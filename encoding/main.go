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

var testData = User{
	ID:        "1k23j123kk2h13-1jk231jh23-1j2h3lk1j2h3lk1-h1j2312f3",
	Mode:      "local",
	CreatedAt: 123123123,
	LastLogin: 123123123,
	Email:     "test@example.com",
	Phone:     "+11231212",

	Data: map[string]interface{}{
		"username": "test",
		"age":      12,
	},

	EmailVerified: true,
	PhoneVerified: true,
}

var testDataProtobuf = &UserProtobuf{
	Id:        "1k23j123kk2h13-1jk231jh23-1j2h3lk1j2h3lk1-h1j2312f3",
	Mode:      "local",
	CreatedAt: 123123123,
	LastLogin: 123123123,
	Email:     "test@example.com",
	Phone:     "+11231212",

	Data: map[string]string{
		"username": "test",
		"age":      "12",
	},

	EmailVerified: true,
	PhoneVerified: true,
}
