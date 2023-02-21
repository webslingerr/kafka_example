package payme

func CheckPerformTransactionAnswer() SuccessResponse {
	return SuccessResponse{
		Result: map[string]interface{}{"allow": true},
	}
}
