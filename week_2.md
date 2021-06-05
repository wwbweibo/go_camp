# 第二周作业

Q:我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

对于`sql.ErrNoRows`错误，包中的注释如下

> ErrNoRows is returned by Scan when QueryRow doesn't return a row. In such a case, QueryRow returns a placeholder *Row value that defers this error until a Scan.

而对于QueryRow方法，注释如下

> QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

所以，该错误出现在我们想从数据库中获取一条记录，而该记录并不存在时。

回到问题中，错误是在dao层出现的，dao层只是负责处理和数据查询相关的逻辑。而查询一条记录，但是该记录并不存在的处理通常是应该交由上层业务逻辑层来具体处理。  
在大多数语言中，dao层对于这种情况的处理通常是返回一个空对象，业务逻辑层判空从而进行对应的处理。但是这种情况下返回值就会存在二义性，在go中，并不推荐这么做。  
参考go中map的实现，使用一个bool值表示对应的Key是否存在。但是这样的做法并不十分适合dao层，因为在dao层与数据库的交互中不仅仅存在`ErrNoRows`错误，也会存在其他的错误。如果我们处理了ErrNoRows错误，并返回了一个bool值来标识行是否存在，我们仍然需要返回其他的错误给业务逻辑层，并不优雅。  

``` go
func QueryUserName(id int) (string, bool, error) {
    userName := ""
    err := db.QueryRow("select user_name from users where id=?", id).Scan(&userName)
    if err == nil {
        return userName, true, nil
    }
    if err == sql.ErrNoRows {
        return userName, false, nil
    } else {
        return userName, false, error
    }
}
```

综上，我认为ErrNoRows错误应该Wrap返回上层。向下面这样

```go
// dao 层我们就可以这样写
func QueryUserName(id int) (string, error) {
    userName := ""
    err := db.QueryRow("select user_name from users where id=?", id).Scan(&userName)
    return userName, errors.Wrap(err, "error when get user name")
}
```
