// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2014 Petar Maymounkov <p@gocircuit.org>

package valve

import (
	"github.com/gocircuit/circuit/kit/fs/rh"
	"github.com/gocircuit/circuit/kit/fs/namespace/file"
)

type TrySendFile struct {
	v *Valve
}

func NewTrySendFile(v *Valve) file.File {
	return &TrySendFile{v: v}
}

func (f *TrySendFile) Perm() rh.Perm {
	return 0222 // -w--w--w-
}

func (f *TrySendFile) Open(flag rh.Flag, intr rh.Intr) (rh.FID, error) {
	if flag.Attr != rh.WriteOnly {
		return nil, rh.ErrPerm
	}
	w, err := f.v.TrySend()
	if err != nil {
		return nil, err
	}
	return file.NewOpenInterruptibleWriterFile(w), nil
}

func (f *TrySendFile) Remove() error {
	return rh.ErrPerm
}
