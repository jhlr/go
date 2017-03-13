package life

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"os"
// )
//
// // Main will use black/white PNG
// // reader = Stdin
// // writer = Stdout
// // iterates N times using the given rules
// func Main(rule Rule, n uint) error {
// 	uni := New(rule)
// 	img, _, err := image.Decode(os.Stdin)
// 	if err != nil {
// 		return err
// 	}
// 	uni.SetImage(img, ColorFunc)
// 	for i := uint(0); i < n; i++ {
// 		uni.Next()
// 	}
// 	res := uni.Image(color.GrayModel, GrayFunc)
// 	err = png.Encode(os.Stdout, res)
// 	return err
// }
