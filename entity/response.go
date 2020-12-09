package entity

//Response ...
type Response struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}
