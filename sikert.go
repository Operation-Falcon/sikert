package main

import (
	"flag"
	"sikert/banner"
	"sikert/sources"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//CLI for the user input
	domain := flag.String("domain", "", "domain name for subdomain enumeration")

	output := flag.String("output", "", "output filename to save the result")

	flag.Parse()

	//Banner design
	banner.Banner_design()

	wg.Add(14)
	//Source Alienvault
	go sources.Alienvault_enum(*domain, *output, &wg)
	//Source Anubis
	go sources.Anubis_enum(*domain, *output, &wg)
	//Source Binary edge
	go sources.Binary_enum(*domain, *output, &wg)
	//Source Certspotter
	go sources.Certspotter_enum(*domain, *output, &wg)
	//Source Chaos
	go sources.Chaos_enum(*domain, *output, &wg)
	//Source Crtsh
	go sources.Crtsh_enum(*domain, *output, &wg)
	//Source Recondev
	go sources.Recondev_enum(*domain, *output, &wg)
	//Source Security trails
	go sources.Security_trails_enum(*domain, *output, &wg)
	//Source Shodan
	go sources.Shodan_enum(*domain, *output, &wg)
	//Source Sonar
	go sources.Sonar_enum(*domain, *output, &wg)
	//Source Sublist3r
	go sources.Sublist3r_enum(*domain, *output, &wg)
	//Source Threatcrowd
	go sources.Threatcrowd_enum(*domain, *output, &wg)
	//Source Threatminer
	go sources.Threatminer_enum(*domain, *output, &wg)
	//Source Virus total
	go sources.Virus_total_enum(*domain, *output, &wg)
	wg.Wait()

}
