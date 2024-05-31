package svr

func GetRequestSub() string {

	return `package requests

type Delete{{.Name | Title}}Request struct {
	Ids []int {{Del}}
}

type {{.Name | Title}}Request struct {
	{{range .Fields}}{{ . | Trans}}
	{{end}}
}`
}
