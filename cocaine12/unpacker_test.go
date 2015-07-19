package cocaine12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tString = "string"
	tUint   = 100
	tInt    = -100
)

var (
	tByte    = []byte("byte")
	tBool    = true
	tFloat64 = 0.5
	in       = []interface{}{tString, tUint, tByte, tInt, tBool, tFloat64}
)

func TestUnpackerToStruct(t *testing.T) {
	var out struct {
		String  string
		Uint8   int
		Byte    []byte
		Int     int
		Bool    bool
		Float64 float64
	}
	err := unpackPayload(in, &out)
	assert.NoError(t, err)
	assert.Equal(t, "string", out.String)
	assert.Equal(t, 100, out.Uint8)
	assert.Equal(t, []byte("byte"), out.Byte)
	assert.Equal(t, -100, out.Int)
	assert.Equal(t, 0.5, out.Float64)
	assert.True(t, out.Bool)

	t.Logf("%+v", out)
}

func TestUnpackerToSlice(t *testing.T) {
	t.FailNow()
}

func TestUnpackerSliceNotEnoughFields(t *testing.T) {
	t.FailNow()
}

func TestUnpackerSliceNotEnoughValues(t *testing.T) {
	t.FailNow()
}

func TestUnpackerSliceNotAssignable(t *testing.T) {
	t.FailNow()
}

func TestUnpackerNotAPointer(t *testing.T) {
	in := []interface{}{}
	var out struct{}
	err := unpackPayload(in, out)
	assert.EqualError(t, err, ErrNotAPointer.Error())
}

func TestUnpackerStuctNotEnoughFields(t *testing.T) {
	var out struct {
		NotString float64
	}

	err := unpackPayload(in, &out)
	assert.EqualError(t, err, ErrNotEnoughFields.Error())
}

func TestUnpackerStuctNotEnoughValues(t *testing.T) {
	var out struct {
		I, J, K, L, M, N, O, P, R int
	}

	err := unpackPayload(in, &out)
	assert.EqualError(t, err, ErrNotEnoughValues.Error())
}

func TestUnpackerStructNotAssignable(t *testing.T) {
	var out struct {
		I, J, K, L, M, N int
	}
	err := unpackPayload(in, &out)
	assert.EqualError(t, err, "(field 0) int is not assignable to string")
}
