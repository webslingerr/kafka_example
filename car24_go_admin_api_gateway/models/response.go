package models

type ResponseOK struct {
	Message string `json:"message"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Error Error `json:"error"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
