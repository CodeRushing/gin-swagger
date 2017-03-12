package client_service

import (
	"github.com/morlay/gin-swagger/example/service/test2"
)

type SomeTest struct {
	Common
	//
	ErrorMap ErrorMap `json:"errorMap"`
	//
	State test2.State `json:"state,string" validate:"@string{TWO}"`
}

type Common struct {
	// 总数
	Total int8 `json:"total"`
}

type ErrorMap map[string]map[string]int64

type ItemData struct {
	//
	Id string `json:"id"`
	//
	Name string `json:"name" validate:"@string[0,)"`
	//
	StartTime test2.Date `json:"startTime,string"`
	//
	State test2.State `json:"state,string"`
}

type ReqBody struct {
	//
	Name string `json:"name"`
	//
	UserName string `json:"username"`
}

type Some struct {
	//
	StartTime test2.Date `json:"startTime,string"`
	// Test
	State test2.State `json:"state,string" validate:"@string{,TWO}"`
	//
	Data []ItemData `json:"data"`
	//
	Name uint64 `json:"name,string"`
}
