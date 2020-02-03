package config

import (
	"fmt"
	"testing"
	"time"
)

func TestMin(t *testing.T) {
	type T struct {
		s    string        `min:"1"`
		S    string        `min:"1"`
		I    int           `min:"1"`
		I16  int16         `min:"1"`
		I32  int32         `min:"1"`
		I64  int64         `min:"1"`
		Ui   uint          `min:"1"`
		Ui16 uint16        `min:"1"`
		Ui32 uint32        `min:"1"`
		Ui64 uint64        `min:"1"`
		F32  float32       `min:"1"`
		F64  float64       `min:"1"`
		B    bool          `min:"true"`
		T    time.Duration `min:"1s"`
	}
	v := &T{}
	err := Revise(v)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	fmt.Printf("%#v\n", v)
}

func TestMax(t *testing.T) {
	type T struct {
		S    string        `max:"1"`
		I    int           `max:"1"`
		I16  int16         `max:"1"`
		I32  int32         `max:"1"`
		I64  int64         `max:"1"`
		Ui   uint          `max:"1"`
		Ui16 uint16        `max:"1"`
		Ui32 uint32        `max:"1"`
		Ui64 uint64        `max:"1"`
		F32  float32       `max:"1"`
		F64  float64       `max:"1"`
		B    bool          `max:"false"`
		T    time.Duration `max:"1s"`
	}
	v := &T{
		S:    "2",
		I:    2,
		I16:  2,
		I32:  2,
		I64:  2,
		Ui:   2,
		Ui16: 2,
		Ui32: 2,
		Ui64: 2,
		F32:  2,
		F64:  2,
		B:    true,
		T:    2 * time.Second,
	}
	err := Revise(v)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	fmt.Printf("%#v\n", v)
}

func TestDefault(t *testing.T) {
	type T struct {
		S    string        `default:"1"`
		I    int           `default:"1"`
		I16  int16         `default:"1"`
		I32  int32         `default:"1"`
		I64  int64         `default:"1"`
		Ui   uint          `default:"1"`
		Ui16 uint16        `default:"1"`
		Ui32 uint32        `default:"1"`
		Ui64 uint64        `default:"1"`
		F32  float32       `default:"1"`
		F64  float64       `default:"1"`
		B    bool          `default:"true"`
		T    time.Duration `default:"1s"`
	}
	v := &T{}
	err := Revise(v)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	fmt.Printf("%#v\n", v)
}

func Test(t *testing.T) {
	type T struct {
		N struct {
			S    string        `default:"1"`
			I    int           `default:"1"`
			I16  int16         `default:"1"`
			I32  int32         `default:"1"`
			I64  int64         `default:"1"`
			Ui   uint          `default:"1"`
			Ui16 uint16        `default:"1"`
			Ui32 uint32        `default:"1"`
			Ui64 uint64        `default:"1"`
			F32  float32       `default:"1"`
			F64  float64       `default:"1"`
			B    bool          `default:"true"`
			T    time.Duration `default:"1s"`
		}
	}
	v := &T{}
	err := Revise(v)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	fmt.Printf("%#v\n", v)
}
