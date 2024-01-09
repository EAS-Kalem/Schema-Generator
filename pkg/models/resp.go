package models

type RespR struct {
	Service  string `json:"Service"`
	Handlers HandlersR  `json:"Handlers"`
}

type HandlersR struct {
	FileName    string `json:"FileName"`
	ServiceName string `json:"ServiceName"`
	Args []string `json:"Args"`
} 