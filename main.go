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
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "", "The file to check against")
	printhashes := flag.Bool("printhashes", false, "Print out the hashes in the format for hashcheck instead of checking against them")
	flag.Parse()
	if *file == "" {
		fmt.Fprintf(os.Stderr, "Missing the -file flag\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	f, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	if *printhashes {
		printHashes(f)
	} else if checkHashes(f) > 0 {
		os.Exit(1)
	}
}
