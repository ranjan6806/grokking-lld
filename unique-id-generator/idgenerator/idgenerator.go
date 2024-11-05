package idgenerator

import (
	"fmt"
	"sync"
	"time"
)

type Generator interface {
	GenerateID() string
}

type CounterGenerator struct {
	counter int64
	mtx     sync.RWMutex
}

var (
	counterGeneratorInstance *CounterGenerator
	once                     sync.Once
)

func NewCounterGenerator() *CounterGenerator {
	once.Do(func() {
		counterGeneratorInstance = &CounterGenerator{
			counter: time.Now().UnixNano(),
		}
	})
	return counterGeneratorInstance
}

func (g *CounterGenerator) GenerateID() string {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	g.counter++
	return fmt.Sprintf("ID-%d", g.counter)
}
