package consume

import (
	"fmt"

	"pkgtest/provide"
	custom "pkgtest/provide2"
)

func main() {
	fmt.Println(provide.ErrorProvideTest.New(nil))
	fmt.Println(custom.ErrorProvide2Test.New(nil))
}

// Nothing public method to do nothing
func Nothing() {}
