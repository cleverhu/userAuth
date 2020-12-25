package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"userAuth/src/common"
	"userAuth/src/data/Getter"
	"userAuth/src/model/UserModel"
)

func LoginGet(ctx *gin.Context) {
	referer := ctx.Query("git ")
	if referer == "" {
		ctx.String(400, "跳转地址错误")
		return
	}
	err := fmt.Errorf("")
	token, _ := ctx.Cookie("jwt")
	if token == "" {
		query := ctx.Query("token")
		if _, err = common.ParseToken(query); err == nil {
			ctx.Redirect(http.StatusFound, referer+"?token="+token)
		} else {
			ctx.SetCookie("jwt", "", -1, "/", "auth.deeplythink.com", false, true)
			ctx.HTML(200, "login.html", nil)

			//ctx.Redirect(http.StatusFound, "http://auth.deeplythink.com/?redirect_url="+referer)
		}
	} else {
		if _, err = common.ParseToken(token); err == nil {
			ctx.Redirect(http.StatusFound, referer+"?token="+token)
		} else {
			ctx.SetCookie("jwt", "", -1, "/", "auth.deeplythink.com", false, true)
			ctx.HTML(200, "login.html", nil)

		}
	}
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "auth.deeplythink.com", false, true)
	//ctx.Redirect(http.StatusFound, "http://auth.deeplythink.com")
}

func LoginPost(ctx *gin.Context) {
	u := UserModel.NewUserLoginInfoImpl()
	referer := ctx.GetHeader("Referer")
	fmt.Println(referer)
	myUrl, _ := url.Parse(referer)

	r := &http.Request{
		Method:           "",
		URL:              myUrl,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}
	_ = r.ParseForm()
	redirect := r.Form.Get("redirect_url")
	if redirect == "" {
		ctx.String(400, "跳转地址错误", nil)
		return
	}
	fmt.Println(redirect)

	if err := ctx.ShouldBind(&u); err != nil {
		panic(err)
	}
	token := Getter.UserGetter.Login(u)

	if token != "" {
		ctx.SetCookie("jwt", token, 3600, "/", "auth.deeplythink.com", false, true)

		ctx.Redirect(http.StatusFound, redirect+"?token="+token)
	} else {
		ctx.Redirect(http.StatusFound, "http://auth.deeplythink.com/?redirect_url="+redirect)
	}
}
