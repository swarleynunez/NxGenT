package inputs

import (
	"github.com/swarleynunez/NxGenT/core/bindings"
	"github.com/swarleynunez/NxGenT/core/utils"
)

var SLA1 = []bindings.TypesSLAMetric{
	{
		Id:        "availability",
		Threshold: utils.ConvertFloatToBigInt(99),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "latency",
		Threshold: utils.ConvertFloatToBigInt(15),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "jitter",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "cpu_cores",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "cpu_frequency",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "memory_total",
		Threshold: utils.ConvertFloatToBigInt(16),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "disk_total",
		Threshold: utils.ConvertFloatToBigInt(50),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "ids_alert",
		Threshold: utils.ConvertFloatToBigInt(0),
		Weighting: utils.ConvertFloatToBigInt(0.5),
	},
}

var SLA2 = []bindings.TypesSLAMetric{
	{
		Id:        "availability",
		Threshold: utils.ConvertFloatToBigInt(99),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "latency",
		Threshold: utils.ConvertFloatToBigInt(25),
		Weighting: utils.ConvertFloatToBigInt(0.3),
	},
	{
		Id:        "jitter",
		Threshold: utils.ConvertFloatToBigInt(50),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "cpu_cores",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "cpu_frequency",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "memory_total",
		Threshold: utils.ConvertFloatToBigInt(16),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "disk_total",
		Threshold: utils.ConvertFloatToBigInt(50),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "ids_alert",
		Threshold: utils.ConvertFloatToBigInt(0),
		Weighting: utils.ConvertFloatToBigInt(0.3),
	},
}

var SLA3 = []bindings.TypesSLAMetric{
	{
		Id:        "availability",
		Threshold: utils.ConvertFloatToBigInt(99),
		Weighting: utils.ConvertFloatToBigInt(0.2),
	},
	{
		Id:        "latency",
		Threshold: utils.ConvertFloatToBigInt(25),
		Weighting: utils.ConvertFloatToBigInt(0.05),
	},
	{
		Id:        "jitter",
		Threshold: utils.ConvertFloatToBigInt(50),
		Weighting: utils.ConvertFloatToBigInt(0.2),
	},
	{
		Id:        "cpu_cores",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.15),
	},
	{
		Id:        "cpu_frequency",
		Threshold: utils.ConvertFloatToBigInt(10),
		Weighting: utils.ConvertFloatToBigInt(0.1),
	},
	{
		Id:        "memory_total",
		Threshold: utils.ConvertFloatToBigInt(16),
		Weighting: utils.ConvertFloatToBigInt(0.15),
	},
	{
		Id:        "disk_total",
		Threshold: utils.ConvertFloatToBigInt(50),
		Weighting: utils.ConvertFloatToBigInt(0.15),
	},
	{
		Id:        "ids_alert",
		Threshold: utils.ConvertFloatToBigInt(0),
		Weighting: utils.ConvertFloatToBigInt(0),
	},
}
