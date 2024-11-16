package fs_utils

import "testing"

func Test_FSDirectoryInit(t *testing.T) {
	t.Run("FSDirectory inits properly", func(t *testing.T) {
		f := NewFSDirectory("~/.config", DirOnlyDataType)
		if f.Path == "" {
			t.Error("FSDirectory initialized with an empty string.")
		}

		if (f.DataType != FileOnlyDataType) && (f.DataType != DirOnlyDataType) && (f.DataType != DirOrFileDataType) {
			t.Errorf("FSDirectory initialized with an unsupported datatype of %v", f.DataType)
		}
	})
}
