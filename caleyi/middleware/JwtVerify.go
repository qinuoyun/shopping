package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID         int64  `json:"id"`
	AccountId  int64  `json:"accountId"`  //A端主账号id
	BusinessID int64  `json:"businessID"` //B端主账号id
	Openid     string `json:"openid"`     //微信openid
	Name       string `json:"name"`
	Username   string `json:"username"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f489")
	// effectTime = 2 * time.Minute //两分钟
)

// 加载配置
// token有效时间（纳秒）
var effectTime = time.Duration(getJwtInt()) * time.Minute //分钟单位

// 写个返回int64-默认2个小时
func getJwtInt() int64 {
	//加载配置

	num := "72000"
	intNum, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 2 * 60 //默认2个小时
	} else {
		return intNum
	}
}

// TokenOutTime 返回超时时间
func TokenOutTime(claims *UserClaims) int64 {
	return time.Now().Add(effectTime).Unix()
}

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) interface{} {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

// JwtVerify 验证token
func JwtVerify(c *gin.Context) {
	//获取白名单列表
	whitelistItems := ci.C("whitelist.items")
	//转换列表数据
	whiteList := strings.Split(whitelistItems, ",")
	if checkWhiteList(whiteList, c.Request.URL.Path) { //不需要token验证-具体路径
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.GetHeader("authorization")
	}
	if token == "" {
		panic("token 不存在")
	}
	//验证token，并存储在请求中
	c.Set("user", ParseToken(token))
}

// ParseToken 解析Token
func ParseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("The token is invalid")
	}
	return claims
}

// Refresh 更新token
func Refresh(tokenString string) interface{} {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("The token is invalid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(effectTime).Unix()
	return GenerateToken(claims)
}

// 检查白名单
func checkWhiteList(whiteList []string, path string) bool {
	for _, p := range whiteList {
		if strings.HasPrefix(p, "^") {
			matched, _ := regexp.MatchString(p, path)
			if matched {
				return true
			}
		} else {
			if p == path {
				return true
			}
		}
	}
	return false
}
