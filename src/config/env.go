package config

type Env struct {
	Username string
	Password string
	Protocol string
	Host     string
	Port     string
	Name     string
	Parse    string
	Driver   string
}

func NewEnv() *Env {
	Env := Env{
		Username: "user",
		Password: "password",
		Protocol: "tcp",
		Host:     "localhost",
		Port:     "3306",
		Name:     "golang-ddd-clear-architecture,
		Parse:    "true",
		Driver:   "mysql",
	}

	return &Env
}
