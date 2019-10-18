/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-18 14:21:25
 * @LastEditTime: 2019-10-18 14:21:25
 * @LastEditors: your name
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><form method = \"POST\" action = \"upload\" "+
			" enctype = \"multipart/form-data\">"+
			"Choose an image to upload: <input name = \"image\" type = \"file\" />"+
			"<input type = \"submit\" value = \"upload\" />"+
			"</form></html>")
		return
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func main() {
	http.HandleFunc("/view", ViewHandler)
	http.HandleFunc("/upload", UploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
