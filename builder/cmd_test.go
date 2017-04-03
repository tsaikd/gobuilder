package builder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getAllSubPackagesRelDir(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	dirs, err := getAllSubPackagesRelDir("")
	require.NoError(err)
	require.Len(dirs, 1)
	require.Equal(dirs[0], ".")
	cmdArgs := getAllCmdArgsForPackages(dirs)
	require.Len(cmdArgs, 1)
	require.Equal(cmdArgs[0], ".")

	dirs, err = getAllSubPackagesRelDir("..")
	require.NoError(err)
	require.NotEmpty(dirs)
	require.True(len(dirs) > 5)
	cmdArgs = getAllCmdArgsForPackages(dirs)
	require.NotEmpty(cmdArgs)
}
