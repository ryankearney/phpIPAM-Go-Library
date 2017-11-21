package phpipamapi

// IP Address
type IP string

// GetSubnet returns a specific subnet by ID
func GetSubnet(id int) error {
	return nil
}

// GetSubnetUsage returns the subents utilization
func GetSubnetUsage(id int) error {
	return nil
}

// GetSubnetFirstFree returns the first available IP address in a subnet
func GetSubnetFirstFree(id int) error {
	return nil
}

// GetSubnetSlaves returns all children subnets of the given subnet
func GetSubnetSlaves(id int) error {
	return nil
}

// GetSubnetSlavesRecursive returns all children subnets of the given subnet and recustively returns their children as well
func GetSubnetSlavesRecursive(id int) error {
	return nil
}

// GetSubnetAddresses returns all IP addresses from the subnet
func GetSubnetAddresses(id int) error {
	return nil
}

// GetSubnetAddress returns an IP address from the subnet
func GetSubnetAddress(id int, ip IP) error {
	return nil
}

// GetFirstSubnet returns the first available subnet within the provided subnet for the given mask size
func GetFirstSubnet(id int, mask int) error {
	return nil
}

// GetAllSubnets returns all available subnets within the provided subnet for the given mask size
func GetAllSubnets(id int, mask int) error {
	return nil
}

// GetSubnetCustomFields returns all subnet custom fields
func GetSubnetCustomFields() error {
	return nil
}

// GetSubnetByCIDR searches for a subnet by CIDR format
func GetSubnetByCIDR(subnet IP) error {
	return nil
}

// CreateSubnet creates a new subnet
func CreateSubnet() error {
	return nil
}

// CreateChildSubnet creates a new child subnet inside the subnet with the specified mask
func CreateChildSubnet(id int, mask int) error {
	return nil
}

// UpdateSubnet updates a subnet
func UpdateSubnet() error {
	return nil
}

// ResizeSubnet resizes a subnet
func ResizeSubnet(id int) error {
	return nil
}

// SplitSubnet splits a subnet into smaller subnets
func SplitSubnet(id int) error {
	return nil
}

// SetSubnetPermissions sets permissions for the given subnet
func SetSubnetPermissions(id int) error {
	return nil
}

// DeleteSubnet deletes the subnet
func DeleteSubnet(id int) error {
	return nil
}

// TruncateSubnet removes all IP addresses form the subnet
func TruncateSubnet(id int) error {
	return nil
}

// DeleteSubnetPermissions removes all permissions
func DeleteSubnetPermissions(id int) error {
	return nil
}
