package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/joanbono/color"
)

// func ParseJSON(jsonfile string) {
// 	println(jsonfile)
// }

// Defining colors
var yellow = color.New(color.Bold, color.FgYellow).SprintfFunc()
var red = color.New(color.Bold, color.FgRed).SprintfFunc()
var cyan = color.New(color.Bold, color.FgCyan).SprintfFunc()

//var green = color.New(color.Bold, color.FgGreen).Printf()
var colorGreen = "\033[32m"

func ParseJSON(dump string) {

	dumpFile, err := os.Open(dump)
	CheckErr(err)

	defer dumpFile.Close()
	//byteValue, _ := ioutil.ReadAll(dumpFile)

	reader := bufio.NewReader(dumpFile)
	buf := make([]byte, 16)
	i := 1
	x := 0

	/*
		fmt.Printf("            ⌌-----------------------⌍\n")
		fmt.Printf("            |       Color code      |\n")
		fmt.Printf("            ⌎-----------⫟-----------⌏\n")
		fmt.Printf("            | \033[33m▶ UID\033[0m     | \033[32m▶ BCC\033[0m     |\n")
		fmt.Printf("            | \033[34m▶ SAK\033[0m     | \033[32m▶ ATQA\033[0m    |\n")
		fmt.Printf("            | \033[32m▶ A Keys\033[0m  | \033[36m▶ B Keys\033[0m  |\n")
		fmt.Printf("            ⌎-----------⫟-----------⌏\n")
	*/
	fmt.Printf("⌌----------⫟--------------⫟----------⫟--------------⌍\n")
	fmt.Printf("|  Offset  |       A      |  Access  |      B       |\n")
	fmt.Printf("⌎----------⫠--------------⫠----------⫠--------------⌏\n")

	for {

		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		//fmt.Printf("%x\n", buf)
		//println(i % 4)
		if i == 1 {
			fmt.Printf("| \033[90m%08d\033[0m | \033[33m%x\033[0m%x | %x | %x |\n", x, buf[0:4], buf[4:6], buf[6:10], buf[10:16])
		} else if i%4 == 0 {
			//fmt.Println(string(colorGreen), "test")
			//fmt.Fprintf(color.Output, "| %08x | %x | %x | %x | <-- KEYS\n", x, buf[0:6], buf[6:10], buf[10:16])
			fmt.Printf("| \033[90m%08d\033[0m | \033[32m%x\033[0m | \033[31m%x\033[0m | \033[34m%x\033[0m |\n", x, buf[0:6], buf[6:10], buf[10:16])
		} else {
			fmt.Printf("| \033[90m%08d\033[0m | %x | %x | %x |\n", x, buf[0:6], buf[6:10], buf[10:16])
		}
		i = i + 1
		x = x + 16
	}

	fmt.Printf("⌎----------⫠--------------⫠----------⫠--------------⌏\n\n")

}

func KeyPrinter(test string) {
	println(test)
}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
