# BUG_REPRO

## Bug 是什么
生成体检报告时没有检查是否所有检查结果都已录入，存在待录入结果时仍然会生成报告。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestReportService_GenerateRequiresAllResultsEntered -count=1
```

测试会创建一条已录入结果和一条待录入结果，再尝试生成报告。

## 错误信息
```
expected AppError code 1404, got ...
```
