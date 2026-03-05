package sharedutil

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Load[T any](instance *T, filenames ...string) error {
	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return err
		}
	}

	if err := env.Parse(instance); err != nil {
		return err
	}

	return nil
}

func Overload[T any](instance *T, filenames ...string) error {
	if len(filenames) > 0 {
		if err := godotenv.Overload(filenames...); err != nil {
			return err
		}
	}

	if err := env.Parse(instance); err != nil {
		return err
	}

	return nil
}

func MustLoad[T any](instance *T, filenames ...string) {
	Must(Load(instance, filenames...))
}
