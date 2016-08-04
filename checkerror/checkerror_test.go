package checkerror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tsaikd/KDGoLib/pkgutil"
)

func Example() {
	pkglist, err := pkgutil.FindAllSubPackages("github.com/tsaikd/gobuilder", "..")
	if err != nil {
		fmt.Println(err)
	}
	if err := Check(pkglist); err != nil {
		fmt.Println(err)
	}
	// Output:
}

func Example_crosspkgtest() {
	pkglist, err := pkgutil.FindAllSubPackages("crosspkgtest", "")
	if err != nil {
		fmt.Println(err)
	}
	if err := Check(pkglist); err != nil {
		fmt.Println(err)
	}
	// Output:
}

func Test_errortest(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pkglist, err := pkgutil.FindAllSubPackages("errortest", "")
	require.NoError(err)
	err = Check(pkglist)
	require.Error(err)
}
