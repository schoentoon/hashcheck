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
	"flag"
	"fmt"
	"hash"
	"io"
)

var printmd5hash *bool = flag.Bool("no-md5", false, "Don't print the md5 hash of this file.")
var printsha1hash *bool = flag.Bool("no-sha1", false, "Don't print the sha1 hash of this file.")
var printsha224hash *bool = flag.Bool("no-sha224", false, "Don't print the sha224 hash of this file.")
var printsha256hash *bool = flag.Bool("no-sha256", false, "Don't print the sha256 hash of this file.")
var printsha384hash *bool = flag.Bool("no-sha384", false, "Don't print the sha384 hash of this file.")
var printsha512hash *bool = flag.Bool("no-sha512", false, "Don't print the sha512 hash of this file.")

func printHashes(input io.Reader) {
	var md5writer hash.Hash = nil
	if *printmd5hash == false {
		md5writer = md5.New()
	}
	var sha1writer hash.Hash = nil
	if *printsha1hash == false {
		sha1writer = sha1.New()
	}
	var sha224writer hash.Hash = nil
	if *printsha224hash == false {
		sha224writer = sha256.New224()
	}
	var sha256writer hash.Hash = nil
	if *printsha256hash == false {
		sha256writer = sha256.New()
	}
	var sha384writer hash.Hash = nil
	if *printsha384hash == false {
		sha384writer = sha512.New384()
	}
	var sha512writer hash.Hash = nil
	if *printsha512hash == false {
		sha512writer = sha512.New()
	}
	hashwriter := NilSafeMultiWriter(md5writer, sha1writer, sha224writer, sha256writer, sha384writer, sha512writer)
	io.Copy(hashwriter, input)
	if md5writer != nil {
		fmt.Printf(" -md5 %x", md5writer.Sum(nil))
	}
	if sha1writer != nil {
		fmt.Printf(" -sha1 %x", sha1writer.Sum(nil))
	}
	if sha224writer != nil {
		fmt.Printf(" -sha224 %x", sha224writer.Sum(nil))
	}
	if sha256writer != nil {
		fmt.Printf(" -sha256 %x", sha256writer.Sum(nil))
	}
	if sha384writer != nil {
		fmt.Printf(" -sha384 %x", sha384writer.Sum(nil))
	}
	if sha512writer != nil {
		fmt.Printf(" -sha512 %x", sha512writer.Sum(nil))
	}
	fmt.Print("\n")
}
