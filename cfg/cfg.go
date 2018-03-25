/*
The MIT License (MIT)

Copyright (c) 2016-2018 Ivan Dejanovic

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cfg

import (
	"fmt"
	"os"
	"strings"
)

const (
	minus        = "-"
	doubleMinus  = "--"
	empty        = ""
	usage        = "Usage: tpm <codefilename>\n\nOptions:\n  -h, --help Prints help\n  -v, --version    Prints version"
	invalidUsage = "Invalid usage. For correct usage examples please try: tpm -h"
	version      = "TPM interpreter version 0.1.0"
)

func HandleArgs() (bool, string) {
	var abort bool = true
	var codeFile string

	args := os.Args[1:]
	argc := len(args)

	for index := 0; index < argc; index++ {
		var flag string = empty
		var flagArg string = args[index]

		if strings.HasPrefix(flagArg, doubleMinus) {
			flag = strings.TrimPrefix(flagArg, doubleMinus)
		} else if strings.HasPrefix(flagArg, minus) {
			flag = strings.TrimPrefix(flagArg, minus)
		}

		if flag != empty {
			switch flag {
			case "h", "help":
				fmt.Println(usage)
			case "v", "version":
				fmt.Println(version)
			default:
				fmt.Println(invalidUsage)
			}
			return abort, codeFile
		}
	}

	if argc != 1 {
		fmt.Println(usage)
		return abort, codeFile
	}

	//If we get this far we have good data to process
	abort = false
	codeFile = args[0]

	return abort, codeFile
}
