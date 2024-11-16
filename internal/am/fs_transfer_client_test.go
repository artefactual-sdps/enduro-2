package am_test

import (
	"context"
	"os"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"

	"github.com/artefactual-sdps/enduro/internal/am"
)

func TestFSTransferClientUpload(t *testing.T) {
	srcDir := fs.NewDir(t, "enduro", fs.WithFile("small.txt", "I'm a smol file."))
	dstDir := fs.NewDir(t, "enduro")

	src := srcDir.Join("small.txt")
	r, err := os.Open(src)
	if err != nil {
		t.Fatalf("Couldn't open: %s", src)
	}

	c := am.NewFSTransferClient(srcDir.Path())
	path, _, err := c.Upload(context.Background(), r, "small.txt")

	assert.NilError(t, err)
	assert.Equal(t, path, dstDir.Join("small.txt"))
}
