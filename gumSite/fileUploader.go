package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 1. parse input, type multipart/form-data.
	r.ParseMultipartForm(10 << 20)

	// 2. retrieve file from posted form-data.
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File size : %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)

	// 3. write temporary file on our server
	// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	f, err := os.Create("static/images" + name + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(fileBytes))
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
