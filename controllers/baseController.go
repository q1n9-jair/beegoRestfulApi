package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type ResultMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//固定常量
const (
	Success    = 200
	SuccessMsg = "success"
	FailureMsg = "Fail"
)

//后续业务增加业务异常常量
const (
	ErrorUserNil = 0
)

/***
成功的返回
*/
func (this *BaseController) Success(data interface{}) {

	res := ResultMsg{
		Code: Success,
		Msg:  SuccessMsg,
		Data: data,
	}
	this.Data["json"] = res
	this.ServeJSON() //对json进行序列化输出
	this.StopRun()
}

/***
失败的返回
*/
func (this *BaseController) Failure(code int) {

	res := ResultMsg{
		Code: code,
		Msg:  FailureMsg,
	}
	this.Data["json"] = res
	this.ServeJSON() //对json进行序列化输出
	this.StopRun()
}

/***
入参封装提取 减少重复代码
*/
func (this *BaseController) JsonStructs(structs interface{}) error {
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &structs)
	if err != nil {
		fmt.Println("err--", err)
		return errors.New("解析失败")
	}
	return nil
}
