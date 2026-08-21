package middleware

// Snapshot 返回每个限流键当前的令牌桶状态，供监控统计使用。
func (rl *RateLimiter) Snapshot() map[string]*bucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	out := make(map[string]*bucket, len(rl.buckets))
	for k, v := range rl.buckets {
		out[k] = v
	}
	return out
}

// Reset 清空某个键的限流桶。
func (rl *RateLimiter) Reset(key string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.buckets, key)
}

// Count 返回当前被限流跟踪的键数量。
func (rl *RateLimiter) Count() int {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return len(rl.buckets)
}

// Peek 返回某个键当前的令牌桶（不存在则返回 false）。
func (rl *RateLimiter) Peek(key string) (*bucket, bool) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	b, ok := rl.buckets[key]
	return b, ok
}

// ClearAll 清空全部限流桶。
func (rl *RateLimiter) ClearAll() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for k := range rl.buckets {
		delete(rl.buckets, k)
	}
}
