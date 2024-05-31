package gsadminGen

import (
	"fmt"
	"github.com/sonhineboy/gsadminGen/src"
	"github.com/sonhineboy/gsadminGen/tmp/svr"
	"os"
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
func GenModel(fileName string, v TableModal) error {
	myFunc := template.FuncMap{
		"Title":         strings.Title,
		"TransFieldAll": TransFieldAll,
	}
	tmpl, err := template.New("model.sub").Funcs(myFunc).Parse(svr.GetModelSub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}
	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenController(fileName string, v TableModal) error {
	myFunc := template.FuncMap{
		"Title": strings.Title,
	}
	tmpl, err := template.New("controller.sub").Funcs(myFunc).Parse(svr.GetControllerSub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}
	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenRepository(fileName string, v TableModal) error {
	myFunc := template.FuncMap{
		"Title": strings.Title,
	}
	tmpl, err := template.New("repository.sub").Funcs(myFunc).Parse(svr.GetRepositorySub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}
	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenRequest(fileName string, v TableModal) error {
	myFunc := template.FuncMap{
		"Title": strings.Title,
		"Del":   TransDelRequest,
		"Trans": TransRequest,
	}
	tmpl, err := template.New("request.sub").Funcs(myFunc).Parse(svr.GetRequestSub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
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

func TransRequest(field Field) string {
	fieldType, ok := src.FieldTypeMapping[field.Type]
	if !ok {
		fmt.Println("类型mapping 不存在")
		return ""
	}

	isNull := " binding:\"required\""
	if field.IsNull {
		isNull = ""
	}

	return fmt.Sprint(
		fmt.Sprintf("%-15s", strings.Title(field.Name)),
		fmt.Sprintf("%-10s", fieldType),
		"`",
		"json:\"",
		field.Json,
		"\"",
		isNull, "`",
	)
}

func TransDelRequest() string {
	return "`binding:\"required\"`"
}

// TransFieldAll 字段转换
func TransFieldAll(field Field) string {
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

	isNull := "not null"
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
		index = fmt.Sprint("index:", fieldName, ",class:fulltext;")
		break
	}

	autoInc := ""
	if strings.ToLower(field.Name) == "id" {
		autoInc = "autoIncrement;"
	}

	fieldSlice := []string{
		fmt.Sprintf("%-15s", strings.Title(field.Name)),
		fmt.Sprintf("%-10s", fieldType),
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
		" json:\"",
		field.Json,
		"\"",
		"`",
	}

	return strings.Join(fieldSlice, "")

}
