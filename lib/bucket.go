package lib

import (
	"sync"
	"time"
)

// 令牌桶，包含四个属性
// 容量（cap）
// 当前令牌数（tokens）
// 互斥锁（lock）
// 每秒放入令牌的速度(rate)
// 上一次放入token的时间（lastTime，时间戳）
type Bucket struct {
	cap    int64
	tokens int64
	rate   int64
	lastTime int64
	lock   sync.Mutex
}

func NewBucket(cap,rate int64) *Bucket {
	if cap <= 0 || rate <= 0 {
		panic("传入参数存在问题")
	}

	bucket := &Bucket{
		cap:    cap,
		tokens: cap,
		rate:   rate,
		lock:   sync.Mutex{},
	}

	return bucket
}

func (this *Bucket) CanGetToken() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	now := time.Now().Unix()
	this.tokens = this.tokens + (now - this.lastTime) * this.rate

	if this.tokens > this.cap {
		this.tokens = this.cap
	}
	this.lastTime = now
	if this.tokens > 0 {
		this.tokens--
		return true
	}

	return false
}