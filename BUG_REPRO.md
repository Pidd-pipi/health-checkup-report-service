# BUG 复现说明

## Bug 是什么
`RegistrationService.Register/UpdateStatus` 与仓储层不检查/不传播 context 取消，登记请求取消或超时后仍被提交。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestRegisterCtxCancel$' -count=1
```

## 错误信息
```
--- FAIL: TestRegisterCtxCancel
    registration_cancel_test.go: Register err = ..., want context.Canceled
```
