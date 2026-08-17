# BUG_REPRO

## Bug 是什么
异常指标列表不传体检人 ID 时，查询条件仍按 examinee_id = 0 过滤，导致列表为空但总数不为 0。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestAbnormalMetricService_ListWithoutExamineeReturnsAll -count=1
```

测试会创建两条不同体检人的异常指标，再用 examinee_id=0 查询全部。

## 错误信息
```
items = 0, total = 2, want both 2
```
