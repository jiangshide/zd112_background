package controllers

import (
	"zd112/models"
)

type UserApiController struct {
	BaseApiController
}

func (this *UserApiController) Login() {
	user := new(models.Admin)
	user.Name = this.getString("username", "the username is null!", 1)
	user.Password = this.getString("password", "the password is null!", defaultMinSize)
	this.content = user
	this.response()
}

func (this *UserApiController) Reg() {
	admin := new(models.Admin)
	admin.Name = this.getString("username", "this username is null", 1)
	admin.Password = this.getString("password", "this password is null!", 1)
	this.content = admin
	this.response()
}

func (this *UserApiController) List() {
	this.response()
}

type SpashApiController struct {
	BaseApiController
}

func (this *SpashApiController) Splash() {

}
