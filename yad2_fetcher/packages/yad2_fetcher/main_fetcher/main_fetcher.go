package main

import (
	"encoding/json"
	"fmt"
	gonadlan "github.com/devbrain/go-nadlan"
)

type Request struct {
	City    int  `json:"city"`
	Page    int  `json:"page"`
	ForSale bool `json:"forSale"`
}

type Workload struct {
	Data     []gonadlan.Yad2Data `json:"data"`
	LastPage int                 `json:"lastPage"`
	Error    string              `json:"error"`
	City     int                 `json:"city"`
	Page     int                 `json:"page"`
	ForSale  bool                `json:"forSale"`
}

type ResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type Response struct {
	Body       string            `json:"body,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

func Main(in Request) (*Response, error) {
	data, lastPage, err := gonadlan.GetYad2Data(in.Page, in.City, in.ForSale)
	errString := ""
	if err != nil {
		errString = fmt.Sprintf("%v", err)
	}
	workload := Workload{
		Data:     data,
		LastPage: lastPage,
		Error:    errString,
		City:     in.City,
		Page:     in.Page,
		ForSale:  in.ForSale,
	}
	body, err := json.Marshal(&workload)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body:       string(body),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
	}, nil
}
