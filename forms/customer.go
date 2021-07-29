package forms

type (
	CustomerRegister struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}

	CustomerReservation struct {
		TableNumber int `json:"table_number"`
	}
)
