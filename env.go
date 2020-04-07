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

func getEnv(rawValue reflect.Value, tagValue string) error {
	envValue := os.Getenv(tagValue)
	if envValue == "" {
		return nil
	}

	v, err := getValue(rawValue, envValue)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(rawValue, v) {
		rawValue.Set(v)
	}
	return nil
}
