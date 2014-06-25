package padding

import (
	"reflect"
	"testing"
)

func TestPad(t *testing.T) {
	type testcase struct {
		length int
		input  []byte
		expect []byte
	}
	var cases = []testcase{
		{5, []byte{0x00, 0x00, 0x00}, []byte{0x00, 0x00, 0x00, 0x02, 0x02}},
		{3, []byte{0x00, 0x00, 0x00}, []byte{0x00, 0x00, 0x00, 0x03, 0x03, 0x03}},
		{256, []byte{0x00, 0x00, 0x00}, nil},
	}

	for k, v := range cases {
		out := Pad(v.input, v.length)
		if !reflect.DeepEqual(out, v.expect) {
			t.Errorf("expect %#v but got %#v in %d", v.expect, out, k)
			continue
		}
	}

}

func TestUnpad(t *testing.T) {
	type testcase struct {
		length int
		input  []byte
		expect []byte
	}
	var cases = []testcase{
		{5, []byte{0x00, 0x00, 0x00, 0x02, 0x02}, []byte{0x00, 0x00, 0x00}},
		{3, []byte{0x00, 0x00, 0x00, 0x03, 0x03, 0x03}, []byte{0x00, 0x00, 0x00}},
		{3, []byte{0x00, 0x00, 0x00, 0xff, 0xff, 0xff}, nil},
		{3, []byte{0x00, 0x00, 0x00, 0x03, 0x03, 0x02}, nil},
		{256, []byte{0x00, 0x00, 0x00}, nil},
	}

	for k, v := range cases {
		out := Unpad(v.input, v.length)
		if !reflect.DeepEqual(out, v.expect) {
			t.Errorf("expect %#v but got %#v in %d", v.expect, out, k)
			continue
		}
	}
}
