package models


type Schema struct {
	Service string `json:"service" validate:"required"`
	Action  string `json:"action" validate:"required"`
	Args    struct {
		Host        string `json:"host" validate:"required"`
		User        string `json:"user" validate:"required"`
		Password    string `json:"password" validate:"required"`
		VM          string `json:"vm" validate:"required"`
		CallbackURL string `json:"callback_url" validate:"required"`
	} `json:"args"`
}

// type Dog struct {
//     Service      string `json:"service" validate:"required"`
//     Action      string `json:"action" validate:"required"`
// 	Args struct {
//     	Host      int    `json:"host" validate:"required"`
//     	User bool   `json:"user" validate:"required"`
// 		Password bool   `json:"password" validate:"required"`
// 	} `json:"args"`
// }

type IError struct {
    Field string
    Tag   string
    Value string
}

