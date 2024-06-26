// Copyright 2017-2024 The Cloudprober Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package resolver provides a caching, non-blocking DNS resolver. All requests
// for cached resources are returned immediately and if cache has expired, an
// offline goroutine is fired to update it.
package resolver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// The max age and the timeout for resolving a target.
const defaultMaxAge = 5 * time.Minute

type cacheRecord struct {
	ip4              net.IP
	ip6              net.IP
	lastUpdatedAt    time.Time
	err              error
	mu               sync.RWMutex
	updateInProgress bool
	callInit         sync.Once
}

// Resolver provides an asynchronous caching DNS resolver.
type Resolver struct {
	cache         map[string]*cacheRecord
	mu            sync.Mutex
	DefaultMaxAge time.Duration
	resolve       func(string) ([]net.IP, error) // used for testing
}

// ipVersion tells if an IP address is IPv4 or IPv6.
func ipVersion(ip net.IP) int {
	if len(ip.To4()) == net.IPv4len {
		return 4
	}
	if len(ip) == net.IPv6len {
		return 6
	}
	return 0
}

// resolveOrTimeout tries to resolve, but times out and returns an error if it
// takes more than defaultMaxAge.
// Has the potential of creating a bunch of pending goroutines if backend
// resolve call has a tendency of indefinitely hanging.
func (r *Resolver) resolveOrTimeout(name string) ([]net.IP, error) {
	var ips []net.IP
	var err error
	doneChan := make(chan struct{})

	go func() {
		ips, err = r.resolve(name)
		close(doneChan)
	}()

	select {
	case <-doneChan:
		return ips, err
	case <-time.After(defaultMaxAge):
		return nil, fmt.Errorf("timed out after %v", defaultMaxAge)
	}
}

// Resolve returns IP address for a name.
// Issues an update call for the cache record if it's older than defaultMaxAge.
func (r *Resolver) Resolve(name string, ipVer int) (net.IP, error) {
	maxAge := r.DefaultMaxAge
	if maxAge == 0 {
		maxAge = defaultMaxAge
	}
	return r.resolveWithMaxAge(name, ipVer, maxAge, nil)
}

// getCacheRecord returns the cache record for the target.
// It must be kept light, as it blocks the main mutex of the map.
func (r *Resolver) getCacheRecord(name string) *cacheRecord {
	r.mu.Lock()
	defer r.mu.Unlock()

	cr := r.cache[name]
	// This will happen only once for a given name.
	if cr == nil {
		cr = &cacheRecord{
			err: errors.New("cache record not initialized yet"),
		}
		r.cache[name] = cr
	}
	return cr
}

// resolveWithMaxAge returns IP address for a name, issuing an update call for
// the cache record if it's older than the argument maxAge.
// refreshedCh channel is primarily used for testing. Method pushes true to
// refreshedCh channel once and if the value is refreshed, or false, if it
// doesn't need refreshing.
func (r *Resolver) resolveWithMaxAge(name string, ipVer int, maxAge time.Duration, refreshedCh chan<- bool) (net.IP, error) {
	cr := r.getCacheRecord(name)
	cr.refreshIfRequired(name, r.resolveOrTimeout, maxAge, refreshedCh)
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	var ip net.IP

	switch ipVer {
	case 0:
		if cr.ip4 != nil {
			ip = cr.ip4
		} else if cr.ip6 != nil {
			ip = cr.ip6
		}
	case 4:
		ip = cr.ip4
	case 6:
		ip = cr.ip6
	default:
		return nil, fmt.Errorf("unknown IP version: %d", ipVer)
	}

	if ip == nil && cr.err == nil {
		return nil, fmt.Errorf("found no IP%d IP for %s", ipVer, name)
	}
	return ip, cr.err
}

