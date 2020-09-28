package singleton

import (
	"sync"
	"time"
)

var once = new(sync.Once)

// 懒汉式单例
type LazyManSingleton struct {
	Data string
}

var LazyManSingletonInstance = new(LazyManSingleton)

func GetLazyManSingleton() *LazyManSingleton {
	once.Do(func(){
		LazyManSingletonInstance.Data = "Hello Lazy Man! " + time.Now().Format("2006-01-02 15:04:05")
	})
	return LazyManSingletonInstance
}

func (h *LazyManSingleton) SingletonData() string {
	return h.Data
}
