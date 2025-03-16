package rpc

import (
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/manager"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/types"
	"net"
	"net/http"
)

func (s *RPC) SendEvidence(r *http.Request, args *types.RPCEvidence, result *any) error {

	// Debug
	/*fmt.Printf(
		"[%d] %s %s %s %s\n",
		time.Now().Unix(),
		r.Method,
		r.RequestURI,
		"SendEvidence",
		args.Target.String(),
	)*/

	// Store evidence locally
	manager.StoreEvidenceLocally(args)

	return nil
}

func (s *RPC) GetReputationScore(r *http.Request, args *net.IP, result *float64) (err error) {

	// Debug
	/*fmt.Printf(
		"[%d] %s %s %s %s\n",
		time.Now().Unix(),
		r.Method,
		r.RequestURI,
		"GetReputationScore",
		args.String(),
	)*/

	// Get trust node's reputation score by IP
	rs, err := manager.GetReputationScore(*args)
	if err == nil {
		*result = rs
	}

	return
}
