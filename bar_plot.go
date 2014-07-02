package plotly

import (
	"fmt"
)

func StackedBarPlot(categories []interface{}, data map[string][]interface{}, filename string, title string, xTitle string, yTitle string, public bool) (Url, error) {
	traces := make([]Trace, 0, len(data))
	for key, values := range data {
		traces = append(traces, Trace{
			X:    values,
			Y:    categories,
			Name: key,
			Type: "bar",
		})
	}
	layout := fmt.Sprintf(`{
    "title":"%v",
    "barmode":"stacked",
    "yaxis":{
      "title":"%v"
    },
    "xaxis":{
      "title":"%v",
      "type":"category"
    }
  }
  `, title, yTitle, xTitle)
	figure := Figure{
		Data:   traces,
		Layout: layout,
	}
	return Create(filename, figure, public)
}
