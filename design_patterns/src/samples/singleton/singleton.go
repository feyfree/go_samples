package singleton

import "sync"

type Singleton struct {
}

var instance *Singleton
var once sync.Once

// 使用 once 保证只会执行一次, 饿汉操作
func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}
