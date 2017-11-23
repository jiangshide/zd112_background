package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"bytes"
	"strconv"
	"github.com/astaxie/beego"
)

type RoleAuth struct {
	AuthId int `orm:"pk"`
	RoleId int64
}

func (this *RoleAuth) TableName() string {
	return TableName("uc_role_auth")
}

func RoleAuthAdd(this *RoleAuth) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func RoleAuthGetById(id int) ([]*RoleAuth, error) {
	list := make([]*RoleAuth, 0)
	_, err := orm.NewOrm().QueryTable(TableName("uc_role_auth")).Filter("role_id", id).All(&list, "AuthId")
	return list, err
}

func RoleAuthDelete(id int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("uc_role_auth")).Filter("role_id", id).Delete()
}

func RoleAuthGetByIds(roleIds string) (string, error) {
	list := make([]*RoleAuth, 0)
	ids := strings.Split(roleIds, ",")
	_, err := orm.NewOrm().QueryTable(TableName("uc_role_auth")).Filter("role_id_in", ids).All(&list, "AuthId")
	beego.Error("------------err:",err)
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	return strings.TrimRight(b.String(), ","), nil
}

func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
	i, _ := orm.NewOrm().QueryTable(TableName("uc_role_auth")).PrepareInsert()
	for _, ra := range ras {
		if _, err := i.Insert(ra); err == nil {
			n += 1
		}
	}
	i.Close()
	return n, err
}
