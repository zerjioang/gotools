package model

// authentication relatead email data wrapper
type AuthMailRequest struct {
	Username string `json:"name,omitempty" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	// for account confirmation
	Confirmation string `json:"confirmation,omitempty"`
}
