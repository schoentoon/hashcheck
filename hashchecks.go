/*  hashcheck
 *  Copyright (C) 2013  Toon Schoenmakers
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
)

type nilSafeMultiWriter struct {
	Writers []io.Writer
}

func NilSafeMultiWriter(writers ...io.Writer) io.Writer {
	return &nilSafeMultiWriter{writers}
}

func (t *nilSafeMultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.Writers {
		if w == nil {
			continue
		}
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

var md5hash *string = flag.String("md5", "", "A md5 hash to check against")

func checkhashes(input io.Reader) (fails uint8) {
	var md5writer hash.Hash = nil
	if *md5hash != "" {
		md5writer = md5.New()
	}
	hashwriter := NilSafeMultiWriter(md5writer)
	io.Copy(hashwriter, input)
	if md5writer != nil {
		md5output := hex.EncodeToString(md5writer.Sum(nil))
		if md5output != *md5hash {
			fmt.Printf("md5 mismatch, expected: %s. Got: %s\n", *md5hash, md5output)
			fails++
		}
	}
	return
}
