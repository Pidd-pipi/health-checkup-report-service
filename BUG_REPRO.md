# BUG 复现说明

## Bug 是什么
`PackageService.BatchUpdateItems/BatchDeleteItems` 循环内 defer 堆积，`BulkUpdate/BulkDelete` 命名返回值被 defer 里的 commit 覆盖吞错，错误分支漏回滚，出错时数据被提交且错误丢失。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestBatchUpdateItemsReleasesAndKeepsError$' -count=1
```

## 错误信息
```
--- FAIL: TestBatchUpdateItemsReleasesAndKeepsError
    package_service_test.go: BatchUpdateItems should return error for empty item name
```
