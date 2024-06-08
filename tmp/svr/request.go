package svr

func GetRequestSub() string {

	return `package requests

type Delete{{.Name | Transform}}Request struct {
	Ids []int {{Del}}
}

type {{.Name | Transform}}Request struct {
	{{range .Fields}}{{ . | Trans}}
	{{end}}
}`
}
