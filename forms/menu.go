package forms

type MenuCreate struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
