/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Defining colors
var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var hired = color.New(color.FgHiRed)
var magenta = color.New(color.FgMagenta)
var gray = color.New(color.FgHiBlack)
var green = color.New(color.FgGreen)
var cyan = color.New(color.FgCyan)
var blue = color.New(color.FgBlue)
var white = color.New(color.Bold, color.FgWhite)

// ParseDump will open the Mifare Dump
// and print it in a readable way
func ParseDump(dump string, keys bool) {
	dumpFile, err := os.Open(dump)
	CheckErr(err)

	defer dumpFile.Close()

	reader := bufio.NewReader(dumpFile)
	buf := make([]byte, 16)
	i := 1
	x := 0
	var keyDictionary []string
	var uid string

	fmt.Printf("⌌----------⫟--------------⫟----------⫟--------------⌍\n")
	fmt.Printf("|  %v  |      %v       |  %v  |      %v       |\n", white.Sprintf("%v", "Offset"), white.Sprintf("%v", "A"), white.Sprintf("%v", "Access"), white.Sprintf("%v", "B"))
	fmt.Printf("⌎----------⫠--------------⫠----------⫠--------------⌏\n")
	for {

		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if i == 1 {
			fmt.Printf("| %v | %v%v%v | %v%x | %x |\n", gray.Sprintf("%08x", x), yellow.Sprintf("%08x", buf[0:4]), cyan.Sprintf("%2x", buf[4]), hired.Sprintf("%02x", buf[5]), magenta.Sprintf("%02x", buf[6]), buf[7:10], buf[10:16])
			uid = fmt.Sprintf("%08x", buf[0:4])
		} else if i%4 == 0 {
			fmt.Printf("| %v | %v | %v | %v |\n", gray.Sprintf("%08x", x), green.Sprintf("%x", buf[0:6]), red.Sprintf("%x", buf[6:10]), blue.Sprintf("%x", buf[10:16]))
			keyDictionary = append(keyDictionary, fmt.Sprintf("%x", buf[0:6]), fmt.Sprintf("%x", buf[10:16]))
		} else {
			fmt.Printf("| %v | %x | %x | %x |\n", gray.Sprintf("%08x", x), buf[0:6], buf[6:10], buf[10:16])
		}
		i = i + 1
		x = x + 16
	}

	fmt.Printf("⌎----------⫠--------------⫠----------⫠--------------⌏\n\n")

	if keys {
		SaveKeys(keyDictionary, uid)
	}

}

// SaveKeys will store the keys of a dumo
// into a file named UID-keys.dic
func SaveKeys(keyDictionary []string, uid string) {
	uniqueKeys := RemoveDuplicates(keyDictionary)

	file, err := os.Create(uid + "-key.dic")
	CheckErr(err)
	defer file.Close()
	for _, key := range uniqueKeys {
		fmt.Fprintf(file, "%v\n", key)
	}
	fmt.Printf("%v Keys saved into %v\n\n", green.Sprintf("[+]"), white.Sprintf("%v-keys.dic", uid))
}

// CodeColor prints the legent
func CodeColor() {
	fmt.Printf("            ⌌-----------------------⌍\n")
	fmt.Printf("            |     %v       |\n", white.Sprintf("%v", "Color Codes"))
	fmt.Printf("            ⌎-----------⫟-----------⌏\n")
	fmt.Printf("            | %v       %v     |\n", yellow.Sprintf("%v", "▶ UID"), cyan.Sprintf("%v", "▶ BCC"))
	fmt.Printf("            | %v      %v  |\n", magenta.Sprintf("%v", "▶ ATQA"), green.Sprintf("%v", "▶ A Keys"))
	fmt.Printf("            | %v       %v  |\n", hired.Sprintf("%v", "▶ SAK"), blue.Sprintf("%v", "▶ B Keys"))
	fmt.Printf("            | %v         |\n", red.Sprintf("%v", "▶ Access Bits"))
	fmt.Printf("            ⌎-----------⫟-----------⌏\n")
}

// RemoveDuplicates will remove the duplicate
// keys found in the dump
func RemoveDuplicates(keyDictionary []string) []string {
	keys := make(map[string]bool)
	keyList := []string{}
	for _, entry := range keyDictionary {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			keyList = append(keyList, entry)
		}
	}
	return keyList
}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
