package controllers

import (
	"zd112/models"
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"zd112/utils"
	"time"
)

type NationController struct {
	BaseController
}

func (this *NationController) List() {
	this.pageTitle("民族列表")
	this.display("backstage/nation/list")
}

func (this *NationController) Add() {
	this.pageTitle("新增民族名称")
	nation, err := models.NationList(1, 1000)
	beego.Info("-----------nation:", nation, " | err:", err)
	this.Data["nation"] = nation
	this.display("backstage/nation/add")
}

func (this *NationController) Edit() {
	this.pageTitle("编辑民族名称")
	if nation, err := models.NationById(this.getInt("id", 0)); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		this.Data["nation"] = nation
	}
	this.display("backstage/nation/edit")
}

func (this *NationController) AjaxSave() {
	id := this.getInt("id", 0)
	if id == 0 {
		nation := new(models.Nation)
		nation.Name = this.getString("name", "名称不能为空!", 1)
		nation.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		beego.Info("--------------name:", nation.Name, " | icon:", nation.Icon)
		nation.CreateId = this.userId
		nation.CreateTime = time.Now().Unix()
		if _, err := models.NationAdd(nation); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	nation, _ := models.NationById(id)
	nation.Name = this.getString("name", "名称不能为空!", 1)
	nation.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	nation.UpdateId = this.userId
	nation.UpdateTime = time.Now().Unix()
	if _, err := nation.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *NationController) Upload() {
	beego.Info("--------------upload")
	f, fh, err := this.GetFile("file")
	beego.Error("----------err:", err, " | fh:", fh.Filename)
	defer f.Close()
	this.SaveToFile("file", "static/upload/"+fh.Filename)
	this.ajaxMsg("成功!", MSG_OK)
}

func (this *NationController) Table() {
	this.pageSize = this.getInt("limit", 30)
	result, count := models.NationList(this.getInt("page", 1), this.pageSize)
	beego.Info("---------result:", result, " | count:", count)
	this.ajaxList("成功", MSG_OK, count, result)
}

func download() {
	mingzu := [...]string{"阿昌族", "白族", "保安族", "布朗族", "布依族", "藏族", "朝鲜族", "达翰尔族", "傣族", "昂德族", "东乡族", "侗族", "独龙族", "俄罗斯族", "鄂伦春族", "鄂温克族", "高山族", "哈尼族", "哈萨克族", "汉族", "赫哲族", "回族", "基诺族", "京族",
		"景颇族", "柯尔克孜族", "拉祜族", "黎族", "傈僳族", "珞巴族", "满族", "毛南族", "门巴族", "蒙古族", "苗族", "仫佬族", "纳西族", "怒族", "普米族", "羌族", "撒拉族", "畲族", "水族", "塔吉克族", "塔塔尔族", "土家族", "图族", "佤族", "维吾尔族", "乌孜别克族", "锡伯族", "瑶族", "彝族", "仡佬族", "裕固族", "壮族"}
	for k, _ := range mingzu {
		index := ""
		if k < 9 {
			index = "0" + fmt.Sprint(k+1)
		} else {
			index = fmt.Sprint(k + 1)
		}
		imgRes, _ := http.Get("http://www.sinobuy.cn/mzfz/img/0" + index + ".jpg")
		defer imgRes.Body.Close()
		imgByte, _ := ioutil.ReadAll(imgRes.Body)
		path := utils.GetCurrentDir("/static/mingzu/img/")
		fn := path + fmt.Sprint(k) + ".jpg"
		_, fErr := os.Stat(fn)
		var fh *os.File
		if fErr != nil {
			fh, _ = os.Create(fn)
		} else {
			fh, _ = os.Open(fn)
		}
		defer fh.Close()
		fh.Write(imgByte)
	}
}
