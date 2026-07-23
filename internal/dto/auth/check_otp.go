package auth

type CheckOtpRequest struct {
	PhoneNumber string `json:"phone-number" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type CheckOtpResponse struct{}
