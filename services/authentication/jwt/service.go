package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"open_emarker/settings"
	"strconv"
	"time"
)

type UserClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var _jwtService = &jwtService{}

func InstanceFromJwtService() *jwtService {
	return _jwtService
}

func (service jwtService) GenerateToken(id uint) string {
	claims := UserClaims{
		strconv.Itoa(int(id)),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * settings.JWT_EXPIRES_HOURS).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "AUTHENTICATION_SYSTEM",
			Subject:   settings.JWT_LOGIN_SUBJECT,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := token.SignedString([]byte(settings.SECRET_KEY))
	if err != nil {
		panic(err.Error())
	}
	return str
}

func (service jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			_, isValid := token.Method.(*jwt.SigningMethodHMAC)
			if !isValid {
				return nil, NotValidToken{"Not valid token!"}
			}
			return []byte(settings.SECRET_KEY), nil
		},
	)
}

func (service jwtService) GetTokenData(token *jwt.Token) map[string]interface{} {
	return token.Claims.(jwt.MapClaims)
}
