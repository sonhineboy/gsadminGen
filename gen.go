package gsadminGen

import (
	"fmt"
	"github.com/sonhineboy/gsadminGen/pkg"
	"github.com/sonhineboy/gsadminGen/tmp/svr"
	"github.com/sonhineboy/gsadminGen/tmp/web"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// GenModel 生成模型结构体
func GenModel(fileName string, v pkg.TableModal) error {
	myFunc := template.FuncMap{
		"Title":         strings.Title,
		"TransFieldAll": TransFieldAll,
		"Transform":     UnderToConvertSore,
	}
	tmpl, err := template.New("model.sub").Funcs(myFunc).Parse(svr.GetModelSub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)
	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenIndex(fileName string, v pkg.TableModal) error {
	myFunc := template.FuncMap{
		"UnderToConvertSoreLow": UnderToConvertSoreLow,
	}
	tmpl, err := template.New("api.sub").Funcs(myFunc).Parse(web.GetIndexSub())
	if err != nil {
		return err
	}

	if err = mkDirByFile(fileName); err != nil {
		return err
	}

	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenForm(fileName string, v pkg.TableModal) error {
	myFunc := template.FuncMap{
		"TransInt": TranInt,
		"HasStr":   HasStr,
	}
	tmpl, err := template.New("api.sub").Funcs(myFunc).Parse(web.GetFormSub())
	if err != nil {
		return err
	}

	if err = mkDirByFile(fileName); err != nil {
		return err
	}

	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

func GenApi(fileName string, v pkg.TableModal) error {

	myFunc := template.FuncMap{
		"GetUrl": GetUrl,
	}

	tmpl, err := template.New("api.sub").Funcs(myFunc).Parse(web.GetApiJsSub())
	if err != nil {
		return err
	}

	if err = mkDirByFile(fileName); err != nil {
		return err
	}

	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

//mkDirByFile 根据文件创建文件
func mkDirByFile(fileName string) error {
	dir := filepath.Dir(fileName)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

//GenController 生成控制器
func GenController(fileName string, v pkg.TableModal, pkgName string) error {
	myFunc := template.FuncMap{
		"Title":     strings.Title,
		"Transform": UnderToConvertSore,
	}
	tmpl, err := template.New("controller.sub").Funcs(myFunc).Parse(svr.GetControllerSub())
	if err != nil {
		return err
	}

	if err = mkDirByFile(fileName); err != nil {
		return err
	}

	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}
	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

	controllerVar := pkg.ControllerVar{TableModal: v, Pkg: pkgName}

	err = tmpl.Execute(wr, controllerVar)
	if err != nil {
		return err
	}
	return nil
}

//GenRepository 生成业务仓储
func GenRepository(fileName string, v pkg.TableModal) error {
	myFunc := template.FuncMap{
		"Title":     strings.Title,
		"Transform": UnderToConvertSore,
	}
	tmpl, err := template.New("repository.sub").Funcs(myFunc).Parse(svr.GetRepositorySub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

	err = tmpl.Execute(wr, v)
	if err != nil {
		return err
	}
	return nil
}

//GenRequest 生成请求结构体
func GenRequest(fileName string, v pkg.TableModal) error {
	myFunc := template.FuncMap{
		"Title":     strings.Title,
		"Del":       TransDelRequest,
		"Trans":     TransRequest,
		"Transform": UnderToConvertSore,
	}
	tmpl, err := template.New("request.sub").Funcs(myFunc).Parse(svr.GetRequestSub())
	if err != nil {
		return err
	}
	wr, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	defer func(wr *os.File) {
		_ = wr.Close()
	}(wr)

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

//UnderToConvertSore 下划线转大驼峰
func UnderToConvertSore(s string) string {
	matchRe := regexp.MustCompile("[-_]").FindString(s)
	if len(matchRe) > 0 {
		strSlice := strings.Split(s, matchRe)
		var str strings.Builder
		for _, s2 := range strSlice {
			str.Write([]byte(strings.Title(s2)))
		}
		return str.String()
	} else {
		return strings.Title(s)
	}
}

//UnderToConvertSoreLow 下划线转小驼峰
func UnderToConvertSoreLow(s string) string {
	s = UnderToConvertSore(s)
	return strings.ToLower(s[0:1]) + s[1:]
}

//TransRequest trans request
func TransRequest(field pkg.Field) string {
	fieldType, ok := pkg.FieldTypeMapping[field.Type]
	if !ok {
		fmt.Println("类型mapping 不存在")
		return ""
	}

	isNull := " binding:\"required\""
	if field.IsNull {
		isNull = ""
	}

	return fmt.Sprint(
		fmt.Sprintf("%-15s", UnderToConvertSore(field.Name)),
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
func TransFieldAll(field pkg.Field) string {
	fieldName := ConvertToUnderScore(field.Name)

	fieldType, ok := pkg.FieldTypeMapping[field.Type]
	if !ok {
		fmt.Println("类型mapping 不存在")
		return ""
	}

	fieldDbType, ok := pkg.FieldDbMapping[field.Type]
	if ok {
		fieldDbType = fmt.Sprint("type:", fieldDbType, ";")
	} else {
		fieldDbType = ""
	}

	primary := ""
	if field.Primary {
		primary = "primaryKey;"
	}

	isNull := "not null;"
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
	indexPrefix := "index:"
	switch field.Index {
	case "NORMAL":
		index = fmt.Sprint(indexPrefix, fieldName, ";")
		break
	case "UNIQUE":
		index = fmt.Sprint(indexPrefix, fieldName, ",class:unique;")
		break
	case "FULLTEXT":
		index = fmt.Sprint(indexPrefix, fieldName, ",class:fulltext;")
		break
	}

	autoInc := ""
	if strings.ToLower(field.Name) == "id" {
		autoInc = "autoIncrement;"
	}

	fieldSlice := []string{
		fmt.Sprintf("%-15s", UnderToConvertSore(field.Name)),
		fmt.Sprintf("%-10s", fieldType),
		"`gorm:\"column:",
		fieldName,
		";",
		fieldDbType,
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

func GetUrl(name string, action string) string {
	return fmt.Sprint("`${config.API_URL}/", UnderToConvertSoreLow(name), "/", action, "`")
}

func TranInt(str string, types string) string {
	if strings.Contains(types, "int") {
		return str
	} else {
		return fmt.Sprint("\"", str, "\"")
	}
}

func HasStr(str string, sub string) bool {
	return strings.Contains(str, sub)
}
