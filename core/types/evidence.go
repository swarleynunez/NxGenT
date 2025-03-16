package types

import (
	"net"
)

type EvidenceType uint8

const (
	LatencyEvidence EvidenceType = iota
	ResourcesEvidence
	IDSEvidence
)

type RPCEvidence struct {
	Type   EvidenceType
	Target net.IP
	Info   *RPCEvidenceInfo
}

type RPCEvidenceInfo struct {
	LostPacket    bool               `json:"lost_pkt,omitempty"`
	PacketLatency float64            `json:"pkt_latency,omitempty"`
	CPUCores      map[uint64]float64 `json:"cpu_cores,omitempty"`
	MemoryTotal   uint64             `json:"memory_total,omitempty"`
	DiskTotal     uint64             `json:"disk_total,omitempty"`
	IDSAlert      bool               `json:"ids_alert,omitempty"`
}

type Evidence struct {
	Target net.IP
	Info   *EvidenceInfo
}

type EvidenceInfo struct {
	Availability float64 `json:"availability"`  // %
	Latency      float64 `json:"latency"`       // In ms
	Jitter       float64 `json:"jitter"`        // In ms
	CPUCores     float64 `json:"cpu_cores"`     // # of logical cores
	CPUFrequency float64 `json:"cpu_frequency"` // CPU frequency in GHz
	MemoryTotal  float64 `json:"memory_total"`  // In GB
	DiskTotal    float64 `json:"disk_total"`    // In GB
	IDSAlert     float64 `json:"ids_alert"`     // Is there an alert? 0 or 1
}
