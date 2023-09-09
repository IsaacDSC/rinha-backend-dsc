package dto

type Person struct {
	ID       string
	Name     string   `json:"nome"`
	LastName string   `json:"apelido"` //unique
	Birthday string   `json:"nascimento"`
	Stack    []string `json:"stack"`
}
