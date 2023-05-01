package car24_car_service

type BrandModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateBrandModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateBrandModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteBrandModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type BrandQueryParamModel struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}

type BrandListModel struct {
	Brands []BrandModel `json:"brands"`
	Count  int          `json:"count"`
}
