# go-util

Go 工具包。

- [go-gin-api](https://github.com/xinliangnote/go-gin-api) 开源项目在使用。

## 安装

```go
go get github.com/xinliangnote/go-util
```

## 方法

- aes
    - aes.Encrypt
    - aes.Decrypt
- base64
    - base64.Encode
    - base64.Decode
- convert
    - convert.StrToInt64
    - convert.StrToInt32
    - convert.StrToInt
    - convert.StrToFloat64
    - convert.StrToByte
    - convert.IntToStr
    - convert.IntToInt32
    - convert.IntToInt64
    - convert.Int32ToInt
    - convert.Int32ToInt64
    - convert.Int64ToInt
    - convert.Int64ToInt32
    - convert.Int64ToStr
- hashids
    - hashids.Encrypt
    - hashids.Decrypt
- json
    - json.Encode
    - json.Decode
- mail
    - mail.Send
- math
    - math.RoundedFixed 小数点后 n 位 - 四舍五入
    - math.TruncRound   小数点后 n 位 - 舍去    
- md5
    - md5.MD5
- rsa
    - rsa.PublicEncrypt
    - rsa.PrivateDecrypt
- time
    - time.GetCurrentDate      获取当前的时间 - 字符串
    - time.GetCurrentUnix      获取当前的时间 - Unix时间戳
    - time.GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
    - time.GetCurrentNanoUnix  获取当前的时间 - 纳秒级时间戳
- uuid
    - uuid.GenUUID
- validator
    - validator.Empty
    
## 依赖

- uuid：github.com/google/uuid
- gomail.v2：gopkg.in/gomail.v2
- go-hashids：github.com/speps/go-hashids

## 一起学习

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

