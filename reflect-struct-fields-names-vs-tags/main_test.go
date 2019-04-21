package main

import (
  "reflect"
	"testing"
)

const tagName = "json"

var user = &User{
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

func BenchmarkFieldNames(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		t := reflect.TypeOf(*user)

		for j := 0; j < t.NumField(); j++ {
			_ = t.Field(j).Name
		}
	}
}

func BenchmarkTags(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		t := reflect.TypeOf(*user)

		for j := 0; j < t.NumField(); j++ {
			_ = t.Field(j).Tag.Get(tagName)
		}
	}
}
