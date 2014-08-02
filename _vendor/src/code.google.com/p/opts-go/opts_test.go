// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opts

import (
	"os"
	"testing"
)

func TestXname(t *testing.T) {
	if Xname != os.Args[0] {
		t.Error("Xname did not match Argv[0]")
	}
}
