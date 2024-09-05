package svr

func GetModelSub() string {

	return `package models

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/global"
)

type {{.Name | Transform}} struct {
	global.GAD_MODEL
	{{range .Fields}}{{ . | TransFieldAll}}
	{{end}}
}

func (m *{{.Name | Transform}}) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "{{.Name}}")
}
`
}
