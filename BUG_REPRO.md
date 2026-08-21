# BUG 复现说明

## Bug 是什么
`StatsService.Dashboard` 与 `DashboardStats.Summarize` 用未初始化的 nil map 汇总统计数据，写入时触发 panic。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestDashboardMonthlySummaryNoPanic$' -count=1
```

## 错误信息
```
panic: assignment to entry in nil map
goroutine 6 [running]:
  (*StatsService).Dashboard(...)
      internal/service/stats_service.go:48
```
