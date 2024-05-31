package gsadminGen

import (
	"fmt"
	"github.com/sonhineboy/gsadminGen"
	"os"
	"testing"
)

func TestTitle(t *testing.T) {

	fmt.Println(fmt.Sprintf("%-5saaaa", "4563asdfasf"))

}

func TestGenModel(t *testing.T) {
	fields := []gsadminGen.Field{
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
	err := gsadminGen.GenModel("./model.text", gsadminGen.TableModal{
		Name:   "test",
		Fields: fields,
	})
	if err != nil {
		fmt.Println("GenModel Error:", err)
	}
}

func TestGenController(t *testing.T) {
	fields := []gsadminGen.Field{
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
	err := gsadminGen.GenController("./controller.text", gsadminGen.TableModal{
		Name:   "user",
		Fields: fields,
	})
	if err != nil {
		t.Error("GenController Error:", err)
	}
}

func TestRequest(t *testing.T) {

	fields := []gsadminGen.Field{
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

	err := gsadminGen.GenRequest("./request.text", gsadminGen.TableModal{
		Name:   "user",
		Fields: fields,
	})
	if err != nil {
		t.Error("GenController Error:", err)
	}
}
func TestRepository(t *testing.T) {

	fields := []gsadminGen.Field{
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

	err := gsadminGen.GenRepository("./repository", gsadminGen.TableModal{
		Name:   "user",
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

	fields := []gsadminGen.Field{
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

	for _, g := range fields {

		fmt.Println(gsadminGen.TransFieldAll(g))

	}
}
