//go:build freebsd && amd64
// +build freebsd,amd64

package assets

import (
    _ "embed"
)

//go:embed node_freebsd_amd64.zst
var EmbeddedNode []byte