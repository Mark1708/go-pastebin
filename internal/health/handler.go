package health

import "net/http"

// Read godoc
//
//	@summary		Check health
//	@description	Endpoint for checking health
//	@tags			Health API
//	@success		200
//	@router			/health [get]
func Read(w http.ResponseWriter, r *http.Request) {}
