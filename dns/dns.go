package dns

import (
	"cache/lruttl"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type resolver struct {
	cache *lruttl.LRU
	*net.Resolver
}

// NewCacheResolver return cached wrapper over net.Resolver with custom its own analogue functions Lookup*
func NewCacheResolver(capacity int, defaultTTL time.Duration) *resolver {
		cache := lruttl.NewCache(capacity, defaultTTL, get)
		r := new(resolver)
		r.cache = cache
		return r
}

const keyDelimiter = "||"

const (
	addr = iota
	cname
	host
	ip
	mx
	ns
	srv
	txt
)

type srvRes struct {
	cname string
	addrs []*net.SRV
}

func get(key string) (interface{}, time.Duration, error) {
	split := strings.Split(key, keyDelimiter)
	if len(split) < 2 {
		return nil, 0, fmt.Errorf("not valid cache key: '%s' (split to %d by '%#v')", key, len(split), split)
	}

	switch atoi(split[0]) {
	case addr:
		r, err := net.LookupAddr(split[1])
		return r, 0, err
	case cname:
		r, err := net.LookupCNAME(split[1])
		return r, 0, err
	case host:
		r, err := net.LookupHost(split[1])
		return r, 0, err
	case ip:
		r, err := net.LookupIP(split[1])
		return r, 0, err
	case mx:
		r, err := net.LookupMX(split[1])
		return r, 0, err
	case ns:
		r, err := net.LookupNS(split[1])
		return r, 0, err
	case srv:
		if len(split) < 4 {
			return "", 0, fmt.Errorf("not valid cache key for srv: '%s' (split to %d by '%#v')", key, len(split), split)
		}
		c, a, err := net.LookupSRV(split[1], split[2], split[3])
		var r srvRes
		r.cname = c
		r.addrs = a
		return r, 0, err
	case txt:
		r, err := net.LookupTXT(split[1])
		return r, 0, err
	default:
		return nil, 0, fmt.Errorf("cache key type not supported")
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func prepareKey(t int, s ...string) string {
	if len(s) > 4 {
		return ""
	}
	tr := make([]string, 4)
	tr[0] = strconv.Itoa(t)
	for i := range s {
		tr[i+1] = strings.ToLower(strings.TrimSpace(s[i]))
	}
	return strings.Join(tr, keyDelimiter)
}

func (r resolver) LookupAddr(name string) ([]string, error) {
	l, err := r.cache.Get(prepareKey(addr, name))
	if err != nil {
		return nil, err
	}
	return l.([]string), nil
}

func (r resolver) LookupCNAME(name string) (string, error) {
	l, err := r.cache.Get(prepareKey(cname, name))
	if err != nil {
		return "", err
	}
	return l.(string), nil
}

func (r resolver) LookupHost(name string) ([]string, error) {
	l, err := r.cache.Get(prepareKey(host, name))
	if err != nil {
		return nil, err
	}
	return l.([]string), nil
}

func (r resolver) LookupIP(name string) ([]net.IP, error) {
	l, err := r.cache.Get(prepareKey(ip, name))
	if err != nil {
		return nil, err
	}
	return l.([]net.IP), nil
}

func (r resolver) LookupMX(name string) ([]*net.MX, error) {
	l, err := r.cache.Get(prepareKey(mx, name))
	if err != nil {
		return nil, err
	}
	return l.([]*net.MX), nil
}

func (r resolver) LookupNS(name string) ([]*net.NS, error) {
	l, err := r.cache.Get(prepareKey(ns, name))
	if err != nil {
		return nil, err
	}
	return l.([]*net.NS), nil
}

func (r resolver) LookupPort(network, service string) (int, error) {
	// Do not need to cache
	return net.LookupPort(network, service)
}

func (r resolver) LookupSRV(service, proto, name string) (string, []*net.SRV, error) {
	l, err := r.cache.Get(prepareKey(srv, service, proto, name))
	if err != nil {
		return "", nil, err
	}
	s := l.(srvRes)
	return s.cname, s.addrs, nil
}

func (r resolver) LookupTXT(name string) ([]string, error) {
	l, err := r.cache.Get(prepareKey(txt, name))
	if err != nil {
		return nil, err
	}
	return l.([]string), nil
}
