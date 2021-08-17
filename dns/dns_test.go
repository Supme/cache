package dns

import (
	"net"
	"testing"
	"time"
)

const (
	testHost = "mail.ru."
	testAddr = "77.88.8.8"
	testCname = "mail.supme.ru."
)

func TestResolver_LookupAddr(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupAddr(testAddr)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", r)
}

func TestResolver_LookupCNAME(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupCNAME(testCname)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestResolver_LookupHost(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupHost(testHost)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", r)
}

func TestResolver_LookupIP(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupIP(testHost)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", r)
}

func TestResolver_LookupMX(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupMX(testHost)
	if err != nil {
		t.Error(err)
	}
	for i := range r {
		t.Logf("%+v", *r[i])
	}
}

func TestResolver_LookupNS(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupNS(testHost)
	if err != nil {
		t.Error(err)
	}
	for i := range r {
		t.Logf("%+v", *r[i])
	}
}

func TestResolver_LookupPort(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupPort("tcp", "imaps")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestResolver_LookupSRV(t *testing.T) {
	//s, r, err := net.LookupSRV("imaps", "tcp", "mail.ru.")
	rslvr := NewCacheResolver(100, time.Minute)
	s, r, err := rslvr.LookupSRV("imaps", "tcp", testHost)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
	for i := range r {
		t.Logf("%+v", *r[i])
	}
}

func TestResolver_LookupTXT(t *testing.T) {
	rslvr := NewCacheResolver(100, time.Minute)
	r, err := rslvr.LookupTXT(testHost)
	if err != nil {
		t.Error(err)
	}
	for i := range r {
		t.Log(r[i])
	}
}

func BenchmarkLookupIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := net.LookupIP(testHost)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCache_LookupIP(b *testing.B) {
	rslvr := NewCacheResolver(100, time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rslvr.LookupIP(testHost)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLookupMX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := net.LookupMX(testHost)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCache_LookupMX(b *testing.B) {
	rslvr := NewCacheResolver(100, time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rslvr.LookupMX(testHost)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLookupAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := net.LookupAddr(testAddr)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCache_LookupAddr(b *testing.B) {
	rslvr := NewCacheResolver(100, time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rslvr.LookupAddr(testAddr)
		if err != nil {
			b.Error(err)
		}
	}
}