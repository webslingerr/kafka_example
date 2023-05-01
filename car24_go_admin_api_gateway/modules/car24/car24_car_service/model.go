package car24_car_service

type CarModel struct {
	ID          string `json:"id"`
	MarkID      string `json:"mark_id"`
	CategoryID  string `json:"category_id"`
	InvestorID  string `json:"investor_id"`
	StateNumber string `json:"state_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCarModel struct {
	ID          string `json:"id"`
	MarkId      string `json:"mark_id"`
	CategoryID  string `json:"category_id"`
	InvestorID  string `json:"investor_id"`
	StateNumber string `json:"state_number"`
}

type UpdateCarModel struct {
	ID          string `json:"id"`
	MarkID      string `json:"mark_id"`
	CategoryID  string `json:"category_id"`
	InvestorID  string `json:"investor_id"`
	StateNumber string `json:"state_number"`
}

type DeleteCarModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type CarQueryParamModel struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}

type CarListModel struct {
	Cars  []CarModel `json:"cars"`
	Count int        `json:"count"`
}
