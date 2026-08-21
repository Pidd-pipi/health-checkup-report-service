# BUG 复现说明

## Bug 是什么
`ExamineeService.BatchImport` 与仓储层不检查/不传播 context 取消，客户端超时取消后仍在往库里插入体检人。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestBatchImportStopsOnCancelledContext$' -count=1
```

## 错误信息
```
--- FAIL: TestBatchImportStopsOnCancelledContext
    examinee_service_test.go: BatchImport err = ..., want context.Canceled
```