// refresh refreshes the cacheRecord by making a call to the provided "resolve" function.
func (cr *cacheRecord) refresh(name string, resolve func(string) ([]net.IP, error), refreshed chan<- bool) {
	// Note that we call backend's resolve outside of the mutex locks and take the lock again
	// to update the cache record once we have the results from the backend.
	ips, err := resolve(name)

	cr.mu.Lock()
	defer cr.mu.Unlock()
	if refreshed != nil {
		refreshed <- true
	}
	cr.err = err
	cr.updateInProgress = false
	// If we have an error, we don't update the cache record so that callers
	// can use cached IP addresses if they want.
	if err != nil {
		return
	}
	cr.lastUpdatedAt = time.Now()
	cr.ip4 = nil
	cr.ip6 = nil
	for _, ip := range ips {
		switch ipVersion(ip) {
		case 4:
			cr.ip4 = ip
		case 6:
			cr.ip6 = ip
		}
	}
}

func (cr *cacheRecord) shouldUpdateNow(maxAge time.Duration) bool {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return !cr.updateInProgress && (time.Since(cr.lastUpdatedAt) >= maxAge || cr.err != nil)
}

// refreshIfRequired does most of the work. Overall goal is to minimize the
// lock period of the cache record. To that end, if the cache record needs
// updating, we do that with the mutex unlocked.
//
// If cache record is new, blocks until it's resolved for the first time.
// If cache record needs updating, kicks off refresh asynchronously.
// If cache record is already being updated or fresh enough, returns immediately.
func (cr *cacheRecord) refreshIfRequired(name string, resolve func(string) ([]net.IP, error), maxAge time.Duration, refreshedCh chan<- bool) {
	cr.callInit.Do(func() { cr.refresh(name, resolve, refreshedCh) })

	// Cache record is old and no update in progress, issue a request to update.
	if cr.shouldUpdateNow(maxAge) {
		cr.mu.Lock()
		cr.updateInProgress = true
		cr.mu.Unlock()
		go cr.refresh(name, resolve, refreshedCh)
	} else if refreshedCh != nil {
		refreshedCh <- false
	}
}

// NewWithResolve returns a new Resolver with the given backend resolver.
// This is useful for testing.
func NewWithResolve(resolveFunc func(string) ([]net.IP, error)) *Resolver {
	return &Resolver{
		cache:         make(map[string]*cacheRecord),
		resolve:       resolveFunc,
		DefaultMaxAge: defaultMaxAge,
	}
}

func parseOverrideAddress(dnsResolverOverride string) (string, string, error) {
	// dnsResolverOverride can be in the format "network://ip:port" or "ip:port",
	// or just "ip". If network is not specified, we use Go's default.
	var networkOverride string
	addrParts := strings.Split(dnsResolverOverride, "://")
	if len(addrParts) == 2 {
		networkOverride = addrParts[0]
		dnsResolverOverride = addrParts[1]
	}

	validNetworks := []string{"", "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6"}
	if !slices.Contains(validNetworks, networkOverride) {
		return "", "", fmt.Errorf("invalid network: %s", networkOverride)
	}

	port := "53"
	// Check if address includes a port number
	ip := net.ParseIP(dnsResolverOverride)
	if ip == nil {
		idx := strings.LastIndex(dnsResolverOverride, ":")
		if idx != -1 {
			port = dnsResolverOverride[idx+1:]
			dnsResolverOverride = dnsResolverOverride[:idx]
		}
	}

	if net.ParseIP(dnsResolverOverride) == nil {
		return "", "", fmt.Errorf("invalid IP address: %s", dnsResolverOverride)
	}

	if _, err := strconv.Atoi(port); err != nil {
		return "", "", fmt.Errorf("invalid port number: %s", port)
	}

	dnsResolverOverride = net.JoinHostPort(dnsResolverOverride, port)
	return networkOverride, dnsResolverOverride, nil
}

func NewWithOverrideResolver(dnsResolverOverride string) (*Resolver, error) {
	networkOverride, dnsResolverOverride, err := parseOverrideAddress(dnsResolverOverride)
	if err != nil {
		return nil, err
	}

	resolveFunc := func(host string) ([]net.IP, error) {
		r := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: defaultMaxAge,
				}
				if networkOverride != "" {
					network = networkOverride
				}
				return d.DialContext(ctx, network, dnsResolverOverride)
			},
		}
		return r.LookupIP(context.Background(), "ip", host)
	}

	return NewWithResolve(resolveFunc), nil
}

// New returns a new Resolver.
func New() *Resolver {
	return NewWithResolve(net.LookupIP)
}
