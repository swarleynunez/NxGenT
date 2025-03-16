package types

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/bindings"
	"math/big"
)

type SLA struct {
	Customer     common.Address
	Metrics      []bindings.TypesSLAMetric
	SetAt        *big.Int
	AcceptedAt   *big.Int
	TerminatedAt *big.Int
}
