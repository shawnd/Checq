package utils

import (
	"net"
	"strings"
)

//
// GetIPAdress queries the given domain name for IP records
//
func GetIPAdress(name string) []string {
	res, err := net.LookupIP(name)

	var results []string
	for _, ip := range res {
		if !stringInSlice(ip.String(), results) {
			results = append(results, ip.String())
		}
	}

	if err == nil {
		return results
	}
	return nil
}

//
// QueryIPAdress queries the given domain name for IP records
//
func QueryIPAdress(name string) []string {
	res, err := net.LookupIP(name)

	var results []string
	for _, ip := range res {
		if !stringInSlice(ip.String(), results) {
			results = append(results, ip.String())
		}
	}

	if err == nil {
		return results
	}
	return nil
}

//
// QueryTXT queries the given domain name for txt records.
// When www is given as a subdomain it is removed
//
func QueryTXT(name string) []string {
	name = strings.TrimLeft(name, "www.")
	res, err := net.LookupTXT(name)

	var results []string
	for _, value := range res {
		if !stringInSlice(value, results) {
			results = append(results, value)
		}
	}

	if err == nil {
		return results
	}
	return nil
}

//
// QueryNS queries the given domain name for nameserver records.
//
func QueryNS(name string) []string {
	name = strings.TrimLeft(name, "www.")
	// test := strings.Split(name, ".")
	// fmt.Println(test)
	res, err := net.LookupNS(name)

	var results []string
	for _, value := range res {
		results = append(results, value.Host)
	}

	if err == nil {
		return results
	}
	return nil
}

//
// QueryHostname queries the given domain name for nameserver records.
//
func QueryHostname(name string) string {
	ip := QueryIPAdress(name)
	if ip != nil {
		res, err := net.LookupAddr(ip[0])

		if err == nil {
			return res[0]
		}
	}
	return ""

}
