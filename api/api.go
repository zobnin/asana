package api

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"asana/config"
	"asana/utils"
)

const (
	GetBase   = "https://app.asana.com"
	PostBase  = "https://app.asana.com/api/1.0"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/36.0.1985.125 Safari/537.36"
)

type Base struct {
	Id   int
	Name string
}

func Get(path string, params url.Values) []byte {
	req, err := http.NewRequest("GET", getURL(path, params), nil)
	utils.Check(err)
	return fire(req)
}

func Post(path string, data string) []byte {
	req, err := http.NewRequest("POST", PostBase+path, strings.NewReader(data))
	utils.Check(err)
	return fire(req)
}

func Put(path string, data string) []byte {
	req, err := http.NewRequest("PUT", PostBase+path, strings.NewReader(data))
	utils.Check(err)
	return fire(req)
}

func Delete(path string) []byte {
	req, err := http.NewRequest("DELETE", PostBase+path, nil)
	utils.Check(err)
	return fire(req)
}

func Upload(uri string, path string) []byte {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	utils.Check(err)
	_, err = io.Copy(part, file)

	err = writer.Close()
	utils.Check(err)

	req, err := http.NewRequest("POST", PostBase+uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	utils.Check(err)
	return fire(req)
}

func getURL(path string, params url.Values) string {
	if params == nil || params.Encode() == "" {
		return GetBase + path
	} else {
		return GetBase + path + "?" + params.Encode()
	}
}

func fire(req *http.Request) []byte {
	client := &http.Client{}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Authorization", "Bearer "+config.Load().Personal_access_token)

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	utils.Check(err)

	if resp.StatusCode >= 300 {
		println(resp.Status)
	}

	return body
}
