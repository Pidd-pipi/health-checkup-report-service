# BUG 复现说明

## Bug 是什么
`GroupOrderRepository.FindByID` 用 `%v` 包装 ErrNotFound 使错误链断裂，`EnterpriseService.DeliverReports` 的 `errors.Is` 判断失效，不存在的团检订单被误判为 500，且已交付订单可被重复交付。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestDeliverReportsMissingOrder404$' -count=1
```

## 错误信息
```
--- FAIL: TestDeliverReportsMissingOrder404
    enterprise_service_test.go: err = 查询团检订单失败: ..., want NotFound(404)
```
