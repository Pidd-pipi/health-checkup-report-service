# BUG 复现说明

## Bug 是什么
`RateLimiter` 的观测方法（Snapshot/Count/Reset/Peek/ClearAll）以及 `Limit` 在加锁区外直接读写内部 `buckets` map，并发调用时存在 data race。

## 如何触发
在 `backend` 目录执行：

```
go test -race ./internal/middleware -run '^TestRateLimiterConcurrentSnapshotAndLimit$' -count=1
```

## 错误信息
```
WARNING: DATA RACE
Read at ... by goroutine 9:
  (*RateLimiter).Snapshot()
      internal/middleware/rate_limiter_observability.go:5
Previous write at ... by goroutine 11:
  runtime.mapaccess1_faststr()
  (*RateLimiter).Reset()
      internal/middleware/rate_limiter_observability.go:10
```
