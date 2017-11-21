package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 1. Authenticate
// 2. Query all addresses
// 3. Erase all reverse IP records
// 4. Write all new records

var BaseURL = ""

type Subnets struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID                    string `json:"id"`
		Subnet                string `json:"subnet"`
		Mask                  string `json:"mask"`
		SectionID             string `json:"sectionId"`
		Description           string `json:"description"`
		LinkedSubnet          string `json:"linked_subnet"`
		FirewallAddressObject string `json:"firewallAddressObject"`
		VrfID                 string `json:"vrfId"`
		MasterSubnetID        string `json:"masterSubnetId"`
		AllowRequests         string `json:"allowRequests"`
		VlanID                string `json:"vlanId"`
		ShowName              string `json:"showName"`
		Device                string `json:"device"`
		Permissions           string `json:"permissions"`
		PingSubnet            string `json:"pingSubnet"`
		DiscoverSubnet        string `json:"discoverSubnet"`
		DNSrecursive          string `json:"DNSrecursive"`
		DNSrecords            string `json:"DNSrecords"`
		NameserverID          string `json:"nameserverId"`
		ScanAgent             string `json:"scanAgent"`
		IsFolder              string `json:"isFolder"`
		IsFull                string `json:"isFull"`
		Tag                   string `json:"tag"`
		Threshold             string `json:"threshold"`
		Location              string `json:"location"`
		EditDate              string `json:"editDate"`
		LastScan              string `json:"lastScan"`
		LastDiscovery         string `json:"lastDiscovery"`
		ResolveDNS            string `json:"resolveDNS"`
	} `json:"data"`
	Time float64 `json:"time"`
}

type Addresses struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID                    string `json:"id"`
		SubnetID              string `json:"subnetId"`
		IP                    string `json:"ip"`
		IsGateway             string `json:"is_gateway"`
		Description           string `json:"description"`
		Hostname              string `json:"hostname"`
		Mac                   string `json:"mac"`
		Owner                 string `json:"owner"`
		Tag                   string `json:"tag"`
		DeviceID              string `json:"deviceId"`
		Location              string `json:"location"`
		Port                  string `json:"port"`
		Note                  string `json:"note"`
		LastSeen              string `json:"lastSeen"`
		ExcludePing           string `json:"excludePing"`
		PTRignore             string `json:"PTRignore"`
		PTR                   string `json:"PTR"`
		FirewallAddressObject string `json:"firewallAddressObject"`
		EditDate              string `json:"editDate"`
	} `json:"data"`
	Time float64 `json:"time"`
}

type AuthenticationToken struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Token   string `json:"token"`
		Expires string `json:"expires"`
	} `json:"data"`
	Time float64 `json:"time"`
}

func main() {
	authenticate("admin", "*****")
}

func authenticate(username string, password string) error {

	url := fmt.Sprintf("%s/dns/user/", BaseURL)

	ca := x509.NewCertPool()

	pemData, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatal("Cannot load ca.pem")
		return err
	}

	ca.AppendCertsFromPEM(pemData)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: ca,
		},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("client.Get: ", err)
		return err
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return err
	}

	defer resp.Body.Close()

	var token AuthenticationToken

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		log.Println(err)
	}

	fmt.Println("token: ", token.Data.Token)

	return nil
}
