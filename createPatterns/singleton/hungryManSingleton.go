package singleton

import "time"

// 饿汉式单例
type HungryManSingleton struct {
	Data string
}

var SingletonInstance = new(HungryManSingleton)

func init() {
	SingletonInstance.Data = "Hello Hungry Man! " + time.Now().Format("2006-01-02 15:04:05")
}

func GetHungryManSingleton() *HungryManSingleton {
	return SingletonInstance
}

func (h *HungryManSingleton) SingletonData() string {
	return h.Data
}
