package tags

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

var (
	ErrNotPtr = errors.New("not ptr")
)

// process tags
func Revise(st interface{}) error {
	value := reflect.ValueOf(st)
	if value.Kind() != reflect.Ptr {
		return ErrNotPtr
	}

	return revise(value.Elem())
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

		for _, v := range parseTags(string(field.Tag)) {
			h, ok := tagHandlers[v.name]
			if !ok {
				continue
			}

			err := h(fieldValue, v.value)
			if err != nil {
				return fmt.Errorf("revise field:%s tag:%s err:%w", field.Name, v.name, err)
			}
		}
	}
	fs.Parse(os.Args[1:])
	return nil
}
