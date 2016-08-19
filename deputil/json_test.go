package deputil

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewJSON_root(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	testdir, err := filepath.Abs("")
	require.NoError(err)

	godepsJSON, err := NewJSON(testdir)
	require.NoError(err)
	require.NotEmpty(godepsJSON.Rev)
	require.NotEmpty(godepsJSON.RevTime)
	require.NotEmpty(godepsJSON.Deps)
	require.Contains(godepsJSON.ImportPath, "github.com/tsaikd/gobuilder")

	data, err := json.MarshalIndent(godepsJSON, "", "  ")
	require.NoError(err)
	require.NotContains(string(data), "github.com/tsaikd/KDGoLib")

	data, err = json.MarshalIndent(godepsJSON.Deps, "", "  ")
	require.NoError(err)
	require.NotContains(string(data), "github.com/tsaikd/gobuilder")
}

func TestNewJSON_lib(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	testdir, err := filepath.Abs("../builder")
	require.NoError(err)

	godepsJSON, err := NewJSON(testdir)
	require.NoError(err)
	require.NotEmpty(godepsJSON.Rev)
	require.NotEmpty(godepsJSON.RevTime)
	require.NotEmpty(godepsJSON.Deps)
	require.Contains(godepsJSON.ImportPath, "github.com/tsaikd/gobuilder/builder")

	data, err := json.MarshalIndent(godepsJSON, "", "  ")
	require.NoError(err)

	data, err = json.MarshalIndent(godepsJSON.Deps, "", "  ")
	require.NoError(err)
	require.NotContains(string(data), "github.com/tsaikd/gobuilder")
}
