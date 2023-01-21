package converter

func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntToPointer(i int) *int {
	return &i
}

func PointerToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func UintToPointer(u uint) *uint {
	return &u
}

func PointerToUint(u *uint) uint {
	if u == nil {
		return uint(0)
	}
	return *u
}
