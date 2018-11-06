package user

import (
	"im/web/models/userModel"
	"im/web/controllers"
	"encoding/json"
	"im/web/logic/userLogic"
)

// Operations about Users
type UserController struct {
	controllers.BaseController
}

// type User struct {
// 	Id     string
// 	Username   string `valid:"Required;MinSize(3);MaxSize(32)"`
// 	Password    int    `valid:"Required;MinSize(6);MaxSize(20)"`
// 	CreateTime int64
// }


// @Title Login
// @Description Logs user into the system
// @Param	username		formData 	string	true		"The username for login"
// @Param	password		formData 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var (
		user userModel.User
		retData userLogic.ReturnInfo
	)
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// userInfo := userModel.GetUserInfoByUserName(user.UserName)
	code, msg  := u.CheckParams(user)
	if code != 0 {
		u.Data["json"] = u.RenderDataSimple(code, msg)
		u.ServeJSON()
		return
	}
	code, msg, retData = userLogic.Login(user)

	if code != 0 {
		u.Data["json"] = u.RenderDataSimple(code, msg)
		u.ServeJSON()
		return
	}

	u.Data["json"] = u.RenderData(code, msg, retData)
	u.ServeJSON()
}

// @Title Register
// @Description Register user
// @Param	username		formData 	string	true		"The username for register"
// @Param	password		formData 	string	true		"The password for register"
// @Success 200 {string} register success
// @Failure 403 user not exist
// @router /register [post]
func (u *UserController) Register() {
	var (
		user userModel.User
	)
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	code, msg  := u.CheckParams(user)

	if code != 0 {
		u.Data["json"] = u.RenderDataSimple(code, msg)
		u.ServeJSON()
		return
	}
	code, msg  = userLogic.CheckUserName(user.UserName)
	if code != 0 {
		u.Data["json"] = u.RenderDataSimple(code, msg)
		u.ServeJSON()
		return
	}
	code, msg = userLogic.AddOne(user)
	u.Data["json"] = u.RenderDataSimple(code, msg)
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {

}




