package plotly

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func init() {
	username = os.Getenv("PLOTLY_USERNAME")
	apikey = os.Getenv("PLOTLY_APIKEY")
	authenticated = true
	var credentialsFile *os.File
	files := []string{
		"/etc/plotly/plotly_credentials.json",
		"/etc/plotly/.plotly_credentials.json",
		path.Join(os.Getenv("HOME"), "plotly_credentials.json"),
		path.Join(os.Getenv("HOME"), ".plotly_credentials.json"),
		"plotly_credentials.json",
		".plotly_credentials.json",
	}
	for _, path := range files {
		file, err := os.Open(path)
		if err == nil {
			credentialsFile = file
			data, err := ioutil.ReadAll(credentialsFile)
			if err != nil {
				println("Cannot read " + file.Name())
				continue
			}
			credentials := struct {
				Username string
				Apikey   string
			}{}
			err = json.Unmarshal(data, &credentials)
			if err != nil {
				println("Badly fomatted " + file.Name())
				continue
			}
			username = credentials.Username
			apikey = credentials.Apikey
		}
	}

	if username == "" || apikey == "" {
		authenticated = false
	}
}

const ROOTURL = "https://plot.ly/"
const POSTURL = ROOTURL + "clientresp/"
const GETURL = ROOTURL + "apigetfile/"
const IMAGEURL = ROOTURL + "apigenimage/"
const VERSION = "2.0"
const PLATFORM = "golang"

var username string
var apikey string
var authenticated bool

type Request struct {
	Un       string
	Key      string
	Origin   string
	Platform string
	Version  string
	Args     string
	Kwargs   string
}

type Response struct {
	ErrorMessage string `json:"error"`
	Warning      string
	Message      string
}

type PostResponse struct {
	Filename string
	Url      string
	Response
}

type Payload struct {
	Figure Figure `json:"figure"`
}

type GetResponse struct {
	Payload Payload `json:"payload"`
	Response
}

type DownloadResponse struct {
	Payload []byte
	Response
}

type Figure struct {
	Layout interface{} `json:"layout"`
	Data   interface{} `json:"data"`
}

type PlotlyError string

func (e PlotlyError) Error() string {
	return string(e)
}

type Url string

func NewRequest() *Request {
	var request = Request{
		Un:       username,
		Key:      apikey,
		Platform: PLATFORM,
		Version:  VERSION,
	}
	return &request
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

func Post(data *Request) (result PostResponse, err error) {
	if !CheckCredentials() {
		return PostResponse{}, PlotlyError("Not authenticated")
	}
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
	if !CheckCredentials() {
		return nil, PlotlyError("Not authenticated")
	}
	request, _ := http.NewRequest("GET", GETURL+username+"/"+id, nil)
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

func (u Url) Id() string {
	fields := strings.Split(string(u), "/")
	if len(fields) == 5 {
		return fields[4]
	}
	return ""
}

func (r Response) Error() string {
	if r.ErrorMessage != "" {
		return r.ErrorMessage
	} else if r.Warning != "" {
		return r.Warning
	} else if r.Message != "" {
		return r.Message
	} else {
		return "An unspecified error occured at Plot.ly"
	}
}

func CheckCredentials() bool {
	return authenticated
}
