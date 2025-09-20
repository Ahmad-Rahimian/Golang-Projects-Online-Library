package auth

type SendOTPRequest struct {
	Phone string `json:"phone" example:"09356769697" binding:"required"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone" example:"09356769697" binding:"required"`
	OTP   string `json:"otp" example:"123456" binding:"required"`
}
