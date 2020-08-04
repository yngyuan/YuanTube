package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}


func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Entered the streamHandler")
	targetUrl := "URL" + p.ByName("vid-id")
	http.Redirect(w, r, targetUrl, 301)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR +filename, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ossfn := "videos/" + filename
	path := "./videos/" + filename
	bn := "avenssi-videos2"
	ret := UploadToOss(ossfn, path, bn)
	if !ret {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}