package firewall

import "net"

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
