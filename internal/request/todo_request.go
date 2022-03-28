package request

type TodoRequest struct {
	Todo        string `json:"todo" binding:"required"`
	Description string `json:"description" binding:"required"`
}
