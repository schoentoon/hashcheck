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
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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
var sha1hash *string = flag.String("sha1", "", "A sha1 hash to check against")
var sha256hash *string = flag.String("sha256", "", "A sha256 hash to check against")
var sha512hash *string = flag.String("sha512", "", "A sha512 hash to check against")

func checkhashes(input io.Reader) (fails uint8) {
	var md5writer hash.Hash = nil
	if *md5hash != "" {
		md5writer = md5.New()
	}
	var sha1writer hash.Hash = nil
	if *sha1hash != "" {
		sha1writer = sha1.New()
	}
	var sha256writer hash.Hash = nil
	if *sha256hash != "" {
		sha256writer = sha256.New()
	}
	var sha512writer hash.Hash = nil
	if *sha512hash != "" {
		sha512writer = sha512.New()
	}
	hashwriter := NilSafeMultiWriter(md5writer, sha1writer, sha256writer, sha512writer)
	io.Copy(hashwriter, input)
	if md5writer != nil {
		md5output := hex.EncodeToString(md5writer.Sum(nil))
		if md5output != *md5hash {
			fmt.Printf("md5 mismatch, expected: %s, got: %s\n", *md5hash, md5output)
			fails++
		}
	}
	if sha1writer != nil {
		sha1output := hex.EncodeToString(sha1writer.Sum(nil))
		if sha1output != *sha1hash {
			fmt.Printf("sha1 mismatch, expected: %s, got: %s\n", *sha1hash, sha1output)
			fails++
		}
	}
	if sha256writer != nil {
		sha256output := hex.EncodeToString(sha256writer.Sum(nil))
		if sha256output != *sha256hash {
			fmt.Printf("sha256 mismatch, expected: %s, got: %s\n", *sha256hash, sha256output)
			fails++
		}
	}
	if sha512writer != nil {
		sha512output := hex.EncodeToString(sha512writer.Sum(nil))
		if sha512output != *sha512hash {
			fmt.Printf("sha512 mismatch, expected: %s, got: %s\n", *sha512hash, sha512output)
			fails++
		}
	}
	return
}
