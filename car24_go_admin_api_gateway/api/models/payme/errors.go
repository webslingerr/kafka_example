package payme

func WrongAmountException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -31001,
			Message: "Wrong Amount",
			Data:    "amount",
		},
	}
}

func InternalServerException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -32400,
			Message: "Internal server error",
			Data:    "error",
		},
	}
}

func UnableCompleteException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -31008,
			Message: "Unable to complete operation",
			Data:    "transaction",
		},
	}
}

func UnauthorizedException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -32504,
			Message: "unauthorized request",
		},
	}
}

func TransactionNotFoundException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -31003,
			Message: "Transaction not found",
			Data:    "transaction",
		},
	}
}

func UnableCancelTransactionException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -31007,
			Message: "Unable to cancel transaction",
			Data:    "transaction",
		},
	}
}

func OrderNotExistsException() *ErrorResponse {
	return &ErrorResponse{
		Error: &Error{
			Code:    -31050,
			Message: "Order not found",
			Data:    "order",
		},
	}
}
