package lruttl

import (
	"strconv"
	"testing"
	"time"
)

func incrementFunc(from int, ttl time.Duration) func(string) (interface{}, time.Duration, error) {
	return func(string) (interface{}, time.Duration, error) {
		r := from
		from++
		return r, ttl, nil
	}
}

func TestCacheWithoutTTL(t *testing.T) {
	cache := NewCache(10, 0, incrementFunc(0, 0))
	for n := 1; n <= 2; n++ {
		for x := 0; x <= 9; x++ {
			r, _ := cache.Get(strconv.Itoa(x))
			if r != x {
				t.Fatalf("itterate %d result %d", x, r)
			}
		}
	}
	r, _ := cache.Get("10")
	if r != 10 {
		t.Fatalf("get 10 result %d", r)
	}
	r, _ = cache.Get("0")
	if r != 11 {
		t.Fatalf("get 0 result %d", r)
	}
}

func TestCacheWithTTL(t *testing.T) {
	cache := NewCache(10, time.Millisecond, incrementFunc(0, 0))
	for n := 0; n <= 1; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func TestCacheWithTTLbyItem(t *testing.T) {
	cache := NewCache(100, 0, incrementFunc(0, time.Millisecond))
	for n := 0; n <= 1; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func TestCacheWithTTLbyItemAndGlobal1(t *testing.T) {
	cache := NewCache(1000, 10*time.Millisecond, incrementFunc(0, time.Millisecond))
	for n := 0; n <= 1; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10 * time.Millisecond)
	for n := 1; n <= 2; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n + 10
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func TestCacheWithTTLbyItemAndGlobal2(t *testing.T) {
	cache := NewCache(1000, time.Millisecond, incrementFunc(0, 10*time.Millisecond))
	for n := 0; n <= 1; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10 * time.Millisecond)
	for n := 2; n <= 3; n++ {
		for i := 1; i <= 2; i++ {
			for x := 0; x <= 9; x++ {
				r, _ := cache.Get(strconv.Itoa(x))
				m := x + 10*n
				if r != m {
					t.Fatalf("itterate %d result %d", m, r)
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
}
