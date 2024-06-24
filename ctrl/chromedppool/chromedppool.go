package chromedppool

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type ChromedpInstance struct {
	ctx    context.Context
	cancel context.CancelFunc
	busy   bool
}

type ChromedpPool struct {
	instances []*ChromedpInstance
	mu        sync.Mutex
}

// NewChromedpPool 创建一个新的 ChromedpPool
func NewChromedpPool(size int) (*ChromedpPool, error) {
	pool := &ChromedpPool{
		instances: make([]*ChromedpInstance, size),
	}

	for i := 0; i < size; i++ {
		ctx, cancel := chromedp.NewContext(context.Background())
		instance := &ChromedpInstance{
			ctx:    ctx,
			cancel: cancel,
			busy:   false,
		}
		pool.instances[i] = instance
	}

	return pool, nil
}

// GetInstance 从池中获取一个空闲的 ChromedpInstance
func (p *ChromedpPool) GetInstance() (*ChromedpInstance, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, instance := range p.instances {
		if !instance.busy {
			instance.busy = true
			return instance, nil
		}
	}

	return nil, errors.New("no available chromedp instances")
}

// ReleaseInstance 释放一个 ChromedpInstance 回到池中
func (p *ChromedpPool) ReleaseInstance(instance *ChromedpInstance) {
	p.mu.Lock()
	defer p.mu.Unlock()

	instance.busy = false
}

// RunChromedpTask 使用 ChromedpInstance 运行任务
func (p *ChromedpPool) RunChromedpTask(instance *ChromedpInstance, tasks chromedp.Tasks) error {
	ctx, cancel := context.WithTimeout(instance.ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctx, tasks)
	return err
}

// Shutdown 关闭池中的所有 ChromedpInstance
func (p *ChromedpPool) Shutdown() {
	for _, instance := range p.instances {
		instance.cancel()
	}
}
