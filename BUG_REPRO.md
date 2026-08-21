# BUG 复现说明

## Bug 是什么
`ExamResultRepository.FindByID` 用 `%v` 包装 ErrNotFound 使错误链断裂，`ExamResultService.Enter/Review` 的 `errors.Is` 判断失效，不存在的检查结果被误判为 500，且已录入结果可重复录入。

## 如何触发
在 `backend` 目录执行：

```
go test ./internal/service -run '^TestEnterMissingResult404$' -count=1
```

## 错误信息
```
--- FAIL: TestEnterMissingResult404
    exam_result_conflict_test.go: err = 查询检查结果失败: ..., want NotFound(404)
```
