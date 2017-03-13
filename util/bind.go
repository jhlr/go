package util

// Bind returns the given method as a function
// and it will be bound to the given object forever
func Bind(i interface{}, method string) interface{} {
	v := valueOf(i)
	return v.MethodByName(method).Interface()
}
