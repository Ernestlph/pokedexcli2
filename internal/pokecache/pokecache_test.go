package pokecache

import (
	"testing"
	"time"
)

func TestPokeCache(t *testing.T) {

	// Create a new cache
	cache := NewCache(time.Millisecond * 1000)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("value3"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		val, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("Expected to find key %s in cache", cas.inputKey)
			continue
		}
		if string(val) != string(cas.inputVal) {
			t.Errorf("Expected to find value %s in cache", cas.inputVal)
			continue
		}
	}
}

func TestPokeCacheReap(t *testing.T) {
	// Create a new cache
	cache := NewCache(time.Millisecond * 10)
	cache.Add("key1", []byte("value1"))
	time.Sleep(time.Millisecond * 20)
	val, ok := cache.Get("key1")
	if ok {
		t.Errorf("Expected to not find key %s in cache", "key1")
	}
	if val != nil {
		t.Errorf("Expected to not find value %s in cache", "key1")
	}
}
