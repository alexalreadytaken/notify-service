package rest

type ExternalMessage struct {
	Id    uint   `json:"id"`
	Phone string `json:"phone"`
	Text  string `json:"text"`
}
