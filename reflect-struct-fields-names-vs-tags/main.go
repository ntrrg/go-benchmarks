package main

type User struct {
	ID        string `json:"id"`
	Mode      string `json:"mode"`
	CreatedAt int64  `json:"createdAt"`
	LastLogin int64  `json:"lastLogin"`

	Data map[string]interface{} `json:"data,omitempty"`

	// Verification methods.
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	EmailVerified bool   `json:"emailVerified"`
	PhoneVerified bool   `json:"phoneVerified"`
}
