package casino

import (
	"encoding/json"
	"math/rand"
	"sync"
)

type Config struct {
	mu     sync.RWMutex
	Module map[string]int32
	Custom *CustomConfig `json:"Custom"`
}

var config *Config

func init() {
	config = &Config{}
	err := json.Unmarshal([]byte(JsonString), config)
	if err != nil {
		panic(err)
	}
}

// WeightChosen 通过权重选择
func WeightChosen(weights []int32) int {
	total := sum(weights)
	if total <= 0 {
		return -1
	}
	r := int32(rand.Intn(int(total)))
	var step int32
	for i, weight := range weights {
		step += weight
		if r < step {
			return i
		}
	}
	return -1
}

func sum(values []int32) int32 {
	var total int32
	for _, v := range values {
		total += v
	}
	return total
}
