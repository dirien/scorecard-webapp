// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetResultHandlerFunc turns a function with the right signature into a get result handler
type GetResultHandlerFunc func(GetResultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResultHandlerFunc) Handle(params GetResultParams) middleware.Responder {
	return fn(params)
}

// GetResultHandler interface for that can handle valid get result params
type GetResultHandler interface {
	Handle(GetResultParams) middleware.Responder
}

// NewGetResult creates a new http.Handler for the get result operation
func NewGetResult(ctx *middleware.Context, handler GetResultHandler) *GetResult {
	return &GetResult{Context: ctx, Handler: handler}
}

/* GetResult swagger:route GET /projects/{platform}/{org}/{repo} results getResult

Get a repository's ScorecardResult

*/
type GetResult struct {
	Context *middleware.Context
	Handler GetResultHandler
}

func (o *GetResult) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetResultParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
