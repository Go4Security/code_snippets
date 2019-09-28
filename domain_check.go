package main
import (
	"net"
	"fmt"
	"os"
	"github.com/akamensky/argparse"
	"golang.org/x/net/publicsuffix"
)
func main(){
	parser := argparse.NewParser("domain_check", "Validates a domain")
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	// Value assigned to a variable
	var domainname = *s
	// Double check if the domain is valid
	eTLD, err := publicsuffix.EffectiveTLDPlusOne(domainname)
	if err != nil {
		fmt.Print(err)
	} else {
		if eTLD != domainname {
			fmt.Printf("effective Domain is %s\n", eTLD)
			domainname = eTLD
		}
	}
	if validateDomainName(domainname) {
		fmt.Printf("The Domain %s is a valid Domain\n",domainname)
	} else {
		fmt.Printf("Invalid Domain %s\n",domainname)
	}
}
func validateDomainName(domain string) bool {
	nss, err := net.LookupNS(domain)
	if err != nil {
		return false
	}
	if len(nss) == 0 {
		return false
	} else {
		return true
	}
}
