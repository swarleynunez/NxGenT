package inputs

import (
	"github.com/swarleynunez/NxGenT/core/types"
)

// All good (100%-0%)
var EV1 = types.EvidenceInfo{
	Availability: 99,
	Latency:      15,
	Jitter:       10,
	CPUCores:     10,
	CPUFrequency: 10,
	MemoryTotal:  16,
	DiskTotal:    50,
	IDSAlert:     0,
}

// Almost all bad (20%-80%)
/*var EV2 = types.EvidenceInfo{
	Availability: float64(rand.Intn(100-91+1) + 91),
	Latency:      float64(rand.Intn(23-14+1) + 14),
	Jitter:       float64(rand.Intn(18-9+1) + 9),
	CPUCores:     float64(rand.Intn(11-2+1) + 2),
	CPUFrequency: float64(rand.Intn(11-2+1) + 2),
	MemoryTotal:  float64(rand.Intn(17-8+1) + 8),
	DiskTotal:    float64(rand.Intn(51-42+1) + 42),
	IDSAlert:     idsa,	// if rand.Intn(10) > 1 {
}*/

// Balanced (50%-50%)
/*var EV3 = types.EvidenceInfo{
	Availability: float64(rand.Intn(103-94+1) + 94),
	Latency:      float64(rand.Intn(20-11+1) + 11),
	Jitter:       float64(rand.Intn(15-6+1) + 6),
	CPUCores:     float64(rand.Intn(14-5+1) + 5),
	CPUFrequency: float64(rand.Intn(14-5+1) + 5),
	MemoryTotal:  float64(rand.Intn(20-11+1) + 11),
	DiskTotal:    float64(rand.Intn(54-45+1) + 45),
	IDSAlert:     idsa,	// if rand.Intn(10) > 4,
}*/

// All bad (0%-100%)
var EV4_1 = types.EvidenceInfo{
	Availability: 98,
	Latency:      16,
	Jitter:       11,
	CPUCores:     9,
	CPUFrequency: 9,
	MemoryTotal:  15,
	DiskTotal:    49,
	IDSAlert:     1,
}

// All bad (0%-100%)
var EV4_2 = types.EvidenceInfo{
	Availability: 1,
	Latency:      100,
	Jitter:       100,
	CPUCores:     1,
	CPUFrequency: 1,
	MemoryTotal:  1,
	DiskTotal:    1,
	IDSAlert:     1,
}

// Almost all good (80%-20%)
/*var EV5 = types.EvidenceInfo{
	Availability: float64(rand.Intn(106-97+1) + 97),
	Latency:      float64(rand.Intn(17-8+1) + 8),
	Jitter:       float64(rand.Intn(12-3+1) + 3),
	CPUCores:     float64(rand.Intn(17-8+1) + 8),
	CPUFrequency: float64(rand.Intn(17-8+1) + 8),
	MemoryTotal:  float64(rand.Intn(23-14+1) + 23),
	DiskTotal:    float64(rand.Intn(57-48+1) + 48),
	IDSAlert:     idsa,	// if rand.Intn(10) > 7
}*/

// All good changing latency (conf. A)
var EV6_1 = types.EvidenceInfo{
	Availability: 99,
	Latency:      25, // Default value: 25
	Jitter:       50,
	CPUCores:     10,
	CPUFrequency: 10,
	MemoryTotal:  16,
	DiskTotal:    50,
	IDSAlert:     0,
}

// All good changing latency (conf. B)
var EV6_2 = types.EvidenceInfo{
	Availability: 300,
	Latency:      24,
	Jitter:       0,
	CPUCores:     150,
	CPUFrequency: 150,
	MemoryTotal:  150,
	DiskTotal:    150,
	IDSAlert:     0,
}
