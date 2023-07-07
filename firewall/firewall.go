package firewall

import (
	"fmt"
	"log"
	"net"
)

// Firewall represents a network security firewall.
type Firewall struct {
	rules []*Rule
}

// Rule represents a firewall rule.
type Rule struct {
	SourceIP   net.IP
	DestIP     net.IP
	Port       int
	Protocol   string
	Action     string
}

// NewFirewall creates a new instance of Firewall.
func NewFirewall() *Firewall {
	return &Firewall{
		rules: []*Rule{},
	}
}

// AddRule adds a new rule to the firewall.
func (fw *Firewall) AddRule(rule *Rule) {
	fw.rules = append(fw.rules, rule)
}

// Block checks if the given packet should be blocked based on the firewall rules.
func (fw *Firewall) Block(packet *Packet) bool {
	for _, rule := range fw.rules {
		if rule.Matches(packet) {
			log.Println("Packet blocked:", packet)
			return true
		}
	}
	return false
}

// Packet represents a network packet.
type Packet struct {
	SourceIP   net.IP
	DestIP     net.IP
	Port       int
	Protocol   string
}

// NewPacket creates a new instance of Packet.
func NewPacket(sourceIP, destIP net.IP, port int, protocol string) *Packet {
	return &Packet{
		SourceIP:   sourceIP,
		DestIP:     destIP,
		Port:       port,
		Protocol:   protocol,
	}
}

// Matches checks if the packet matches the firewall rule.
func (rule *Rule) Matches(packet *Packet) bool {
	if rule.SourceIP != nil && !rule.SourceIP.Equal(packet.SourceIP) {
		return false
	}
	if rule.DestIP != nil && !rule.DestIP.Equal(packet.DestIP) {
		return false
	}
	if rule.Port != 0 && rule.Port != packet.Port {
		return false
	}
	if rule.Protocol != "" && rule.Protocol != packet.Protocol {
		return false
	}
	return true
}

// Example usage:
func main() {
	// Create a new firewall instance
	fw := NewFirewall()

	// Add firewall rules
	rule1 := &Rule{
		SourceIP:   net.ParseIP("192.168.1.0"),
		DestIP:     nil,
		Port:       80,
		Protocol:   "tcp",
		Action:     "block",
	}
	fw.AddRule(rule1)

	rule2 := &Rule{
		SourceIP:   nil,
		DestIP:     net.ParseIP("10.0.0.1"),
		Port:       22,
		Protocol:   "tcp",
		Action:     "block",
	}
	fw.AddRule(rule2)

	// Create a sample packet to test the firewall
	packet := NewPacket(net.ParseIP("192.168.1.10"), net.ParseIP("10.0.0.1"), 80, "tcp")

	// Check if the packet should be blocked
	if fw.Block(packet) {
		fmt.Println("Packet blocked!")
	} else {
		fmt.Println("Packet allowed.")
	}
}
