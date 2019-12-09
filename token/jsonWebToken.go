package jsonWebToken

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"crypto/rsa"
)

// JWT签名结构
type JsonWebToken struct {
	// 签名所需的私钥
	PrivateKey *rsa.PrivateKey
	// 验签所需的公钥
	PublicKey *rsa.PublicKey
}

// 自定义载荷 必须继承 jwt.StandardClaims
type customClaims struct {
	Data map[string]interface{}
	jwt.StandardClaims
}

func New(tokenConfig *TokenConfig, privateKey []byte, publicKey []byte) (*JsonWebToken, error) {
	var priKey *rsa.PrivateKey
	var err error
	if len(privateKey) != 0 {
		priKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			return nil, err
		}
	}

	var pubKey *rsa.PublicKey
	if len(publicKey) != 0 {
		pubKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
	}

	if tokenConfig == nil {
		tokenConfig = &TokenConfig{}
		tokenConfig.defaultValue()
	}
	return &JsonWebToken{PrivateKey: priKey, PublicKey: pubKey}, nil
}

// 创建token
func (j *JsonWebToken) CreateToken(data map[string]interface{}) (string, error) {
	claims := &customClaims{Data: data, StandardClaims: jwt.StandardClaims{
		//签名生效时间
		NotBefore: int64(time.Now().Unix() - 1000),
		//签名过期时间
		ExpiresAt: int64(time.Now().Unix() + GetTokenConfig().ExpireTime),
		//签名发行者
		Issuer: GetTokenConfig().Issuer,
	}}
	return j.createToken(claims)
}

//解析token
func (j *JsonWebToken) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			return claims.Data, nil
		}
	}
	return nil, err
}

//更新Token
func (j *JsonWebToken) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			jwt.TimeFunc = time.Now
			refreshExpireTime := time.Duration(GetTokenConfig().RefreshExpireTime)
			claims.ExpiresAt = time.Now().Add(refreshExpireTime * time.Second).Unix()
			return j.createToken(claims)
		}
	}
	return "", err
}

// 创建token
func (j *JsonWebToken) createToken(claims *customClaims) (string, error) {
	// RAS 私钥签名/公钥验签
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(j.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
