# 学习笔记

## Question
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## Answer 

Wrap the error at where it happens (low level error should be handled by upper level/business level).

## Learning Material

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

## Answer in short:
dao:
```go
return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
```

business:
```go
if errors.Is(err, code.NotFound} {

}
```