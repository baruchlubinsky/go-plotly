package main

import (
  "github.com/baruchlubinsky/plotly"
  "fmt"
  "strings"
)

// This program creates a plot on plotly, retrieves it, and saves it as an image.
func main () {
  request := plotly.NewRequest()
  request.Origin = "plot"
  request.Args = `[{"name":"0","y":[0,27,32,37,39,40]},{"name":"10","y":[0,30,37,39,40,0]},{"name":"20","y":[0,32,37,39,40,0]},{"name":"30","y":[0,25,30,37,38,40]},{"name":"40","y":[0,25,30,35,37,40]},{"name":"50","y":[0,30,33,37,38,40]},{"name":"60","y":[0,18,24,35,37,40]},{"name":"70","y":[0,25,29,34,37,40]},{"name":"80","y":[0,27,30,35,37,40]},{"name":"90","y":[0,21,26,31,35,39]},{"name":"100","y":[0,17,21,27,32,39]},{"name":"110","y":[0,15,19,25,30,39]},{"name":"120","y":[0,15,18,23,27,37]},{"name":"130","y":[0,19,24,28,32,39]},{"name":"140","y":[0,24,28,33,37,39]},{"name":"150","y":[0,21,25,30,35,39]},{"name":"160","y":[0,17,21,27,33,39]},{"name":"170","y":[0,21,25,30,34,40]},{"name":"180","y":[0,21,25,30,33,39]},{"name":"190","y":[0,20,24,28,32,40]},{"name":"200","y":[0,21,25,30,33,39]},{"name":"210","y":[0,22,26,31,35,40]},{"name":"220","y":[0,20,25,30,35,40]},{"name":"230","y":[0,20,25,30,35,40]},{"name":"240","y":[0,21,25,30,34,40]},{"name":"250","y":[0,21,25,30,32,40]},{"name":"260","y":[0,17,19,23,27,38]},{"name":"270","y":[0,13,17,21,24,39]},{"name":"280","y":[0,13,17,22,26,38]},{"name":"290","y":[0,23,25,29,31,39]},{"name":"300","y":[0,19,22,25,28,38]},{"name":"310","y":[0,17,20,23,27,38]},{"name":"320","y":[0,15,18,22,25,36]},{"name":"330","y":[0,16,19,23,25,35]},{"name":"340","y":[0,17,20,23,26,35]},{"name":"350","y":[0,18,22,24,26,35]},{"name":"360","y":[0,17,19,23,26,35]},{"name":"370","y":[0,17,19,23,26,35]},{"name":"380","y":[0,15,17,21,23,33]},{"name":"390","y":[0,14,17,21,23,33]},{"name":"400","y":[0,19,22,24,26,33]},{"name":"410","y":[0,17,20,22,23,26]},{"name":"420","y":[11,15,17,20,22,28]},{"name":"430","y":[11,13,16,19,22,32]},{"name":"440","y":[9,17,19,21,23,31]},{"name":"450","y":[9,19,21,23,25,28]},{"name":"460","y":[12,21,22,24,26,30]},{"name":"470","y":[16,19,20,22,25,27]}]`
  opts := `{"filename": "plot from golang api",
        "fileopt": "overwrite",
        "world_readable": true,
        "style":{"type": "box", "boxpoints": false, "marker":{"color":"rgb(0.1, 0.3, 0.9)"}},
        "layout":{
          "title":"Distribution of quality scores",
          "yaxis":{
            "title": "Quality score",
            "zeroline": true},
          "xaxis":{
            "title":"Read position"
          },
          "showlegend":false
        }
  }`
  request.Kwargs = opts
  result, err := plotly.Post(request)
  fmt.Println(result)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("Succesfully create plot!\nFilename:%v\nURL:%v\n", result.Filename, result.Url)
  }
  fields := strings.Split(result.Url, "/")
  id := fields[4]
  response, err := plotly.Get(id)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Downloaded plot")
  }
  err = plotly.Download(response.Payload.Figure, "image.png")
  if err != nil {
    fmt.Println(err)
  }
}
