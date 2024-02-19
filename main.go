package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	domain    string
	domainF   string
	subdomian string
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.domainF, "l", "", "The target inscope domains file")
	flag.StringVar(&cfg.domain, "d", "", "The target inscope domain")
	flag.StringVar(&cfg.subdomian, "s", "", "Subdomains file")
	flag.Parse()

	if (cfg.domainF == "" || cfg.domain == "") && cfg.subdomian == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	subdomain_got, err := readLines(cfg.subdomian)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	var inscope_domain []string
	if cfg.domainF != "" {
		inscope_domain, err = readLines(cfg.domainF)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	} else if cfg.domain != "" {
		inscope_domain = append(inscope_domain, cfg.domain)
	}

	for _, domain := range inscope_domain {
		for _, subdomain := range subdomain_got {
			if strings.HasSuffix(subdomain, "."+domain) {
				fmt.Println(subdomain)
			}
		}
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
