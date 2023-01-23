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
