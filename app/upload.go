package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (a *App) upload(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	// fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	io.Copy(&buf, file)
	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()
	return
}
