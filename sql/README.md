# mysql

## database/sql APIs

### func Open

```go
func Open(driverName, dataSourceName string) (*DB, error)
```

`Open` 只是验证数据库连接参数是否有效，并没有实际上创建连接。如果需要验证「数据源有效」，需要调用 `Ping`。

每个进程只需要 `sql.Open()` 一次：

> database/sql 自己会维护连接池，每次 sql.Open() 会新建一套连接池。虽然不会报错，但是会导致资源浪费。

### type DB

底层数据库连接池的操作句柄，它是 goroutine「并发安全」的。

#### func (*DB) Ping

```go
func (db *DB) Ping() error
```

验证数据库这个连接是不是还存活着，如果必要就建立连接。

#### func (*DB) Close

```go
func (db *DB) Close() error
```

关闭数据库连接，阻止新的 query 执行，对于已经执行的 query，它会等待关闭。

#### func (*DB) Query

```go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

执行 `SELECT` 操作返回结果。 `args` 参数是 `query` 中占位符代表的实参，比如：

```go
rows, err := db.Query("SELECT name FROM user WHERE name = ? or name = ?", "jack", "luck")
```

#### func (*DB) Exec

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

像 `INSERT`、`UPDATE` 等通过 `Exec` 来执行。

#### func (*DB) QueryRow

```go
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

返回一条数据，如果没有数据 `Row.Scan` 会返回错误 `ErrNoRows`。

### type Rows

代表了 `(*DB).Query` 的查询结果，指针指向结果的第一行，取其中的数据需要迭代查询。

#### func (*Rows) Close

```go
func (rs *Rows) Close() error
```

首先，这个函数是「幂等」的，可以多次调用。当调用 `(*Rows).Next()` 返回 `false` 或者没有结果集时，它会自动的关闭。

为什么需要手动 `defer` 关闭它呢？

> 如果不 Close，这个 row 就一直保持着与当前 connection pool 中的 sql 连接的依赖关系，连接也就不会被释放。最终导致资源不必要的堆积，甚至崩溃。

#### func (*Rows) Next

```go
func (rs *Rows) Next() bool
```

在调用 `(*Rows).Scan` 前，「必须」要通过 `(.Rows).Next` 来为它准备数据。

#### func (*Rows) Scan

```go
func (rs *Rows) Scan(dest ...interface{}) error
```

`dest` 对应 `SELECT` 中的字段，需要传递「指针」。

#### func (*Rows) Err

```go
func (rs *Rows) Err() error
```

返回结果集迭代过程中的错误。

## Usage

### 直接使用

```
import (
	"database/sql"
	_ "github.com/go-sql-dirver/mysql"
)
db, err := sql.Open("mysql", "<user>:<pwd>@tcp(<host>:<port>)/<db>")
if err != nil {
    panic(err)
}
defer db.Close()

// SELECT

```

DSN 配置详情：https://github.com/go-sql-driver/mysql#dsn-data-source-name

Query 步骤：



## Demos

> Link: https://www.cnblogs.com/hanyouchun/p/6708037.html

### Init

create mysql database and table:

```text
CREATE DATABASE `gotest`;

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT '',
  `age` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` VALUES (null, "jack", 11), (null, "lucy", 22), (null, "dylan", null), (null, null, 12);
```

## 

