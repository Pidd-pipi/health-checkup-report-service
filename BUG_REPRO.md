# BUG 复现说明

## Bug 是什么
`UserRepository.FindByID/FindByPhone` 未命中时返回 typed-nil 指针配 nil error，`UserService.UpdateProfile/Login` 忽略错误直接解引用触发空指针 panic。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestUpdateProfileMissingUser404$' -count=1
```

## 错误信息
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation]
  (*UserService).UpdateProfile(...)
      internal/service/user_service.go:85
```
