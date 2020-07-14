package controllers

import (
	"beeApi/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type UserController struct {
	BaseController
}

type LoginStruct struct {
	Name string `json:"name"`
	Pwd  string `json:pwd`
}

func (userController *UserController) Login() {
	var loginStruct = &LoginStruct{}
	userController.JsonStructs(loginStruct)
	userInfo, err := models.Login(loginStruct.Name, loginStruct.Pwd)
	fmt.Println(userInfo)
	if err != nil {
		fmt.Println(err)
	}

	if userInfo == nil {
		userController.Failure(ErrorUserNil)
	}
	userController.SetSession("uid", userInfo.Id)
	userController.SetSession("userInfo", *userInfo)
	rc := models.RedisPool().Get()
	defer rc.Close()
	_, rcerr := rc.Do("set", "user:id:"+strconv.Itoa(userInfo.Id), userInfo.Token)
	if rcerr != nil {
		logs.Error(rcerr)
	}
	userController.Success(userInfo)
}

/***
注册
*/
func (userController *UserController) Regin() {
	var userInfo = &models.TUser{}
	userController.JsonStructs(userInfo)
	count, err := models.AddTUser(userInfo)
	if err != nil {
		fmt.Println(err)
	}
	if count == 0 {
		userController.Failure(ErrorUserNil)
	}
	userController.Success(userInfo)
}

/***
根据id查找用户
*/
func (userController *UserController) GetUser() {
	var userInfo = &models.TUser{}
	userController.JsonStructs(userInfo)
	retUser, err := models.GetUser(*userInfo)
	if err != nil {
		fmt.Println(err)
	}
	userController.Success(retUser)
}
