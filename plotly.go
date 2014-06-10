package plotly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
)

const ROOTURL = "https://plot.ly/"
const POSTURL = ROOTURL + "clientresp/"
const GETURL = ROOTURL + "apigetfile/"
const IMAGEURL = ROOTURL + "apigenimage/"
const VERSION = "2.0"
const PLATFORM = "golang"

var username = "baruchlubinsky"
var apikey = "knk0dbvmeu"

type Request struct {
	Un       string
	Key      string
	Origin   string
	Platform string
	Version  string
	Args     string
	Kwargs   string
}

type PostResponse struct {
	Filename string
	Url      string
	Error    string
	Warning  string
	Message  string
}

type Figure struct {
	Layout interface{} `json:"layout"`
	Data interface{} `json:"data"`
}

type Payload struct {
	Figure Figure `json:"figure"`
}

type GetResponse struct {
	Message string `json:"message"`
	Warning string `json:"warning"`
	Payload Payload `json:"payload"`
	Error string `json:"error"`
}

type DownloadResponse struct {
	Message string
	Warning string
	Payload []byte
	Error string
}

func (r *Request) urlEncode() url.Values {
	v := url.Values{}
	v.Set("un", r.Un)
	v.Set("key", r.Key)
	v.Set("origin", r.Origin)
	v.Set("platform", r.Platform)
	v.Set("version", r.Version)
	v.Set("args", r.Args)
	v.Set("kwargs", r.Kwargs)
	return v
}

func setHeaders(request *http.Request) {
	request.Header.Set("plotly-username", username)
	request.Header.Set("plotly-apikey", apikey)
	request.Header.Set("plotly-version", VERSION)
	request.Header.Set("plotly-platform", PLATFORM)
}

func Post(data Request) (result PostResponse, err error) {
	client := http.DefaultClient
	response, err := client.PostForm(POSTURL, data.urlEncode())
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &result)
	return
}

func Get(id string) (result *GetResponse, err error) {
	request, _ := http.NewRequest("GET", GETURL + username + "/" + id, nil)
	setHeaders(request)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &result)
	return
}

func Download(figure Figure, filename string) (err error) {
	data, err := json.Marshal(figure)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	request, _ := http.NewRequest("POST", IMAGEURL, bytes.NewReader(data))
	setHeaders(request)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	var result = DownloadResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, result.Payload, 0777)
	return
}
