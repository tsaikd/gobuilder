package provide

import (
	"fmt"

	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrorProvideTest     = errutil.NewFactory("provide test")
	ErrorProvideTest2    = errutil.NewFactory("provide test 2")
	ErrorProvideTest3    = errutil.NewFactory("provide test 3")
	ErrorProvideSelfTest = errutil.NewFactory("provide self test")
)

func main() {
	func(fac errutil.ErrorFactory) {
		if ErrorProvideTest2 != nil {
			fmt.Println(ErrorProvideSelfTest.New(nil))
		}
	}(ErrorProvideTest3)
}

// Nothing public method to do nothing
func Nothing() {}
