package pasha_bot

type Input struct {
	Date    string `json:"date" binding:"required"`
	Img     string `json:"img" binding:"required"`
	Message string `json:"message" binding:"required"`
}
