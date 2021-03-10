package lib

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type cacheData struct {
	key string
	value interface{}
	expireAt time.Time
}

func newCacheData(key string,value interface{}) *cacheData {
	return &cacheData{
		key:key,
		value: value,
	}
}

type LRUCache struct {
	maxSize int
	eList *list.List
	edata map[string]*list.Element
	lock sync.Mutex
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		maxSize: 1024,
		eList: list.New(),
		edata: make(map[string]*list.Element),
		lock:  sync.Mutex{},
	}
}

// 获取缓存
func (this *LRUCache) Get(key string) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if v,ok := this.edata[key]; ok {
		this.eList.MoveToFront(v)
		return v.Value.(*cacheData).value
	}

	return nil
}

func (this *LRUCache) Set(key string,new interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	newCache := newCacheData(key,new)
	if v,ok := this.edata[key]; ok {
		v.Value = newCache
		this.eList.MoveToFront(v)
	} else {
		this.edata[key] = this.eList.PushFront(newCache)
	}
}

func (this *LRUCache) Print() {
	e := this.eList.Front()
	for e != nil {
		fmt.Println(e.Value.(*cacheData).value)
		e = e.Next()
	}
}

// 删除最后一个元素
func (this *LRUCache) RemoveOldest() {
	this.lock.Lock()
	defer this.lock.Unlock()
	back := this.eList.Back()
	if back == nil {
		return
	}
	this.removeItem(back)
}

func (this *LRUCache) removeItem(e *list.Element) {
	k := e.Value.(*cacheData).key
	delete(this.edata,k) // 删除map里的key
	this.eList.Remove(e)
}