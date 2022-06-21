package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = os.Getenv("MailPassword")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretID = "AKIDQfd2A0auxhCRP8ZiUyA4h49zlSyHKVOy"
var TencentSecretKey = "MPBh02mNuIDxXw62h6GBrHSOAftrZ1NX"
var CosBucket = "https://1-1312455940.cos.ap-chongqing.myqcloud.com"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
