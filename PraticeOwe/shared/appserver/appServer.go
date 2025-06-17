package appserver

import (
	"PraticeOwe/shared/types"
	"net/http"
)

type ServiceApiRoute struct {
	Method             string
	Pattern            string
	Handler            http.HandlerFunc
	IsAuthReq          bool
	GroupAllowedAccess []types.UserGroup
}

type ApiRoutes []ServiceApiRoute