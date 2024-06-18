package gsadminGen

import (
	"fmt"
	"github.com/sonhineboy/gsadminGen"
	"github.com/sonhineboy/gsadminGen/pkg"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var MyData = pkg.TableModal{Name: "userMember", Fields: []pkg.Field{
	{
		Name:     "userName",
		Json:     "user_name",
		Default:  "",
		Describe: "用户名",
		Primary:  false,
		Index:    "FULLTEXT",
		IsNull:   true,
		Type:     "varchar",
		Transfer: "用户名",
	},
	{
		Name:     "age",
		Json:     "age",
		Default:  "0",
		Describe: "年龄",
		Primary:  false,
		Index:    "Null",
		IsNull:   true,
		Type:     "int",
		Transfer: "年龄",
	},
}}

func TestTitle(t *testing.T) {

	fmt.Println(fmt.Sprintf("%-5saaaa", "4563asdfasf"))

}

func TestGenModel(t *testing.T) {
	fields := []pkg.Field{
		{
			Name:     "id",
			Json:     "id",
			Default:  "",
			Describe: "Id",
			Primary:  true,
			Index:    "Null",
			IsNull:   true,
			Type:     "int",
			Transfer: "Id",
		},
		{
			Name:     "userName",
			Json:     "user_name",
			Default:  "",
			Describe: "用户名",
			Primary:  false,
			Index:    "FULLTEXT",
			IsNull:   true,
			Type:     "varchar",
			Transfer: "用户名",
		},
		{
			Name:     "age",
			Json:     "age",
			Default:  "0",
			Describe: "年龄",
			Primary:  false,
			Index:    "Null",
			IsNull:   true,
			Type:     "int",
			Transfer: "年龄",
		},
	}
	err := gsadminGen.GenModel("./model.text", pkg.TableModal{
		Name:   "test",
		Fields: fields,
	})
	if err != nil {
		fmt.Println("GenModel Error:", err)
	}
}

func TestGenController(t *testing.T) {
	fields := []pkg.Field{
		{
			Name:     "userName",
			Json:     "user_name",
			Default:  "",
			Describe: "用户名",
			Primary:  false,
			Index:    "FULLTEXT",
			IsNull:   true,
			Type:     "varchar",
			Transfer: "用户名",
		},
		{
			Name:     "age",
			Json:     "age",
			Default:  "0",
			Describe: "年龄",
			Primary:  false,
			Index:    "Null",
			IsNull:   true,
			Type:     "int",
			Transfer: "年龄",
		},
	}
	err := gsadminGen.GenController("./sss/controllexxx.text", pkg.TableModal{
		Name:   "user",
		Fields: fields,
	}, "txss")
	if err != nil {
		t.Error("GenController Error:", err)
	}
}

func TestRequest(t *testing.T) {

	fields := []pkg.Field{
		{
			Name:     "userName",
			Json:     "user_name",
			Default:  "",
			Describe: "用户名",
			Primary:  false,
			Index:    "FULLTEXT",
			IsNull:   true,
			Type:     "varchar",
			Transfer: "用户名",
		},
		{
			Name:     "age",
			Json:     "age",
			Default:  "0",
			Describe: "年龄",
			Primary:  false,
			Index:    "Null",
			IsNull:   false,
			Type:     "int",
			Transfer: "年龄",
		},
	}

	err := gsadminGen.GenRequest("./request.text", pkg.TableModal{
		Name:   "user",
		Fields: fields,
	})
	if err != nil {
		t.Error("GenController Error:", err)
	}
}
func TestRepository(t *testing.T) {

	fields := []pkg.Field{
		{
			Name:     "user_name",
			Json:     "user_name",
			Default:  "",
			Describe: "用户名",
			Primary:  false,
			Index:    "FULLTEXT",
			IsNull:   true,
			Type:     "varchar",
			Transfer: "用户名",
		},
		{
			Name:     "age",
			Json:     "age",
			Default:  "0",
			Describe: "年龄",
			Primary:  false,
			Index:    "Null",
			IsNull:   false,
			Type:     "int",
			Transfer: "年龄",
		},
	}

	err := gsadminGen.GenRepository("./repository", pkg.TableModal{
		Name:   "user_member",
		Fields: fields,
	})
	if err != nil {
		t.Error("GenController Error:", err)
	}
}

func TestCreateFile(t *testing.T) {
	var err error
	f, e := os.Create("./hello.txt")
	if e != nil {
		fmt.Println("文件创建失败:", e)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = f.Write([]byte("hello"))
	if err != nil {
		return
	}

}

func TestMapping(t *testing.T) {

	fields := []pkg.Field{
		{
			Name:     "user_name",
			Json:     "user_name",
			Default:  "",
			Describe: "用户名",
			Primary:  false,
			Index:    "FULLTEXT",
			IsNull:   true,
			Type:     "varchar",
			Transfer: "用户名",
		},
		{
			Name:     "age",
			Json:     "age",
			Default:  "0",
			Describe: "年龄",
			Primary:  false,
			Index:    "Null",
			IsNull:   true,
			Type:     "int",
			Transfer: "年龄",
		},
	}

	for _, g := range fields {

		fmt.Println(gsadminGen.TransFieldAll(g))

	}
}

func TestDir(t *testing.T) {
	path := "./aaa/ttt.txt"
	fileInfo, err := os.Stat(filepath.Dir(path))

	fmt.Println(err)

	fmt.Printf("%v", fileInfo.IsDir())
}
func TestStringC(t *testing.T) {
	s := "user_name"
	matchRe := regexp.MustCompile("[-_]").FindString(s)

	fmt.Println(matchRe)
	if len(matchRe) > 0 {
		strSlice := strings.Split(s, matchRe)
		var str strings.Builder
		for _, s2 := range strSlice {
			str.Write([]byte(strings.Title(s2)))
		}
		fmt.Println(str.String())
	}

	fmt.Println(strings.Title(s))

	fmt.Println(gsadminGen.UnderToConvertSoreLow(s))
}

func TestDefer(t *testing.T) {

	func() {
		defer func() {
			fmt.Println("-----")
		}()
	}()

	defer func() {
		fmt.Println(2222)
	}()
}

func TestGenJs(t *testing.T) {

	err := gsadminGen.GenApi("./userMember.js", MyData)
	if err != nil {

		t.Error(err)
		return
	}
}

func TestGenForm(t *testing.T) {
	err := gsadminGen.GenForm("./form.vue", MyData)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGenIndex(t *testing.T) {
	err := gsadminGen.GenIndex("./index.vue", MyData)
	if err != nil {
		t.Error(err)
		return
	}
}
