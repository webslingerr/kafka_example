package payme

type Params struct {
	Id      string            `json:"id"`
	Amount  int64             `json:"amount"`
	Time    int64             `json:"time"`
	Account map[string]string `json:"account"`
	Reason  int               `json:"reason"`
	From    int64             `json:"from"`
	To      int64             `json:"to"`
}

type Body struct {
	Method string `json:"method"`
	Params Params `json:"params"`
}

type Error struct {
	Code    int32       `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}

type SuccessResponse struct {
	Result map[string]interface{} `json:"result"`
}

type Response struct {
	SuccessResponse SuccessResponse
	ErrorResponse   *ErrorResponse
}

type PaymentNotification struct {
	ID          string `json:"id"`
	CelebrityID string `json:"celebrity_id"`
	FcmToken    string `json:"fcm_token"`
	Platform    int32  `json:"platform"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	Content     string `json:"content"`
	Amount      int64  `json:"amount"`
	SessionID   string `json:"-"`
}
