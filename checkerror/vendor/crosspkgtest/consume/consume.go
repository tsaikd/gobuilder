package consume

import (
	"crosspkgtest/provide"
	custom "crosspkgtest/provide2"
	"fmt"
)

func main() {
	fmt.Println(provide.ErrorProvideTest.New(nil))
	fmt.Println(custom.ErrorProvide2Test.New(nil))
}
