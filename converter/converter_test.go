package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerToString(t *testing.T) {
	x := "x"
	y := PointerToString(&x)
	assert.Equal(t, x, y, "value variable \"x\" should be equal with \"y\"")
}
func TestPointerToString_Empty(t *testing.T) {
	y := PointerToString(nil)
	assert.Equal(t, "", y, "value should be empty")
}

func TestStringToPointer(t *testing.T) {
	x := "x"
	y := StringToPointer(x)
	assert.NotNil(t, y)
}

func TestStringToPointer_Nil(t *testing.T) {
	x := ""
	y := StringToPointer(x)
	assert.Nil(t, y)
}

func TestIntToPointer(t *testing.T) {
	x := 1
	y := IntToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToInt(t *testing.T) {
	x := 1
	y := PointerToInt(&x)
	assert.Equal(t, x, y)
}

func TestPointerToInt_Zero(t *testing.T) {
	y := PointerToInt(nil)
	assert.Equal(t, 0, y)
}

func TestInt8ToPointer(t *testing.T) {
	x := int8(1)
	y := Int8ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToInt8(t *testing.T) {
	x := int8(1)
	y := PointerToInt8(&x)
	assert.Equal(t, x, y)
}

func TestPointerToInt8_Zero(t *testing.T) {
	y := PointerToInt8(nil)
	assert.Equal(t, int8(0), y)
}

func TestInt16ToPointer(t *testing.T) {
	x := int16(1)
	y := Int16ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToInt16(t *testing.T) {
	x := int16(1)
	y := PointerToInt16(&x)
	assert.Equal(t, x, y)
}

func TestPointerToInt16_Zero(t *testing.T) {
	y := PointerToInt16(nil)
	assert.Equal(t, int16(0), y)
}

func TestInt32ToPointer(t *testing.T) {
	x := int32(1)
	y := Int32ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToInt32(t *testing.T) {
	x := int32(1)
	y := PointerToInt32(&x)
	assert.Equal(t, x, y)
}

func TestPointerToInt32_Zero(t *testing.T) {
	y := PointerToInt32(nil)
	assert.Equal(t, int32(0), y)
}

func TestInt64ToPointer(t *testing.T) {
	x := int64(1)
	y := Int64ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToInt64(t *testing.T) {
	x := int64(1)
	y := PointerToInt64(&x)
	assert.Equal(t, x, y)
}

func TestPointerToInt64_Zero(t *testing.T) {
	y := PointerToInt64(nil)
	assert.Equal(t, int64(0), y)
}

func TestUintToPointer(t *testing.T) {
	x := uint(1)
	y := UintToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToUint(t *testing.T) {
	x := uint(1)
	y := PointerToUint(&x)
	assert.Equal(t, x, y)
}

func TestPointerToUint_Zero(t *testing.T) {
	y := PointerToUint(nil)
	assert.Equal(t, uint(0), y)
}

func TestUint8ToPointer(t *testing.T) {
	x := uint8(1)
	y := Uint8ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToUint8(t *testing.T) {
	x := uint8(1)
	y := PointerToUint8(&x)
	assert.Equal(t, x, y)
}

func TestPointerToUint8_Zero(t *testing.T) {
	y := PointerToUint8(nil)
	assert.Equal(t, uint8(0), y)
}

func TestUint64ToPointer(t *testing.T) {
	x := uint64(1)
	y := Uint64ToPointer(x)
	assert.NotNil(t, y)
}

func TestPointerToUint64(t *testing.T) {
	x := uint64(1)
	y := PointerToUint64(&x)
	assert.Equal(t, x, y)
}

func TestPointerToUint64_Zero(t *testing.T) {
	y := PointerToUint64(nil)
	assert.Equal(t, uint64(0), y)
}

func TestFloat64ToPointer(t *testing.T) {
	f := 1.23334567
	y := Float64ToPointer(f)
	assert.NotNil(t, y)
}

func TestPointerToFloat64(t *testing.T) {
	f := 1.23334567
	y := PointerToFloat64(&f)
	assert.Equal(t, f, y)
}

func TestPointerToFloat64_Nil(t *testing.T) {
	y := PointerToFloat64(nil)
	assert.Equal(t, float64(0), y)
}

func TestFloat64ToString(t *testing.T) {
	f := 1.2
	g := Float64ToString(f)
	assert.Equal(t, "1.2", g)
}

func TestToString(t *testing.T) {
	val := 123
	r := ToString(val)

	assert.Equal(t, "123", r)
}
