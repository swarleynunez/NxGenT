package rpc

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
)

// Receiver
type RPC struct{}

func NewRouter() *mux.Router {

	// Create and configure a new RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Add service to the RPC server
	err := s.RegisterService(new(RPC), "")
	utils.CheckError(err, utils.ErrorMode)

	// Create and configure a new router
	r := mux.NewRouter().StrictSlash(true)

	// Register the RPC server endpoint
	r.Handle("/rpc", s)

	return r
}
