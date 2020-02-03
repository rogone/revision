package config

import (
	"fmt"
	"reflect"
)

func Revise(st interface{}) error {
	value := reflect.ValueOf(st).Elem()
	return revise(value)
}

func revise(value reflect.Value) error {
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.FieldByName(field.Name)

		if !fieldValue.CanSet() {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			err := revise(fieldValue)
			if err != nil {
				return err
			}
			continue
		}

		for k, v := range tagHandlers {
			tagValue := field.Tag.Get(k)
			if tagValue != "" {
				destValue, err := v(fieldValue, tagValue)
				if err != nil {
					return fmt.Errorf("revise field:%s tag:%s err:%w", field.Name, k, err)
				}
				if !reflect.DeepEqual(fieldValue, destValue) {
					fieldValue.Set(destValue)
				}
			}
		}
	}
	return nil
}
