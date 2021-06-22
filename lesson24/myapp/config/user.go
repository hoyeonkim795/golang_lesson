package config

type User struct {
	Username string `json:"username"`
	Points	int	`json:"points"`
	Rank	int	`json:"rank"`
}