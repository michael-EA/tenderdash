package types

import (
	"encoding/json"
	"fmt"
	"net"

	"inet.af/netaddr"
)

type IPAddress struct {
	data netaddr.IP
}

// Returns underlying netaddr.IP
func (ip IPAddress) ToNetaddrIP() netaddr.IP {
	return ip.data
}

// IPAddr returns the net.IPAddr representation of an IP.
// The returned value is always non-nil, but the IPAddr.
// IP will be nil if ip is the zero value.
// If ip contains a zone identifier, IPAddr.Zone is populated.
func (ip IPAddress) ToIPAddr() *net.IPAddr {
	return ip.data.IPAddr()
}

// METHODS REQUIRED BY GOGO PROTOBUF

func (ip IPAddress) Marshal() ([]byte, error) {
	data := make([]byte, net.IPv6len)
	n, err := ip.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	if n != net.IPv6len {
		return nil, fmt.Errorf("length mismatch: want %d, got %d", net.IPv6len, n)
	}

	return data, nil
}

// MarshalTo puts binary form of IP Address into data slice.
// Data needs 4 bytes for ipv4 address and 16 bytes for ipv6 address
func (ip *IPAddress) MarshalTo(data []byte) (n int, err error) {
	ret := ip.data.As16()

	if len(data) < len(ret) {
		return 0, fmt.Errorf("IP address requires at least %d bytes, has only %d", len(ret), len(data))
	}

	return copy(data, ret[:]), nil
}

// Unmarshal parses BigEndian-encoded IPv6 address.
// It also supports IPv4 addresses mapped to IPv6.
func (ip *IPAddress) Unmarshal(data []byte) error {
	if len(data) != net.IPv6len {
		return fmt.Errorf("invalid length of data to unmarshal: have %d, want %d", len(data), net.IPv6len)
	}

	var ipaddr [16]byte
	copy(ipaddr[:], data)
	ip.data = netaddr.IPFrom16(ipaddr)
	return nil
	// ip.high = binary.BigEndian.Uint64(data[:8])
	// ip.low = binary.BigEndian.Uint64(data[8:])
	// return nil
}

func (ip *IPAddress) Size() int {
	return net.IPv6len
}

func (ip IPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.data.String())
}
func (ip *IPAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	newIP, err := netaddr.ParseIP(s)
	if err != nil {
		return err
	}

	ip.data = newIP
	return nil
}

// Compare returns an integer comparing two IPs.
// The result will be 0 if ip==ip2, -1 if ip < ip2, and +1 if ip > ip2.
// The definition of "less than" is the same as the IP.Less method.
func (ip IPAddress) Compare(other IPAddress) int {
	return ip.data.Compare(other.data)
}

// Equal returns true if both addresses are equal, false otherwise
func (ip IPAddress) Equal(other IPAddress) bool {
	return (ip.Compare(other) == 0)
}

// only required if populate option is set
// func NewPopulatedT(r randyThetest) *IPAddress {}
