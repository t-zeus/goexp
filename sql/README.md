# mysql

> Link: https://www.cnblogs.com/hanyouchun/p/6708037.html

## Init

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
