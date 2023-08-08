package visitmap

import (
	"fmt"
	"sync"
)

// 配置Map变量安全访问示例

type ExampleConfig struct {
	mu  sync.RWMutex
	Map map[string]string
}

var config *ExampleConfig // 配置对象

// MapKeys 获取Map配置键名切片
// TODO 方式一的第一个方法
func MapKeys() []string {
	config.mu.RLock()
	defer config.mu.RUnlock()
	if len(config.Map) == 0 {
		return nil
	}
	var keys []string
	for k := range config.Map {
		keys = append(keys, k)
	}
	return keys
}

// MapValue map[key]
// TODO 方式一的第二个方法
func MapValue(key string) string {
	config.mu.RLock()
	defer config.mu.RUnlock()
	if v, ok := config.Map[key]; ok {
		return v
	}
	return ""
}

// MapCopy 获取Map配置副本
// TODO 方式二
func MapCopy() map[string]string {
	config.mu.RLock()
	defer config.mu.RUnlock()
	if len(config.Map) == 0 {
		return nil
	}
	var copyMap = make(map[string]string, len(config.Map))
	for k, v := range config.Map {
		copyMap[k] = v
	}
	return copyMap
}

// MapRange map遍历
func MapRange(f func(key, value string)) {
	config.mu.RLock()
	defer config.mu.RUnlock()
	for k, v := range config.Map {
		f(k, v)
	}
}

func LoadConfig() {
	config = new(ExampleConfig)
	config.Map = map[string]string{
		"key1": "config1",
		"key2": "config2",
		"key3": "config3",
	}
}

func ShowResult() {
	// 方式一测试，两步结合使用
	keys := MapKeys()
	for _, key := range keys {
		value := MapValue(key)
		fmt.Printf("方式一：%s: %s\n", key, value)
	}
	println()
	// 方式二
	copyMap := MapCopy()
	for k, v := range copyMap {
		fmt.Printf("方式二：%s: %s\n", k, v)
	}
	println()
	MapRange(func(key, value string) {
		fmt.Println(key, "=", value)
	})

}
