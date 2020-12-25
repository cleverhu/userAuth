package Getter

import (
	"time"
	"userAuth/src/common"
	"userAuth/src/dbs"
	"userAuth/src/model/UserModel"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	Login(u *UserModel.UserLoginInfoImpl) string
}

type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (this *UserGetterImpl) Login(u *UserModel.UserLoginInfoImpl) string {

	u.PassWord = common.MD5(u.PassWord)
	token := dbs.Rds.Get("users:username:" + u.Name + u.PassWord).Val()
	if token == "" {
		token = dbs.Rds.Get("users:email:" + u.Name + u.PassWord).Val()
	}
	if token != "" {
		return token
	}

	tmp :=UserModel.NewUserLoginInfoImpl()

	if dbs.Orm.Table("t_user").Find(&tmp, "u_name = ? or u_email = ?", u.Name, u.Name).RecordNotFound() {
		return ""
	} else {
		if dbs.Orm.Table("t_user").Find(&u, "u_name = ?  and u_password = ?", tmp.Name,  u.PassWord).RowsAffected == 1 {

			if token, err := common.CreateToken(u); err != nil {
				return ""
			} else {
				dbs.Rds.Set("users:username:"+u.Name+u.PassWord, token, 10*time.Second)
				dbs.Rds.Set("users:email:"+u.Email+u.PassWord, token, 10*time.Second)
				return token
			}
		} else {
			return ""
		}

	}
}
