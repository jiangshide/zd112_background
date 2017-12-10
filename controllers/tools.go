package controllers

import (
	"zd112/models"
	"os"
	"github.com/astaxie/beego"
	"time"
	"bufio"
	"io"
	"strings"
)

type FormatTypeController struct {
	BaseController
}

func (this *FormatTypeController) List() {
	this.pageTitle("文件类型列表")
	this.display(this.getBgToolAction("format/type/list"))
}

func (this *FormatTypeController) Add() {
	this.pageTitle("增加文件类型")
	this.display(this.getBgToolAction("format/type/add"))
}

func (this *FormatTypeController) Edit() {
	formatType := new(models.FormatType)
	formatType.Id = this.getInt("id", 0)
	if err := formatType.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = formatType.Id
	row["name"] = formatType.Name
	this.Data["row"] = row
	this.display(this.getBgToolAction("format/type/edit"))
}

func (this *FormatTypeController) AjaxSave() {
	formatType := new(models.FormatType)
	formatType.Id = this.getInt("id", 0)
	if formatType.Id == 0 {
		formatType.Name = this.getString("name", "名称不能为空!", 1)
		formatType.CreateId = this.userId
		formatType.CreateTime = time.Now().Unix()
		if _, err := formatType.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	formatType.Query()
	formatType.Name = this.getString("name", "名称不能为空!", 1)
	formatType.UpdateId = this.userId
	formatType.UpdateTime = time.Now().Unix()
	if _, err := formatType.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatTypeController) Table() {
	formatType := new(models.FormatType)
	result, count := formatType.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *FormatTypeController) AjaxDel() {
	formatType := new(models.FormatType)
	formatType.Id = this.getInt("id", 0)
	if _, err := formatType.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type FormatController struct {
	BaseController
}

func (this *FormatController) List() {
	this.pageTitle("文件格式列表")
	this.display(this.getBgToolAction("format/list"))
}

func (this *FormatController) Add() {
	this.pageTitle("增加文件格式")
	this.parent(0)
	this.display(this.getBgToolAction("format/add"))
}

func (this *FormatController) Edit() {
	this.pageTitle("编辑文件格式")
	format := new(models.Format)
	format.Id = this.getInt("id", 0)
	if err := format.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	this.parent(format.ParentId)
	row["id"] = format.Id
	row["name"] = format.Name
	this.Data["row"] = row
	this.display(this.getBgToolAction("format/edit"))
}

func (this *FormatController) AjaxSave() {
	format := new(models.Format)
	format.ParentId = this.getInt("group_id", 0)
	format.Name = this.getString("name", "名称不能为空!", 1)
	format.Id = this.getInt("id", 0)
	if format.Id == 0 {
		format.CreateId = this.userId
		format.UpdateTime = time.Now().Unix()
		if _, err := format.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	format.UpdateId = this.userId
	format.UpdateTime = time.Now().Unix()
	if _, err := format.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatController) parent(id int) {
	formatType := new(models.FormatType)
	result, count := formatType.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *FormatController) Table() {
	format := new(models.Format)
	result, count := format.List(this.pageSize, this.offSet)
	if count == 0 {
		this.Export()
	}
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		formatType := new(models.FormatType)
		formatType.Id = v.ParentId

		if err := formatType.Query(); err == nil {
			row["parent"] = formatType.Name
		} else {
			beego.Error(err)
		}
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *FormatController) AjaxDel() {
	format := new(models.Format)
	format.Id = this.getInt("id", 0)
	if _, err := format.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatController) Export() {
	f, _ := os.Open("format")
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, ":") {
			lineStr := strings.Split(line, ":")
			if lineStr[0] == "video" {
				this.addDb("视频", lineStr[1])
			} else if lineStr[0] == "audio" {
				this.addDb("音频", lineStr[1])
			} else if lineStr[0] == "picture" {
				this.addDb("图片", lineStr[1])
			} else if lineStr[0] == "text" {
				this.addDb("文字", lineStr[1])
			}
		}
	}
}

func (this *FormatController) addDb(name, format string, ) {
	formatType := new(models.FormatType)
	formatType.Name = name
	formatType.CreateId = this.userId
	formatType.CreateTime = time.Now().Unix()
	if _, err := formatType.Add(); err != nil {
		beego.Error(err)
	}
	formatType.Query()
	if formatType.Id > 0 {
		formatStr := strings.Split(format, ".")
		for _, v := range formatStr {
			formatTable := new(models.Format)
			formatTable.ParentId = formatType.Id
			formatTable.Name = v;
			formatTable.CreateId = this.userId
			formatTable.CreateTime = time.Now().Unix()
			if _, err := formatTable.Add(); err != nil {
				beego.Error(err)
			}
		}
	}
}

type CompressController struct {
	BaseController
}

func (this *CompressController) List() {
	this.pageTitle("压缩文件列表")
	this.display(this.getBgToolAction("compress/list"))
}

func (this *CompressController) Add() {
	this.pageTitle("添加文件")
	this.display(this.getBgToolAction("compress/add"))
}

func (this *CompressController) Edit() {
	this.pageTitle("编辑文件")
	compress := new(models.Compress)
	compress.Id = this.getInt("id", 0)
	if err := compress.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = compress.Id
	row["name"] = compress.Name
	row["file"]=compress.Url
	row["format"] = compress.Format
	row["size"] = compress.Size
	row["resize"] = compress.ReSize
	row["compress"] = compress.Compress
	row["downs"] = compress.Downs
	row["descript"] = compress.Descript
	this.Data["row"] = row
	this.display(this.getBgToolAction("compress/edit"))
}

func (this *CompressController) AjaxSave() {
	compress := new(models.Compress)
	compress.Id = this.getInt("id", 0)
	compress.Name = this.getString("name", "名称不能为空!", 1)
	compress.Url = this.getString("file", "上传文件URL为空!", defaultMinSize)
	compress.Descript = this.getString("descript", "", 0)
	compress.Size = this.getInt64("size", 0)
	compress.ReSize = this.getInt64("resize", 0)
	compress.Compress = this.getInt("compress", 0)
	if compress.ReSize > 0 {
		compress.Compress = 1
	}
	if compress.Id == 0 {
		id, reSize, sufix := this.getFileFormat(compress.Url)
		compress.Type = id
		compress.Format = sufix
		compress.ReSize = reSize
		compress.CreateId = this.userId
		compress.CreateTime = time.Now().Unix()
		if _, err := compress.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	compress.Downs = this.getInt("downs", 0)
	compress.UpdateId = this.userId
	compress.UpdateTime = time.Now().Unix()
	if _, err := compress.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CompressController) Table() {
	compress := new(models.Compress)
	result, count := compress.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Url
		row["parent"] = this.getFileType(v.Url)
		row["format"] = v.Format
		row["size"] = v.Size
		row["resize"] = v.ReSize
		row["compress"] = v.Compress
		row["downs"] = v.Downs
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CompressController) AjaxDel() {
	compress := new(models.Compress)
	compress.Id = this.getInt("id", 0)
	if _, err := compress.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
