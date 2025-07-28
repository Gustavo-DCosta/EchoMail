package model

type SignupRequest struct {
	Phone_Number  string `json:"phone_number"`
	Email_Address string `json:"email_address"`
}

type VerifyOTPpayload struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

type VerifySupabaseResponse struct {
	AccesToken string `json:"access_token"`
}
