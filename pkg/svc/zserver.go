package svc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/seizadi/cost-insights-backend/pkg/pb"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~ A BRIEF DEVELOPMENT GUIDE ~~~~~~~~~~~~~~~~~~~~~~~~~
//
// TODO: Extend the AwsCost service by defining new RPCs and
// and message types in the pb/service.proto file. These RPCs and messages
// compose the API for your service. After modifying the proto schema in
// pb/service.proto, call "make protobuf" to regenerate the protobuf files.
//
// TODO: Create an implementation of the AwsCost server
// interface. This interface is generated by the protobuf compiler and exists
// inside the pb/service.pb.go file. The "server" struct already provides an
// implementation of AwsCost server interface service, but only
// for the GetVersion function. You will need to implement any new RPCs you
// add to your protobuf schema.
//
// TODO: Update the GetVersion function when newer versions of your service
// become available. Feel free to change GetVersion to better-suit how your
// versioning system, or get rid of it entirely. GetVersion helps make up a
// simple "starter" example that allows an end-to-end example. It is not
// required.
//
// TODO: Oh yeah, delete this guide when you no longer need it.
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~ FAREWELL AND GOOD LUCK ~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	// version is the current version of the service
	version = "0.0.1"
)

// Default implementation of the AwsCost server interface
type server struct{}

// GetVersion returns the current version of the service
func (server) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

// NewBasicServer returns an instance of the default server interface
func NewBasicServer() (pb.AwsCostServer, error) {
	return &server{}, nil
}
