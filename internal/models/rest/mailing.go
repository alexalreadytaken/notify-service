package rest

import "time"

type Mailing struct {
	ID            uint      `json:"id"`
	StartingTime  time.Time `json:"starting_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	Text          string    `json:"text" binding:"required"`
	SendindFilter string    `json:"sending_filter" binding:"required,oneof=BY_TAG BY_OPERATOR"`
	FilterValue   string    `json:"filter_value" binding:"required"`
	EndingTime    time.Time `json:"end_time" binding:"required" time_format:"2006-01-02 15:04:05"`
}

type NewMailingResponse struct {
	ID uint `json:"id"`
}
