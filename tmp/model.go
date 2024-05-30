package tmp

func GetModelSub() string {

	return `package models

import "github.com/sonhineboy/gsadmin/service/global"

type {{.Name | Title}} struct {
	*global.GAD_MODEL
	{{range .Fields}}{{ . | TransFieldAll}}
	{{end}}
}

func (m *{{.Name | Title}}) TableName() string {
	return "{{.Name}}"
}
`
}
