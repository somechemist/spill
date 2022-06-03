package main

import (
	"fmt"
	"os"
)

var domain string = "domain"

func Help() {
	fmt.Printf("spill [-switch] {%s}", domain)
	fmt.Println("\nThis application allows you to search for whois and dig information on the fly")
	fmt.Println("Without any switches, It will give all available information")
	fmt.Println("Use -a for A records")
	fmt.Println("Use -m for MX records")
	fmt.Println("Use -t for TXT records")
	fmt.Println("Use -n for Name Servers")
	fmt.Println("Use -w for WHOIS lookup")
	fmt.Println("Use -h for this help menu")
	fmt.Println("Use -i for Author information")
}

func Info() {
	fmt.Println("\n      `.-::::::-.`\n  .:-::::::::::::::-:.\n  `_:::    ::    :::_`\n   .:( ^   :: ^   ):.   Written in GO\n   `:::   (..)   :::.   By Justin Powell\n   `:::::::UU:::::::`\n   .::::::::::::::::.\n   O::::::::::::::::O   06-02-2022\n   -::::::::::::::::-\n   `::::::::::::::::`\n    .::::::::::::::.\n      oO:::::::Oo\n")
}
func main() {
	var lengthArgs int = len(os.Args)
	if lengthArgs < 3 {
		fmt.Println("Please enter the domain:")
		fmt.Scanln(&domain)
		if domain == "" {
			domain = "domain"
		}
	} else {
		domain = os.Args[2]
	}
	if lengthArgs > 3 {
		fmt.Println("Too many arguments given!")
		Help()
	} else if lengthArgs > 1 {
		switch os.Args[1] {
		case "-i":
			Info()
		case "-h":
			Help()
		default:
			fmt.Println("\n\nNo recognized switch given!\n")
			Help()
		}
	} else {
		Help()
	}
}
