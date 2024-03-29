package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

// Erebrus Version
var Version = "1.0"

//Hostname
var hostname, _ = os.Hostname()

var (
	// RegexpWalletEth check valid Eth Wallet Address
	RegexpWalletEth = regexp.MustCompile("^0x[a-fA-F0-9]{40}$")
)

//add wallet regex

// StandardFields for logger
var StandardFields = log.Fields{
	"hostname": "host-server",
	"appname":  "erebrus",
}

var StandardFieldsGRPC = log.Fields{
	"hostname": hostname,
	"appname":  "erebrus",
	"service":  "gRPC",
}

// CheckError for checking any errors
func CheckError(message string, err error) {
	if err != nil {
		log.WithFields(StandardFields).Fatalf("%s %+v", message, err.Error())
	}
}

// LogError for logging any errors
func LogError(message string, err error) {
	if err != nil {
		log.WithFields(StandardFields).Warnf("%s %+v", message, err)
	}
}

// ReadFile file content
func ReadFile(path string) (bytes []byte, err error) {
	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// WriteFile content to file
func WriteFile(path string, bytes []byte) (err error) {
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FileExists check if file exists
func FileExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists check if directory exists
func DirectoryExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// GetAvailableIP search for an available ip in cidr against a list of reserved ips
func GetAvailableIP(cidr string, reserved []string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	// this two addresses are not usable
	broadcastAddr := BroadcastAddr(ipnet).String()
	networkAddr := ipnet.IP.String()

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ok := true
		address := ip.String()
		for _, r := range reserved {
			if address == r {
				ok = false
				break
			}
		}
		if ok && address != networkAddr && address != broadcastAddr {
			return address, nil
		}
	}

	return "", errors.New("no more available address from cidr")
}

// IsIPv6 check if given ip is IPv6
func IsIPv6(address string) bool {
	ip := net.ParseIP(address)
	if ip == nil {
		return false
	}
	return ip.To4() == nil
}

// IsValidIP check if ip is valid
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsValidCidr check if CIDR is valid
func IsValidCidr(cidr string) bool {
	_, _, err := net.ParseCIDR(cidr)
	return err == nil
}

// GetIPFromCidr get ip from cidr
func GetIPFromCidr(cidr string) (string, error) {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	return ip.String(), nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// BroadcastAddr returns the last address in the given network, or the broadcast address.
func BroadcastAddr(n *net.IPNet) net.IP {
	// The golang net package doesn't make it easy to calculate the broadcast address. :(
	var broadcast net.IP
	if len(n.IP) == 4 {
		broadcast = net.ParseIP("0.0.0.0").To4()
	} else {
		broadcast = net.ParseIP("::")
	}
	for i := 0; i < len(n.IP); i++ {
		broadcast[i] = n.IP[i] | ^n.Mask[i]
	}
	return broadcast
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
