package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pad2ipfs ".."
	"github.com/fatih/color"
)

const checkIcon = "\xE2\x9C\x94 "

func main() {
	asciiart := `
                  _   ___    _____ _____  ______ _____
                 | | |__ \  |_   _|  __ \|  ____/ ____|
  _ __   __ _  __| |    ) |   | | | |__) | |__ | (___
 | '_ \ / _  |/ _  |   / /    | | |  ___/|  __| \___ \
 | |_) | (_| | (_| |  / /_   _| |_| |    | |    ____) |
 | .__/ \__,_|\__,_| |____| |_____|_|    |_|   |_____/  - cli
 | |
 |_|
		`
	color.Blue(asciiart)
	fmt.Println("							v0.0.1")
	color.Blue("https://github.com/arnaucode/pad2ipfs")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Pad Link to IPFS
	2 - IPFS hash to file
	0 - Exit cli
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Pad Link to IPFS")
			padlinkToIPFS()
			break
		case "2":
			fmt.Println("selected 2 - IPFS hash to file")
			hashToFile()
			break
		case "0":
			fmt.Println("selected 0 - exit cli")
			os.Exit(3)
			break
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}
func padlinkToIPFS() {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the pad link: ")
	link, _ := newcommand.ReadString('\n')
	link = strings.Replace(link, "\n", "", -1)

	newcommand = bufio.NewReader(os.Stdin)
	formats := `	Available formats:
		- md (by default)
		- txt
		- html
		- pdf
		- odt`
	fmt.Println(formats)
	fmt.Print("	Enter the pad format: ")
	format, _ := newcommand.ReadString('\n')
	format = strings.Replace(format, "\n", "", -1)
	if format == "" {
		format = "md"
	}
	if format != "md" && format != "txt" && format != "html" && format != "pdf" && format != "odt" {
		fmt.Println("		wrong format, using md format")
		format = "md"
	}

	hash, err := pad2ipfs.Add(link, format)
	if err != nil {
		color.Red(err.Error())
	} else {
		color.Green(checkIcon + "File added to IPFS network")
		fmt.Print("IPFS hash: ")
		color.Blue(hash)
	}
}
func hashToFile() {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the IPFS hash: ")
	hash, _ := newcommand.ReadString('\n')
	hash = strings.Replace(hash, "\n", "", -1)
	err := pad2ipfs.Get(hash, hash+".md")
	if err != nil {
		color.Red(err.Error())
	} else {

		color.Green(checkIcon + "File downloaded from IPFS network")
		fmt.Print("File stored in: ")
		color.Blue(pad2ipfs.GettedPads + "/" + hash + ".md")
	}
}
