# BUG_REPRO

## Bug 是什么
更新套餐时，未传的价格和说明字段会被覆盖成 0 和空字符串，导致套餐资料丢失。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestPackageService_UpdatePreservesOmittedFields -count=1
```

测试会创建一个价格为 1999、带说明的套餐，再用空参数调用更新接口。

## 错误信息
```
price = 0, want 1999
```
