package webpagetest

import (
	"fmt"
)

// Connectivity settings for test
type Connectivity struct {
	// Profile name
	Name string `json:"connectivity"`
	// Download bandwidth in Kbps
	BandwidthDown int `json:"bwDown"`
	// Upload bandwidth in Kbps
	BandwidthUp int `json:"bwUp"`
	// First-hop Round Trip Time in ms
	Latency int `json:"latency"`
	// Packet loss rate - percent of packets to drop
	PacketLossRate int `json:"plr,string"`
}

// String gives human readable string for connectivity profile
func (c Connectivity) String() string {
	return fmt.Sprintf("%v (%dKbps/%dKbps) %vms, Packet Loss %d%%",
		c.Name, c.BandwidthDown, c.BandwidthUp, c.Latency, c.PacketLossRate)
}
