package server

import (
	"fmt"
	"net/http"
	"strconv"
	"webdraw/mouse"
)

func indexAction (responseWriter http.ResponseWriter, request *http.Request) {
	var query  = request.URL.Query()
	var i byte = 1
	var cursor *mouse.Mouse = mouse.New()
	var xStr, yStr string

	for {
		xStr = query.Get(fmt.Sprintf("x%v", i))
		yStr = query.Get(fmt.Sprintf("y%v", i))

		if xStr != "" && yStr != "" {
			x, xErr := strconv.ParseFloat(xStr, 64)
			y, yErr := strconv.ParseFloat(yStr, 64)

			if xErr == nil && yErr == nil {
				cursor.SetXY(int32(x), int32(y))
				cursor.Set()
			}
		} else {
			break
		}

		i++
	}
}

func Start (port string) {
	http.HandleFunc("/webdraw/", indexAction)
	http.ListenAndServe(port, nil)
}
