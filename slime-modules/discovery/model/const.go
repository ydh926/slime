package model

import (
	"istio.io/istio-mcp/pkg/config/schema/resource"
	"slime.io/slime/slime-framework/util"
)

var GVKServiceEntry = resource.TypeUrlToGvk(util.TypeUrl_ServiceEntry)
var GVKSidecar = resource.TypeUrlToGvk(util.TypeUrl_Sidecar)