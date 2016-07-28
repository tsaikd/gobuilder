package errorcheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Example() {
	if err := Check("github.com/tsaikd/gobuilder"); err != nil {
		fmt.Println(err)
	}
	// Output:
}

func Example_crosspkgtest() {
	if err := Check("crosspkgtest"); err != nil {
		fmt.Println(err)
	}
	// Output:
}

func Test_errortest(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	err := Check("errortest")
	require.Error(err)
}
