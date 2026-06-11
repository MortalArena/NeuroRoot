package acp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// TaskHandler معالج مهمة ACP
type TaskHandler func(ctx context.Context, input json.RawMessage) (json.RawMessage, error)

// Router يوجّه المهام إلى المعالجات المسجّلة
type Router struct {
	mu       sync.RWMutex
	handlers map[string]TaskHandler
}

// NewRouter ينشئ router
func NewRouter() *Router {
	r := &Router{handlers: make(map[string]TaskHandler)}
	registerBuiltinTasks(r)
	return r
}

// Register يسجّل معالج مهمة
func (r *Router) Register(task string, h TaskHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[task] = h
}

// Handle ينفّذ مهمة
func (r *Router) Handle(ctx context.Context, task string, input json.RawMessage) (json.RawMessage, error) {
	r.mu.RLock()
	h, ok := r.handlers[task]
	r.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("مهمة غير مدعومة: %s", task)
	}
	return h(ctx, input)
}

// SupportedTasks يرجع قائمة المهام المدعومة
func (r *Router) SupportedTasks() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]string, 0, len(r.handlers))
	for t := range r.handlers {
		tasks = append(tasks, t)
	}
	return tasks
}

