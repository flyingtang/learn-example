package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

const uploadFilePath = "static"

func main() {

	hfs := http.FileServer(http.Dir("static"))

	// 单文件上传
	http.HandleFunc("/file/one", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			// 文件上传解析
			// 限制在内存中的占用
			//req.ParseMultipartForm(32 << 20)
			f, h, err := req.FormFile("myfile")

			if err != nil {
				returnMessage(res, http.StatusBadRequest, "解析错误")
				return
			}
			var currentfilePath = strings.Join([]string{uploadFilePath, "img", h.Filename}, "/")
			if file, err := os.Create(currentfilePath); err != nil {
				returnMessage(res, http.StatusBadRequest, "解析错误"+err.Error())
				return
			} else {
				io.Copy(file, f)
				file.Close()
				returnMessage(res, http.StatusOK, "上传成功!!!")
				return
			}
		} else {
			returnMessage(res, http.StatusBadRequest, "暂时未提供服务")
		}
	})

	// 多文件上传
	http.HandleFunc("/file/more", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			// 文件上传解析
			mp, err := req.MultipartReader()
			if err != nil {
				returnMessage(res, http.StatusBadRequest, err.Error())
				return
			}

			for {

				part, err := mp.NextPart()
				if err != nil {
					returnMessage(res, http.StatusBadRequest, err.Error())
					return
				}

				// 当前文件名字
				fileName := part.FileName()
				if len(fileName) == 0 {
					continue
				}

				var currentfilePath = strings.Join([]string{uploadFilePath, "img", fileName}, "/")
				err = writeFile(currentfilePath, part)
				if err != nil {
					returnMessage(res, http.StatusBadRequest, err.Error())
					return
				}
			}
		} else {
			returnMessage(res, http.StatusBadRequest, "暂时未提供服务")
		}
	})

	http.Handle("/", hfs)

	panic(http.ListenAndServe("0.0.0.0:3005", nil))
}

func writeFile(filePath string, part *multipart.Part) error {

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New(" os.Create " + err.Error())
	}
	var buf = make([]byte, 4*1024)
	defer file.Close()
	for {
		n, err := part.Read(buf)
		if err != nil {
			if err == io.EOF {
				_, err := file.Write(buf[:n])
				if err != nil {
					return  err
				}
				return nil
			} else {
				return err
			}
		} else {
			_, err := file.Write(buf[:n])
			if err != nil {
				return  err
			}
		}
	}

}

func returnMessage(res http.ResponseWriter, code int, message string) {
	res.WriteHeader(code)
	r := map[string]string{
		"message": message,
	}
	data := dataEncode(r)
	res.Write(data)
}

func dataEncode(v interface{}) []byte {

	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println("json.Marshal err ", err.Error())
	}
	return data
}
