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

// PostResultHandlerFunc turns a function with the right signature into a post result handler
type PostResultHandlerFunc func(PostResultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostResultHandlerFunc) Handle(params PostResultParams) middleware.Responder {
	return fn(params)
}

// PostResultHandler interface for that can handle valid post result params
type PostResultHandler interface {
	Handle(PostResultParams) middleware.Responder
}

// NewPostResult creates a new http.Handler for the post result operation
func NewPostResult(ctx *middleware.Context, handler PostResultHandler) *PostResult {
	return &PostResult{Context: ctx, Handler: handler}
}

/* PostResult swagger:route POST /projects/{platform}/{org}/{repo} results postResult

Publish a repository's OIDC verified ScorecardResult

*/
type PostResult struct {
	Context *middleware.Context
	Handler PostResultHandler
}

func (o *PostResult) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostResultParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}