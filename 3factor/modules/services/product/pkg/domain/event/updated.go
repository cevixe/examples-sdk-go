package event

type ProductUpdated struct {
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
