package gsadminGen

import (
	"fmt"
	"github.com/sonhineboy/gsadminGen/src"
	"github.com/sonhineboy/gsadminGen/tmp"
	"io"
	"regexp"
	"strings"
	"text/template"
)

type Field struct {
	Name     string
	Json     string
	Default  string
	Describe string
	Primary  bool
	Index    string
	IsNull   bool
	Type     string
	Transfer string
}

type TableModal struct {
	Name   string
	Fields []Field
}

// GenModel 生成模型结构体
func GenModel(wr io.Writer, v TableModal) error {

	MyFunc := template.FuncMap{
		"Title":         strings.Title,
		"TransFieldAll": TransFieldAll,
	}
	tmpl, err := template.New("model.sub").Funcs(MyFunc).Parse(tmp.GetSub())
	if err != nil {
		return err
	}
	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

// ConvertToUnderScore 将驼峰命名转换为下划线命名
func ConvertToUnderScore(camelCaseName string) string {
	reg := regexp.MustCompile("([a-z0-9A-Z])([A-Z])")
	underScoreName := reg.ReplaceAllString(camelCaseName, "${1}_${2}")
	underScoreName = strings.ToLower(underScoreName)
	return underScoreName
}

// TransFieldAll 字段转换
func TransFieldAll(field Field, maxLen int) string {
	separator := "    "

	fieldName := ConvertToUnderScore(field.Name)

	fieldType, ok := src.FieldTypeMapping[field.Type]
	if !ok {
		fmt.Println("类型mapping 不存在")
		return ""
	}

	primary := ""
	if field.Primary {
		primary = "primaryKey;"
	}

	isNull := ""
	if field.IsNull {
		isNull = ""
	}

	comment := ""
	if len(field.Describe) > 0 {
		comment = fmt.Sprint("comment:", field.Describe, ";")
	}

	defaultVal := ""
	if len(field.Default) > 0 {
		defaultVal = fmt.Sprint("default:", field.Default, ";")
	}

	index := ""
	switch field.Index {
	case "NORMAL":
		index = fmt.Sprint("index:", fieldName, ";")
		break
	case "UNIQUE":
		index = fmt.Sprint("index:", fieldName, ",unique;")
		break
	case "FULLTEXT":
		index = fmt.Sprint("index:", fieldName, ",fulltext;")
		break
	}

	autoInc := ""
	if strings.ToLower(field.Name) == "id" {
		autoInc = "autoIncrement;"
	}

	fieldSlice := []string{
		strings.Title(field.Name),
		separator,
		fieldType,
		separator,
		"`gorm:\"column:",
		fieldName,
		";",
		primary,
		autoInc,
		index,
		isNull,
		defaultVal,
		comment,
		"\"",
		" json:",
		field.Json,
		"`",
	}

	return strings.Join(fieldSlice, "")

}
