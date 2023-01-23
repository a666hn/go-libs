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
