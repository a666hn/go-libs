package converter

import "fmt"

// StringToPointer convert string to pointer
func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// PointerToString convert pointer to string type
func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// IntToPointer convert int to pointer
func IntToPointer(i int) *int {
	return &i
}

// PointerToInt convert pointer to int type
func PointerToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Int8ToPointer convert int8 type to pointer
func Int8ToPointer(i int8) *int8 {
	return &i
}

// PointerToInt8 convert pointer to int8 type
func PointerToInt8(i *int8) int8 {
	if i == nil {
		return int8(0)
	}
	return *i
}

// Int16ToPointer convert int16 type to pointer
func Int16ToPointer(i int16) *int16 {
	return &i
}

// PointerToInt16 convert pointer to int16 type
func PointerToInt16(i *int16) int16 {
	if i == nil {
		return int16(0)
	}
	return *i
}

// Int32ToPointer convert int32 type to pointer
func Int32ToPointer(i int32) *int32 {
	return &i
}

// PointerToInt32 convert pointer to int32 type
func PointerToInt32(i *int32) int32 {
	if i == nil {
		return int32(0)
	}
	return *i
}

// Int64ToPointer convert int64 type to pointer
func Int64ToPointer(i int64) *int64 {
	return &i
}

// PointerToInt64 convert pointer to int64 type
func PointerToInt64(i *int64) int64 {
	if i == nil {
		return int64(0)
	}
	return *i
}

// UintToPointer convert uint to pointer
func UintToPointer(u uint) *uint {
	return &u
}

// PointerToUint convert Pointer to uint type
func PointerToUint(u *uint) uint {
	if u == nil {
		return uint(0)
	}
	return *u
}

// Uint8ToPointer convert uint8 type to pointer
func Uint8ToPointer(u uint8) *uint8 {
	return &u
}

// PointerToUint8 convert Pointer to uint8 type
func PointerToUint8(u *uint8) uint8 {
	if u == nil {
		return uint8(0)
	}
	return *u
}

// Uint64ToPointer convert uint64 type to pointer
func Uint64ToPointer(u uint64) *uint64 {
	return &u
}

// PointerToUint64 convert Pointer to uint64 type
func PointerToUint64(u *uint64) uint64 {
	if u == nil {
		return uint64(0)
	}
	return *u
}

// Float64ToPointer convert float64 to pointer
func Float64ToPointer(f float64) *float64 {
	return &f
}

// PointerToFloat64 convert pointer float64 to float64 type
func PointerToFloat64(f *float64) float64 {
	if f == nil {
		return float64(0)
	}
	return *f
}

// Float64ToString format float to string
func Float64ToString(f float64) string {
	return fmt.Sprintf("%g", f)
}

// ToString convert type to string
func ToString(s interface{}) string {
	return fmt.Sprintf("%v", s)
}
