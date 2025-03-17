package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swarleynunez/NxGenT/core/bindings"
	"math/big"
)

type SLA struct {
	Customer     common.Address
	Metrics      []bindings.TypesSLAMetric
	SetAt        *big.Int
	AcceptedAt   *big.Int
	TerminatedAt *big.Int
}
