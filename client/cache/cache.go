// Added library to cache information
// Will be usefull to avoid user typing multiple times his password

package cache

import "sync"

var Cache = struct {
	sync.RWMutex
	data map[string]string
}{
	data: make(map[string]string),
}

func Set(key, value string) {
	Cache.Lock()
	defer Cache.Unlock()
	Cache.data[key] = value
}

func Get(key string) (string, bool) {
	Cache.RLock()
	defer Cache.RUnlock()
	val, ok := Cache.data[key]
	return val, ok
}
