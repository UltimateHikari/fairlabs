package adapter

import (
	"fairlabs-server/api"
)

var LookupGoalsKind = api.Kind{Action: api.Get, Endpoint: "/goals"}
var ProgressKind = api.Kind{Action: api.Get, Endpoint: "/progress"}
