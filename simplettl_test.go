package simplettl_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Konstantin8105/SimpleTTL"
)

func TestNewAtHour(t *testing.T) {
	cache := simplettl.NewCache(time.Hour)
	if cache == nil {
		t.Errorf("Cannot create cache")
	}
}

func TestGet(t *testing.T) {
	cache := simplettl.NewCache(2 * time.Second)
	key := "foo"
	value := "bar"
	cache.Add(key, value, time.Second)
	// checking before
	{
		r, ok := cache.Get(key)
		if !ok {
			t.Errorf("Cannot take \"ok\" from cache")
		}
		if r != value {
			t.Errorf("Returned value is not correct")
		}
	}

	time.Sleep(time.Second)
	// checking after
	{
		r, ok := cache.Get(key)
		if ok {
			t.Errorf("We have \"ok\" from cache after deadline")
		}
		if r == value {
			t.Errorf("We have correct value after deadline")
		}
	}
}

func TestCount(t *testing.T) {
	cache := simplettl.NewCache(time.Second / 10) // It will be the second
	key := "foo"
	value := "bar"
	cache.Add(key, value, time.Second/2)
	// checking before
	{
		count := cache.Count()
		if count != 1 {
			t.Errorf("Cannot take correct counts of map elements before deadline. Count = %v", count)
		}
	}

	time.Sleep(2 * time.Second)
	// checking after
	{
		count := cache.Count()
		if count != 0 {
			t.Errorf("Cannot take correct counts of map elements after deadline. Count = %v", count)
		}
	}
}

func TestGetKeys(t *testing.T) {
	cache := simplettl.NewCache(time.Second)
	key := "foo"
	value := "bar"
	cache.Add(key, value, time.Second)
	// checking before
	{
		keys := cache.GetKeys()
		if len(keys) != 1 {
			t.Errorf("Not correct length of of keys before deadline. Len = %v", len(keys))
		}
		if keys[0] != key {
			t.Errorf("Returned value is not correct")
		}
	}

	time.Sleep(time.Second * 2)
	// checking after
	{
		keys := cache.GetKeys()
		if len(keys) != 0 {
			t.Errorf("Not correct length of of keys after deadline. Len = %v", len(keys))
		}
	}
}

func TestCondition(t *testing.T) {
	cache := simplettl.NewCache(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		cache.Add("Dream", 42, time.Hour)
		wg.Done()
	}()

	cache.Add("Dream", -1, time.Second)

	wg.Wait()

	if cache.Count() != 1 {
		t.Errorf("Checking race condition")
	}
}

func ExampleSimple() {
	cache := simplettl.NewCache(2 * time.Second)
	key := "foo"
	value := "bar"
	cache.Add(key, value, time.Second)
	if r, ok := cache.Get(key); ok {
		fmt.Printf("Value for key %v is %v", key, r)
	}
	// Output: Value for key foo is bar
}
