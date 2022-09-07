package lock_tool

import (
	"net/http"
	"sync"
)

const url = "http://192.168.5.176:9001/version"

var lock sync.Mutex

// TestWorkLock 全局锁
func TestWorkLock() {
	lock.Lock()
	_, _ = http.Get(url)
	lock.Unlock()
}

// TestWork 不加锁
func TestWork() {
	_, _ = http.Get(url)
}

// TestWorkPrivate 私有锁
func TestWorkPrivate(k int) {
	v := GetLock(k)

	v.Lock()
	_, _ = http.Get(url)
	v.Unlock()
}

type Mutex struct {
	Key int
	M   sync.Mutex
}

func (m *Mutex) Lock() {
	m.M.Lock()
}

func (m *Mutex) Unlock() {
	m.M.Unlock()
}

var MMap = map[int]*Mutex{}
var MMapLock = sync.Mutex{}

func GetLock(key int) *Mutex {
	MMapLock.Lock()
	v, ok := MMap[key]
	if !ok {
		v = &Mutex{
			Key: key,
			M:   sync.Mutex{},
		}
		MMap[key] = v
	}
	MMapLock.Unlock()
	return v
}
