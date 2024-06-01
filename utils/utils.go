package utils

// ToPtr 将任何类型的值转换为指针
func ToPtr[T any](value T) *T {
	return &value
}

// FromPtr 将指针转换为其指向的值
// 如果指针为空，返回该类型的零值
func FromPtr[T any](ptr *T) T {
	if ptr == nil {
		var zeroValue T
		return zeroValue
	}
	return *ptr
}
