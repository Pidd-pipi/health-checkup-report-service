# BUG_REPRO

## Bug 是什么
检查结果重复录入时，每次异常都会新建一条 AbnormalMetric，旧记录没有清理，导致同一检查项目出现多条重复异常指标。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestExamResultService_ReenterKeepsSingleAbnormalMetric -count=1
```

测试会对同一条 ExamResult 连续录入两次异常值。

## 错误信息
```
abnormal metrics = 2, want 1
```
