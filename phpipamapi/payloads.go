package phpipamapi

import (
	"fmt"
)

type convertibleBoolean bool

func (bit *convertibleBoolean) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == "true" {
		*bit = true
	} else if s == "0" || s == "false" {
		*bit = false
	} else {
		return fmt.Errorf("Boolean unmarshal error: invalid input %s", s)
	}
	return nil
}

type revokedSession struct {
	Code    int                `json:"code"`
	Success convertibleBoolean `json:"success"`
	Message string             `json:"message,omitempty"`
	Data    []string           `json:"data,omitempty"`
	Time    float64            `json:"time"`
}

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
