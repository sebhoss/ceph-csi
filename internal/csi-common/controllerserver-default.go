/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csicommon

import (
	"context"

	"github.com/ceph/ceph-csi/internal/util/log"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DefaultControllerServer points to default driver.
type DefaultControllerServer struct {
	csi.UnimplementedControllerServer
	Driver *CSIDriver
}

// ControllerGetCapabilities implements the default GRPC callout.
// Default supports all capabilities.
func (cs *DefaultControllerServer) ControllerGetCapabilities(
	ctx context.Context,
	req *csi.ControllerGetCapabilitiesRequest,
) (*csi.ControllerGetCapabilitiesResponse, error) {
	log.TraceLog(ctx, "Using default ControllerGetCapabilities")
	if cs.Driver == nil {
		return nil, status.Error(codes.Unimplemented, "Controller server is not enabled")
	}

	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: cs.Driver.capabilities,
	}, nil
}
