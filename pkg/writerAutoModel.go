package pkg

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type WriterAutoModel struct {
	Slices
	File         string
	Content      string
	ContentSlice []string
	Flag         string
}

func NewWriterAutoModel(file, flag string) *WriterAutoModel {
	return &WriterAutoModel{
		File: file,
		Flag: flag,
	}
}

func (w *WriterAutoModel) Write(str []string) error {
	err := w.ReadFile()
	if err != nil {
		return fmt.Errorf("read %v", err)
	}
	w.TransSlice()
	index := w.SliceIndex(w.ContentSlice, w.Flag)
	if index == -1 {
		return fmt.Errorf("not flag   %s ", w.Flag)
	}

	w.ContentSlice = w.append(w.ContentSlice, index, str)
	content := strings.Join(w.ContentSlice, "\n")

	err = ioutil.WriteFile(w.File, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Write   %v ", err)
	}
	return nil
}

func (w *WriterAutoModel) ReadFile() error {
	file, err := ioutil.ReadFile(w.File)
	if err != nil {
		return fmt.Errorf("ReadFile %v", err)
	}
	w.Content = string(file)
	return nil
}

func (w *WriterAutoModel) TransSlice() {
	w.ContentSlice = strings.Split(w.Content, "\n")
}
