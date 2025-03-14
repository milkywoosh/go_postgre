package models

type BodyReponseAPI struct {
	Data    any    `json:"data"` // same as interface{} type
	Message string `json:"message"`
}
