# BUG 复现说明

## Bug 是什么
`AbnormalMetricService.UpdateFollowUp` 状态机转换表漏了 processing 中间态且 done 终态校验缺失，`List/CountPending` 按旧状态过滤漏掉 processing，复查跟踪可被非法回退并从列表消失。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestUpdateFollowUpDoneIsTerminal$' -count=1
```

## 错误信息
```
--- FAIL: TestUpdateFollowUpDoneIsTerminal
    abnormal_metric_followup_test.go: err = ..., want Conflict(409)
```
