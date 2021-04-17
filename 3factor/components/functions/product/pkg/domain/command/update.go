package command

type UpdateProduct struct {
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
