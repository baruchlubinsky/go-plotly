package plotly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Trace struct {
	X      []interface{} `json:"x"`
	Y      []interface{} `json:"y,omitempty"`
	Name   string        `json:"name"`
	YAxis  string        `json:"yaxis"`
	Type   string        `json:"type"`
	Marker Marker        `json:"marker,omitempty"`
	Line   Line          `json:"line,omitempty"`
	Fill   string        `json:"fill,omitempty"`
}

type Marker struct {
	Line    Line    `json:"line,omitempty"`
	Opacity float64 `json:"opacity"`
	Color   string  `json:"color"`
}

type Line struct {
	Color string `json:"color,omitempty"`
}

func Create(filename string, figure Figure, public bool) (url Url, err error) {
	request := NewRequest()
	request.Origin = "plot"
	args, err := json.Marshal(figure.Data)
	if err != nil {
		return
	}
	request.Args = string(args)
	request.Kwargs = fmt.Sprintf(`{"filename":"%v",
        "fileopt":"overwrite",
        "world_readable":%v,
        "layout":%v
  }`, filename, public, figure.Layout)
	result, err := Post(request)
	if err != nil {
		return
	}
	if result.Url == "" {
		return Url(""), result
	}
	return Url(result.Url), nil
}

func Download(figure Figure, filename string) (err error) {
	payload := Payload{Figure: figure}
	data, err := json.Marshal(payload)
	if err != nil {
		return
	}
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
	err = ioutil.WriteFile(filename, body, 0777)
	return
}

func Save(id string, filename string) error {
	response, err := Get(id)
	if err != nil {
		return err
	} else if response.Payload.Figure.Data == nil {
		return response
	}
	err = Download(response.Payload.Figure, filename)
	return err
}
