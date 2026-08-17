# BUG_REPRO

## Bug 是什么
检查项目已经被检查结果引用时，仍然可以调用删除接口，导致后续报告里项目名称等字段变成空。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestPackageService_DeleteItemInUseFails -count=1
```

测试会先给套餐添加检查项目，再为该项目录入一条检查结果，然后删除该项目。

## 错误信息
```
expected error when deleting item in use
```
