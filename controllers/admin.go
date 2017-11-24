package controllers

import (
	"zd112/models"
	"github.com/astaxie/beego"
	"time"
	"zd112/utils"
	"strconv"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) List() {
	this.pageTitle("管理员管理")
	this.display("backstage/admin/list")
	//go download()
}

func download(){
	mingzu :=[...]string{"阿昌族","白族","保安族","布朗族","布依族","藏族","朝鲜族","达翰尔族","傣族","昂德族","东乡族","侗族","独龙族","俄罗斯族","鄂伦春族","鄂温克族","高山族","哈尼族","哈萨克族","汉族","赫哲族","回族","基诺族","京族",
	"景颇族","柯尔克孜族","拉祜族","黎族","傈僳族","珞巴族","满族","毛南族","门巴族","蒙古族","苗族","仫佬族","纳西族","怒族","普米族","羌族","撒拉族","畲族","水族","塔吉克族","塔塔尔族","土家族","图族","佤族","维吾尔族","乌孜别克族","锡伯族","瑶族","彝族","仡佬族","裕固族","壮族"}
	for k,_ := range mingzu{

		index := ""
		if k < 9{
			index = "0"+fmt.Sprint(k+1)
		}else{
			index = fmt.Sprint(k+1)
		}
		imgRes,_ := http.Get("http://www.sinobuy.cn/mzfz/img/0"+index+".jpg")
		defer imgRes.Body.Close()
		imgByte,_ := ioutil.ReadAll(imgRes.Body)
		path := utils.GetCurrentDir("/static/mingzu/img/")
		fn := path+fmt.Sprint(k)+".jpg"
		_,fErr := os.Stat(fn)
		var fh *os.File
		if fErr != nil{
			fh,_ = os.Create(fn)
		}else{
			fh,_ = os.Open(fn)
		}
		defer fh.Close()
		fh.Write(imgByte)
	}
}

func (this *AdminController) Add() {
	this.pageTitle("新增管理员")
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/add")
}

func (this *AdminController) Edit() {
	this.pageTitle("编辑管理员")
	id := this.getInt("id")
	admin, _ := models.AdminGetById(id)
	row := make(map[string]interface{})
	row["id"] = admin.Id
	row["login_name"] = admin.Name
	row["real_name"] = admin.RealName
	row["phone"] = admin.Phone
	row["email"] = admin.Email
	row["role_ids"] = admin.RoleIds
	this.Data["admin"] = row

	role_ids := strings.Split(admin.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(role_ids); i++ {
			role_id, _ := strconv.Atoi(role_ids[i])
			if role_id == v.Id {
				row["checked"] = 1
			}
		}
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/edit")
}

func (this AdminController) AjaxSave() {
	id := this.getInt("id")
	if id == 0 {
		admin := new(models.Admin)
		this.getParam(admin)
		admin.Salt = utils.GetRandomString(10)
		admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
		admin.CreateTime = time.Now().Unix()
		admin.CreateId = this.userId
		if _, err := models.AdminAdd(admin); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	admin, _ := models.AdminGetById(id)
	admin.Id = id
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	this.getParam(admin)

	resetPsw := this.getInt("reset_pwd")
	if resetPsw == 1 {
		admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
		admin.Salt = utils.GetRandomString(10)
	}
	if err := admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg(strconv.Itoa(resetPsw), MSG_OK)
}

func (this AdminController) getParam(admin *models.Admin) {
	admin.Name = this.getString("login_name", "账号不能为空!", defaultMinSize)
	admin.RealName = this.getString("real_name", "真实名字不能为空!", defaultMinSize)
	admin.Phone = this.getString("phone", "联系电话不能为空!", defaultMinSize)
	admin.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
	admin.RoleIds = this.getString("roleids", "没有选择权限项!", 0)
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	admin.Status = 1
}

func (this AdminController) AjaxDel(){
	id := this.getInt("id")
	beego.Info("---------------id:",id)
	admin,_ := models.AdminGetById(id)
	admin.UpdateTime = time.Now().Unix()
	admin.Status = 0
	admin.Id = id
	if id == 1{
		this.ajaxMsg("超级管理员不允许删除!",MSG_ERR)
	}
	if err := admin.Update();err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	this.ajaxMsg("",MSG_OK)
}

func (this *AdminController) Table() {
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	this.pageSize = limit
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AdminList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.Name
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
