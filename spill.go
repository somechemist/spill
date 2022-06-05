package main

import (
	"fmt"
	"github.com/likexian/whois"
	"net"
	"os"
	"strings"
)

var domain string = "domain"

func GetDomain() {
	fmt.Println("Please enter the domain:")
	fmt.Scanln(&domain)
	if domain == "\n" {
		domain = "domain"
		Help()
		os.Exit(1)
	}
}

func SwitchHandler() {
	switch os.Args[1] {
	case "-a":
		ARecords()
	case "-m":
		MailExchange()
	case "-t":
		TextRecords()
	case "-n":
		NameServers()
	case "-i":
		Info()
	case "-h":
		Help()
	case "-w":
		WhoIs()
	default:
		ARecords()
		MailExchange()
		TextRecords()
		NameServers()
		WhoIs()
	}
}

func ARecords() {
	fmt.Println(":::A Records:::")
	arecords, _ := net.LookupIP(domain)
	for _, a := range arecords {
		fmt.Println(a)
	}
}

func MailExchange() {
	fmt.Println(":::MX Records:::")
	mxrecords, _ := net.LookupMX(domain)
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}

func TextRecords() {
	fmt.Println(":::TXT Records:::")
	textrecords, _ := net.LookupTXT(domain)
	for _, txt := range textrecords {
		fmt.Println(txt)
	}
}

func NameServers() {
	fmt.Println(":::Name Servers:::")
	nameservers, _ := net.LookupNS(domain)
	for _, ns := range nameservers {
		fmt.Println(ns)
	}
}

func WhoIs() {
	fmt.Println(":::WHOIS:::")
	result, err := whois.Whois(domain)
	if err == nil {
		fmt.Println(result)
	}
}

func AllRecords() {
	ARecords()
	MailExchange()
	TextRecords()
	NameServers()
	WhoIs()
}

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

	if lengthArgs == 1 {
		GetDomain()
		AllRecords()
	} else if lengthArgs == 2 {
		if strings.Contains(os.Args[1], "-") {
			GetDomain()
			SwitchHandler()
		} else if strings.Contains(os.Args[1], ".") {
			domain = os.Args[1]
			AllRecords()
		} else {
			GetDomain()
			AllRecords()
		}
	} else if lengthArgs == 3 {
		if strings.Contains(os.Args[2], ".") {
			domain = os.Args[2]
		} else {
			fmt.Println("invalid domain, please include .com, .net, etc.")
			os.Exit(1)
		}
		SwitchHandler()
	} else {
		fmt.Println("Too many arguments given.\nPlease try again.")
	}
}
