package model

// She Struct prefix on the structure fields helps disambiguate
type SignupRequest struct {
	StructPhone     string `json:"phone"`
	StructEmaill    string `json:"email"`
	StructAccStatus bool   `json:"status"`
}

type VerifyOTPrequest struct {
	StructUuid  string `json:"uuid"`
	StructToken string `json:"token"`
}

type UUIDResponse struct {
	StructUUID string `json:"uuid"`
}

type VerifySupabaseResponse struct {
	StructAccessToken string `json:"access_token"`
}

type AccessTkJsonObject struct {
	StructAccessTk string
}
