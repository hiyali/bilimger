package lib

import (
	"flag"
)

type Environment struct {
	Env string
}

var (
	environment = Environment{
		Env: "dev", // default env dev
	}
)

func InitialEnvironment() {
	env := flag.String("env", "dev", "Environment")
	flag.Parse()

	environment = Environment{
		Env: *env,
	}
}

func GetEnv() string {
	return environment.Env
}
