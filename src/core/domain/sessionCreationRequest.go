package domain

type Credentials struct {
	Password string `json:"password" binding:"required"`
}

type SessionCreationRequest struct {
	User        string      `json:"user" binding:"required"`
	Credentials Credentials `json:"credentials" binding:"required"`
}
