package tags

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFlag(t *testing.T) {
	fs = flag.NewFlagSet("TestFlag", flag.ExitOnError)
	type T struct {
		N struct {
			S string `flag:"flag_S"`
			I int    `flag:"flag_I"`
			//I16  int16         `flag:"flag_I16"`
			//I32  int32         `flag:"flag_I32"`
			I64 int64 `flag:"flag_I64"`
			Ui  uint  `flag:"flag_U"`
			//Ui16 uint16        `flag:"flag_U16"`
			//Ui32 uint32        `flag:"flag_U32"`
			Ui64 uint64 `flag:"flag_U64"`
			//F32  float32       `flag:"flag_F32"`
			F64 float64       `flag:"flag_F64"`
			B   bool          `flag:"flag_B"`
			T   time.Duration `flag:"flag_T"`
		}
	}

	os.Args = []string{
		"",
		"-flag_S=1",
		"-flag_I=1",
		"-flag_I64=1",
		"-flag_U=1",
		"-flag_U64=1",
		"-flag_F64=1",
		"-flag_B=1",
		"-flag_T=1s",
	}
	v := &T{}

	err := Revise(v)
	if err != nil {
		t.Errorf("Testflag %s", err.Error())
	}

	fmt.Printf("Testflag %#v\n", v)
}

func TestFlagDefault(t *testing.T) {
	fs = flag.NewFlagSet("TestFlagDefault", flag.ExitOnError)
	type T struct {
		N struct {
			S string `flag:"flag_S,1"`
			I int    `flag:"flag_I,1"`
			//I16  int16         `flag:"flag_I16"`
			//I32  int32         `flag:"flag_I32"`
			I64 int64 `flag:"flag_I64,1"`
			Ui  uint  `flag:"flag_U,1"`
			//Ui16 uint16        `flag:"flag_U16"`
			//Ui32 uint32        `flag:"flag_U32"`
			Ui64 uint64 `flag:"flag_U64,1"`
			//F32  float32       `flag:"flag_F32"`
			F64 float64       `flag:"flag_F64,1"`
			B   bool          `flag:"flag_B,1"`
			T   time.Duration `flag:"flag_T,1s"`
		}
	}

	v := &T{}

	err := Revise(v)
	if err != nil {
		t.Errorf("TestFlagDefault %s", err.Error())
	}

	fmt.Printf("TestflagDefault %#v\n", v)
}
