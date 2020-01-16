[TOC]

# API

## 通用规范

### null 的处理

1. 如果字段的值是string类型，为空返回“”
2. 如果字段的值是指针类型，为空返回null

例如

``` json
//deleted_at 是 TIMESTAMP 类型，属于指针类型，为空返回null
{
  "code": 200,
  "msg": "",
  "data": {
    "id": 2,
    "uid": 123,
    "filename": "团队项目表格",
    "created_at": "2017-08-16T19:47:47+08:00",
    "updated_at": "2017-08-16T19:47:47+08:00",
    "deleted_at": null
  }
}
```

或者

``` json
{
  "code": 404,
  "msg": "Document not found.",
  "data": null
}
```

### RESTfull风格命名规范

| 方法 | 路径 | 动作 | 函数名 |
| :-: | --- | --- | --- |
| GET | /documents | list | DocumentList |
| POST  | /documents | create | DocumentCreate |
| GET | /documents/{did} | show | DocumentShow |
| PUT / PATCH | /documents/{did} | update | DocumentUpdate |
| DELETE | /documents/{did} | destroy | DocumentDestroy |

``` go
const (
    MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH" // RFC 5789
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
)
```

### 状态码


``` go
//Auth 10xxx

//Document 20xxx
var ResultDocument200 = Result{Code: 20200, Msg: "Success."}
var ResultDocument201 = Result{Code: 20201, Msg: "Document created successfully."}
var ResultDocument400 = Result{Code: 20400, Msg: "Fail to create a document."}
var ResultDocument404 = Result{Code: 20404, Msg: "Document not found."}

```


