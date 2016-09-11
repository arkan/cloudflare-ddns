package ddns

import (
	"errors"
	"net"

	"github.com/miekg/dns"
)

// errNoResult is used when no result can be returned
var errNoResult = errors.New("No result found")

// getIPAddress returns the current public ip
func getIPAddress() (*net.IP, error) {
	target := "myip.opendns.com"
	server := "resolver1.opendns.com"

	client := dns.Client{}
	requestMessage := dns.Msg{}
	requestMessage.SetQuestion(target+".", dns.TypeA)
	responseMessage, _, err := client.Exchange(&requestMessage, server+":53")
	if err != nil {
		return nil, err
	}

	if len(responseMessage.Answer) == 0 {
		return nil, err
	}

	for _, answer := range responseMessage.Answer {
		switch answer.(type) {
		case *dns.A:
			answerA := answer.(*dns.A)
			return &answerA.A, nil
		}
	}

	return nil, errNoResult
}

// getIPAddressFromHost returns the IP for a given host
func getIPAddressFromHost(host string) (*net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, errNoResult
	}

	return &ips[0], nil
}
