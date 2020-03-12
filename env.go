package tags

import (
	"os"
	"reflect"
)

const (
	ENV = "env"
)

func init() {
	tagHandlers[ENV] = getEnv
}

func getEnv(rawValue reflect.Value, tagValue string) (reflect.Value, error) {
	envValue := os.Getenv(tagValue)
	if envValue == "" {
		return rawValue, nil
	}
	return getValue(rawValue, envValue)
}
