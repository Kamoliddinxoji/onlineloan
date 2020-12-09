package entity

//Customer ...
type customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	PhoneNum string `json:"phone_num"`
}
