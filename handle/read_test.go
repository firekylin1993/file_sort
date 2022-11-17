package handle

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	dir := t.TempDir()
	depDir := "../records"
	dstDir := filepath.Join(dir, "result")
	reader := NewMyReader()
	err := reader.ReadFile("./records/record_0.txt", depDir, dstDir)
	assert.Nil(t, err, "database select error")
}
