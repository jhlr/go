package util

// Catch will keep the given function from panicking
// and returns the pacicked object or nil
func Catch(try func()) (err interface{}) {
	if try != nil {
		defer func() {
			err = recover()
		}()
		try()
	}
	return nil
}
