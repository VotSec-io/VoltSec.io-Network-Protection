package firewall

import "net"

// Rule represents a firewall rule.
type Rule struct {
	SourceIP   net.IP
	DestIP     net.IP
	Port       int
	Protocol   string
	Action     string
}

// NewRule creates a new instance of Rule.
func NewRule(sourceIP, destIP net.IP, port int, protocol, action string) *Rule {
	return &Rule{
		SourceIP:   sourceIP,
		DestIP:     destIP,
		Port:       port,
		Protocol:   protocol,
		Action:     action,
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
