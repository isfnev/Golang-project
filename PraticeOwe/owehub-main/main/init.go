package main

import (
	"PraticeOwe/shared/appserver"
	"PraticeOwe/shared/types"
	"strings"
	apiHandler "PraticeOwe/owehub-main/services"
)

var apiRoutes = appserver.ApiRoutes{
	{
		Method:             strings.ToUpper("post"),
		Pattern:            "/pratice-owe/v1/login",
		Handler:            apiHandler.HandleLoginRequest,
		IsAuthReq:          false,
		GroupAllowedAccess: []types.UserGroup{types.GroupAdmin},
	},
}
