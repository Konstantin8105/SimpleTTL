# SimpleTTL


[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/SimpleTTL/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/SimpleTTL?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/SimpleTTL.svg?branch=master)](https://travis-ci.org/Konstantin8105/SimpleTTL)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/SimpleTTL)](https://goreportcard.com/report/github.com/Konstantin8105/SimpleTTL)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/SimpleTTL/blob/master/LICENSE)


Simple TTL on golang for map[string]interface{}, so keys is string and values is something.

See more detail: https://en.wikipedia.org/wiki/Time_to_live

Minimal example of using:

```golang
func main(){
	cache := simplettl.NewCache(2 * time.Second)
	key := "foo"
	value := "bar"
	cache.Add(key, value, time.Second)
	if r, ok := cache.Get(key); ok {
		fmt.Printf("Value for key %v is %v",key,r)
	}
}
```
