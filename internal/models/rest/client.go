package rest

type Client struct {
	ID                 uint   `json:"id"`
	PhoneNumber        string `json:"phone_number" binding:"required,min=11,max=11"`
	MobileOperatorCode string `json:"mobile_operator_code" binding:"required,min=3,max=3"`
	Tag                string `json:"tag" binding:"required"`
	Timezone           string `json:"timezone" binding:"required"`
}

type NewClientResponse struct {
	ID uint `json:"id"`
}
