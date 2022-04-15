#### API版本： v1

## 响应格式
```
{
    code: errorcode, //子错误码
    msg: str,        //子错误码文本串
    data: object,    //数据(可选)
}
```
以下只显示**data**

## 一. 用户操作

#### (1) 用户登录
Request:
```
POST /api/v1/login
{
    "username" : str,
    "passwd"   : str, //hash
}
```
Response:
```
成功设置Cookie
{
    "userid": str
}
```

#### (2) 用户退出
Request:
```
POST /api/v1/logout
{
    "userid" : str,
}
```
Response:

#### (3) 获取用户信息
Request:
```
GET /api/v1/userinfo
{
    "userid" : str,
}
```
Response:
```
{
    "userid": str,
    "username" : str,
    "desc" : str,
}
```

#### (4) 修改用户信息
Request:
```
PUT /api/v1/userinfo
{
    "userid" : str,
    "username" : str, //opt
    "desc" : str,     //opt
}
```
Response:

#### (5) 用户注册
Request:
```
POST /api/v1/userinfo
{
    "username" : str,
    "desc" : str
}
```
Response:
```
{
    "userid" : str,
}
```






