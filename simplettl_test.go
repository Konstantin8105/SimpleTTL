package simplettl_test

import (
	"testing"
	"time"

	"github.com/Konstantin8105/SimpleTTL"
)

func TestNewAtHour(t *testing.T) {
	c := simplettl.NewCache(time.Hour)
	if c == nil {
		t.Errorf("Cannot create cache")
	}
}

func TestGet(t *testing.T) {
	c := simplettl.NewCache(2 * time.Second)
	key := "foo"
	value := "bar"
	c.Add(key, value, time.Second)
	// checking before
	{
		r, ok := c.Get(key)
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
		r, ok := c.Get(key)
		if ok {
			t.Errorf("We have \"ok\" from cache after deadline")
		}
		if r == value {
			t.Errorf("We have correct value after deadline")
		}
	}
}

func TestCount(t *testing.T) {
	c := simplettl.NewCache(time.Second / 10) // It will be the second
	key := "foo"
	value := "bar"
	c.Add(key, value, time.Second/2)
	// checking before
	{
		count := c.Count()
		if count != 1 {
			t.Errorf("Cannot take correct counts of map elements before deadline. Count = %v", count)
		}
	}

	time.Sleep(time.Second)
	// checking after
	{
		count := c.Count()
		if count != 0 {
			t.Errorf("Cannot take correct counts of map elements after deadline. Count = %v", count)
		}
	}
}

func TestGetKeys(t *testing.T) {
	c := simplettl.NewCache(2 * time.Second)
	key := "foo"
	value := "bar"
	c.Add(key, value, time.Second)
	// checking before
	{
		keys := c.GetKeys()
		if len(key) != 1 {
			t.Errorf("Not correct length of of keys")
		}
		if keys[0] != value {
			t.Errorf("Returned value is not correct")
		}
	}

	time.Sleep(time.Second)
	// checking after
	{
		keys := c.GetKeys()
		if len(key) != 1 {
			t.Errorf("Not correct length of of keys")
		}
		if keys[0] == value {
			t.Errorf("We have correct value after deadline")
		}
	}
}
