package filesystem_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wengchaoxi/terraform-provider-filesystem/filesystem"
)

func TestFileSystem(t *testing.T) {
	fm := filesystem.FileManager{}
	path := "./test.txt"
	t.Run("create", func(t *testing.T) {
		fm.CreateFile(path, "Hello World")
		fi, _ := fm.ReadFile(path)
		require.Equal(t, "Hello World", fi.FileContent)
	})

	t.Run("update", func(t *testing.T) {
		fm.UpdateFile(path, "你好世界")
		fi, _ := fm.ReadFile(path)
		require.Equal(t, "你好世界", fi.FileContent)
	})

	t.Run("delete", func(t *testing.T) {
		err := fm.DeleteFile(path)
		require.NoError(t, err, "delete success.")
	})
}
