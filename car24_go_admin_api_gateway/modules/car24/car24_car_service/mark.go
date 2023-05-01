package car24_car_service

type MarkModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateMarkModel struct {
	ID      string `json:"id"`
	BrandID string `json:"brand_id"`
	Name    string `json:"name"`
}

type UpdateMarkModel struct {
	ID      string `json:"id"`
	BrandID string `json:"brand_id"`
	Name    string `json:"name"`
}

type DeleteMarkModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type MarkQueryParamModel struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}

type MarkListModel struct {
	Marks []MarkModel `json:"marks"`
	Count int         `json:"count"`
}
