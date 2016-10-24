package provide

import (
	"fmt"

	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrorProvideTest     = errutil.NewFactory("provide text")
	ErrorProvideSelfTest = errutil.NewFactory("provide self text")
)

func main() {
	fmt.Println(ErrorProvideSelfTest.New(nil))
}

// Nothing public method to do nothing
func Nothing() {}
