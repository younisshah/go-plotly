package main

import (
	"encoding/json"
	"fmt"
	j "github.com/ricardolonga/jsongo"
	"strings"
)

// Support for {"style":{"type": "bar"}} type element in Kwargs

const (
	BASE_URL  = "https://plot.ly/clientresp"
	PLOT      = "plot"
	STYLE     = "style"
	LAYOUT    = "layout"
	NEW       = "new"
	OVERWRITE = "overwrite"
	APPEND    = "append"
	EXTEND    = "extend"
	TYPE      = "type"
)

type Request struct {
	un       string // done
	key      string // done
	origin   string //done
	platform string //done
	kwargs   Kwargs // done
	args     string // done
}

type Kwargs struct {
	FileName      string `json:"filename"`
	FileOpt       string `json:"fileopt"`
	Style         string `json:"style"`
	Traces        []int  `json:"traces"`
	WorldReadable bool   `json:"world_readable"`
}

type RequestBuilder interface {
	setKStyle(request Request) string
	build(request Request) string
}

type PlotlyRequest struct {
	BASE_URL string
	data     Request
}

func (req Request) build() (plotlyRequest PlotlyRequest, errors []string) {

	errors = validate(req)
	if len(errors) != 0 {
		return
	}

	req.platform = "go"
	req.kwargs.Style = j.Object().Put("style", j.Object().Put("type", req.kwargs.Style)).String()
	args, err := json.Marshal(req.args)
	if err != nil {
		fmt.Println(err)
		errors = []string{"JSON Error in KWARGS"}
		return
	}
	req.args = string(args)

	plotlyRequest.BASE_URL = BASE_URL
	plotlyRequest.data = req

	return
}

func validate(req Request) (errors []string) {
	if len(strings.Trim(req.un, " ")) == 0 {
		errors = append(errors, "'un' cannot be empty")
	}
	if len(strings.Trim(req.key, " ")) == 0 {
		errors = append(errors, "'key' cannot be empty")
	}
	if len(strings.Trim(req.origin, " ")) == 0 {
		errors = append(errors, "'key' cannot be empty")
	} else if req.origin != PLOT && req.origin != STYLE && req.origin != LAYOUT {
		errors = append(errors, "origin must be either plot, style or layout")
	}
	if len(strings.Trim(req.kwargs.FileName, " ")) == 0 {
		errors = append(errors, "'FileName' cannot be empty")
	}
	if len(strings.Trim(req.kwargs.FileOpt, " ")) == 0 {
		errors = append(errors, "'FileOpt' cannot be empty")
	}
	if len(strings.Trim(req.kwargs.Style, " ")) == 0 {
		errors = append(errors, "'Style' cannot be empty")
	}
	if len(strings.Trim(req.args, " ")) == 0 {
		errors = append(errors, "'args' cannot be empty")
	}
	return
}

func main() {

	var kwargs = Kwargs{}
	kwargs.FileName = "go-plotly-test"
	kwargs.FileOpt = OVERWRITE
	kwargs.Traces = []int{0, 3, 5}
	kwargs.WorldReadable = false
	kwargs.Style = "bar"

	var args = `[{"x": [0, 1, 2], "y": [3, 1, 6]}]`

	var req = Request{}
	req.un = "younisashah"
	req.key = "3rf0bfpzer"
	req.origin = PLOT
	req.kwargs = kwargs
	req.args = args

	plotlyReq, err := req.build()

	if err != nil {
		fmt.Println(err)
	} else {
		// TODO: Creating a POST request
		fmt.Println(plotlyReq)
	}
}
