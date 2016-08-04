package deputil

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewJSONRoot(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pwd, err := os.Getwd()
	require.NoError(err)
	testdir := filepath.Dir(pwd)

	godepsJSON, err := NewJSON(testdir)
	require.NoError(err)
	require.NotEmpty(godepsJSON.Deps)
	require.Contains(godepsJSON.ImportPath, "github.com/tsaikd/gobuilder")

	data, err := json.MarshalIndent(godepsJSON, "", "  ")
	require.NoError(err)

	count := strings.Count(string(data), "github.com/tsaikd/KDGoLib")
	require.Equal(int(1), count)

	data, err = json.MarshalIndent(godepsJSON.Deps, "", "  ")
	require.NoError(err)
	require.NotContains(string(data), "github.com/tsaikd/gobuilder")
}

func Test_NewJSONLib(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	pwd, err := os.Getwd()
	require.NoError(err)
	testdir := filepath.Dir(pwd)
	testdir = filepath.Join(testdir, "builder")

	godepsJSON, err := NewJSON(testdir)
	require.NoError(err)
	require.NotEmpty(godepsJSON.Deps)
	require.Contains(godepsJSON.ImportPath, "github.com/tsaikd/gobuilder/builder")

	data, err := json.MarshalIndent(godepsJSON, "", "  ")
	require.NoError(err)

	data, err = json.MarshalIndent(godepsJSON.Deps, "", "  ")
	require.NoError(err)
	require.NotContains(string(data), "github.com/tsaikd/gobuilder")
}
