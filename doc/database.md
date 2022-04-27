# Database

## Tables

**user**

|  字段名  |  数据类型   |   字段描述   |    备注     |
| :------: | :---------: | :----------: | :---------: |
| username | varchar(20) |    用户名    | primary key |
| password | varchar(20) |     密码     |             |
| role_id  |   tinyint   | 角色类型编号 | default: 1  |

**problem**



**submission**

| 字段名          | 数据类型    | 字段描述 | 备注                          |
| --------------- | ----------- | -------- | ----------------------------- |
| submission_id   | bigint      |          |                               |
| submission_time | timestamp   |          |                               |
| problem_id      | bigint      |          |                               |
| username        | varchar(20) |          |                               |
| language        | varchar(20) |          |                               |
| code            | mediumtext  |          |                               |
| status          | varchar(20) |          | "submitted", "pending" ...... |
| run_time        | int         |          | timeunit: millisecond         |



**case**

|   字段名   |  数据类型  | 字段描述 |         备注          |
| :--------: | :--------: | :------: | :-------------------: |
| problem_id |   bigint   |          | composite primary key |
|  case_id   |    int     |          | composite primary key |
|    case    | mediumtext |          |                       |
|  expected  | mediumtext |          |                       |

