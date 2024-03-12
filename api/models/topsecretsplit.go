package models

type TSSRequestBody struct {
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}
