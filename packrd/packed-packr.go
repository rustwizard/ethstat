// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "3fdfa968bd61164eb1bfeee245a2a4cf"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	return nil
}()
