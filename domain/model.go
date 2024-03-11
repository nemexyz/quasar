package domain

type AllianceResponse struct {
	Position PositionResponse `json:"position"`
	Message  string           `json:"message"`
}

type PositionResponse struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
