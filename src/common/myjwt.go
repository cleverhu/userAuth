package common
import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"userAuth/src/model/UserModel"
)

const (
	SECRET = "deeplyThink"
)

func CreateToken(u *UserModel.UserLoginInfoImpl) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      u.ID,
		"username": u.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	token, err := at.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
