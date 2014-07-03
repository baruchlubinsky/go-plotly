package plotly

import (
	"fmt"
)

func StackedBarPlot(categories []string, data map[string][]interface{}, filename string, title string, xTitle string, yTitle string, public bool) (Url, error) {
	traces := make([]Trace, 0, len(data))
	for i, category := range categories {
		x := make([]interface{}, 0)
		y := make([]interface{}, 0)
		for key, values := range data {
			x = append(x, key)
			y = append(y, values[i])
		}
		traces = append(traces, Trace{
			X:    x,
			Y:    y,
			Name: category,
			Type: "bar",
		})
	}
	layout := fmt.Sprintf(`{
    "title":"%v",
    "barmode":"stack",
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
