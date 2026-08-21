# BUG 复现说明

## Bug 是什么
`ReportService` 生成报告时用 `results[:0]` 原地压缩异常项，共享底层数组污染原始结果，导致 PDF 报告行里正常检查项目丢失。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestReportGenerateKeepsAllResultRows$' -count=1
```

## 错误信息
```
--- FAIL: TestReportGenerateKeepsAllResultRows
    report_rows_race_test.go:31: row[0] item = "血糖", want "血压"
```
