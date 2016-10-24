package checkerror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/KDGoLib/pkgutil"
)

func Example_gobuilder() {
	pkglist, err := pkgutil.FindAllSubPackages("github.com/tsaikd/gobuilder", "..")
	if err != nil {
		fmt.Println(err)
	}
	errs := Check(pkglist, false)
	for _, err := range errs {
		fmt.Println(err)
	}
	// Output:
}

func Example_errortest() {
	pkglist, err := pkgutil.FindAllSubPackages("errortest", "")
	if err != nil {
		fmt.Println(err)
	}
	errs := Check(pkglist, false)
	for _, err := range errs {
		fmt.Println(err)
	}
	// Output:
	// error factory "github.com/tsaikd/gobuilder/checkerror/vendor/errortest/provide" "ErrorProvideTest" declared but not used
	// error factory "github.com/tsaikd/gobuilder/checkerror/vendor/errortest/provide2" "ErrorProvide2Test" declared but not used
}

func Example_pkgtest() {
	pkglist, err := pkgutil.FindAllSubPackages("pkgtest", "")
	if err != nil {
		fmt.Println(err)
	}
	errs := Check(pkglist, false)
	for _, err := range errs {
		fmt.Println(err)
	}
	// Output:
}

func Test_errortest(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pkglist, err := pkgutil.FindAllSubPackages("errortest", "")
	require.NoError(err)
	errs := Check(pkglist, false)
	require.Len(errs, 2)
}

func Test_noerrfac(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pkglist, err := pkgutil.FindAllSubPackages("noerrfac", "")
	require.NoError(err)
	errs := Check(pkglist, false)
	require.Len(errs, 1)
	errs = Check(pkglist, true)
	require.Len(errs, 0)
}

func Test_pkgtest(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pkglist, err := pkgutil.FindAllSubPackages("pkgtest", "")
	require.NoError(err)
	errs := Check(pkglist, false)
	require.Len(errs, 0)
}

func Test_import(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	if pkg, err := defaultImporter.Import("errortest"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/errortest", pkg.Path())
	}
	if pkg, err := defaultImporter.Import("errortest/provide"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/errortest/provide", pkg.Path())
	}
	if pkg, err := defaultImporter.Import("errortest/provide2"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/errortest/provide2", pkg.Path())
	}

	if pkg, err := defaultImporter.Import("pkgtest"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/pkgtest", pkg.Path())
	}
	if pkg, err := defaultImporter.Import("pkgtest/consume"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/pkgtest/consume", pkg.Path())
	}
	if pkg, err := defaultImporter.Import("pkgtest/provide"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/pkgtest/provide", pkg.Path())
	}
	if pkg, err := defaultImporter.Import("pkgtest/provide2"); assert.NoError(err) {
		require.Equal("github.com/tsaikd/gobuilder/checkerror/vendor/pkgtest/provide2", pkg.Path())
	}
}
