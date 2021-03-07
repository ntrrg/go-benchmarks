package main

type Car struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type Person struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Alive   bool      `json:"alive"`
	Numbers []int     `json:"numbers"`
	Car     Car       `json:"car"`
	Family  []*Person `json:"family"`

	Data []byte `json:"data"`
}
