package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	contextNotErr       = errors.New("context is empty ")
	routersNotErr       = errors.New("routers is empty ")
	routerFlagNotHasErr = errors.New("router flag not has ")
	routerHasErr        = errors.New("writer router is has ")
)

type WriterRouter struct {
	Slices
	Path         string
	RouterFlag   string
	PkgName      string
	context      []byte
	contextLines []string
	routers      []string
}

func NewWriterRouter(path, routerFlag, pkgName string) *WriterRouter {
	return &WriterRouter{
		Path:       path,
		RouterFlag: routerFlag,
		PkgName:    pkgName,
	}
}

func (w *WriterRouter) Write(routers []string) error {

	err := w.setFileContext()
	if err != nil {
		return fmt.Errorf("setFileContext err %v", err)
	}

	err = w.transFileContext()
	if err != nil {
		return fmt.Errorf("transFileContext err %v", err)
	}

	err = w.writerPkg()
	if err != nil {
		return fmt.Errorf("writerPkg err %v", err)
	}

	w.routers = routers
	err = w.writerRouter()

	if err != nil {
		return fmt.Errorf("writerRouter err %v", err)
	}

	str := strings.Join(w.contextLines, "\n")
	err = ioutil.WriteFile(w.Path, []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil

}

// writerPkg 把导入包写入
func (w *WriterRouter) writerPkg() error {
	isHas, err := w.hasPkg()
	if err != nil {
		return fmt.Errorf("writerPkg %v", err)
	}

	if !isHas {
		w.contextLines = w.append(w.contextLines, 3, []string{w.getAllPkg()})
	}
	return nil
}

func (w *WriterRouter) getAllPkg() string {
	return fmt.Sprintf("\t\"github.com/sonhineboy/gsadmin/service/app/controllers/%s\"", w.PkgName)
}

func (w *WriterRouter) hasPkg() (bool, error) {
	if w.context == nil {
		return false, contextNotErr
	}
	return strings.Contains(string(w.context), w.getAllPkg()), nil
}

//writerRouter 路由内容导入
func (w *WriterRouter) writerRouter() error {
	if w.routers == nil {
		return routersNotErr
	}

	index := w.SliceIndex(w.contextLines, w.RouterFlag)

	if index == -1 {
		return fmt.Errorf("%v flag:%s", routerFlagNotHasErr, w.RouterFlag)
	}

	if w.SliceIndex(w.contextLines, w.routers[1]) != -1 {
		return fmt.Errorf("%v flag:%s", routerHasErr, w.routers[1])
	}

	w.contextLines = w.append(w.contextLines, index, w.routers)
	return nil

}

// setFileContext 读取文件内容
func (w *WriterRouter) setFileContext() (err error) {
	if w.context == nil {
		w.context, err = ioutil.ReadFile(w.Path)
		if err != nil {
			return fmt.Errorf("writeRouter setFileContext %v", err)
		}
	}
	return err
}

//transFileContext 把文件内容分割到切片
func (w *WriterRouter) transFileContext() (err error) {
	if w.contextLines == nil {
		if w.context == nil {
			err = w.setFileContext()
			if err != nil {
				return fmt.Errorf("writeRouter transFileContext %v", err)
			}
		}
		w.contextLines = strings.Split(string(w.context), "\n")
	}
	return err
}
