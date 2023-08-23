// SPDX-License-Identifier: Apache-2.0
//
// The OpenSearch Contributors require contributions made to
// this file be licensed under the Apache-2.0 license or a
// compatible open source license.
//
// Modifications Copyright OpenSearch Contributors. See
// GitHub history for details.

// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package opensearchapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func newRoleMappingDeleteFunc(t Transport) RoleMappingDelete {
	return func(role string, o ...func(*RoleMappingDeleteRequest)) (*Response, error) {
		var r = RoleMappingDeleteRequest{Role: role}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// RoleMappingDelete creates an role with optional settings and mappings.
type RoleMappingDelete func(role string, o ...func(*RoleMappingDeleteRequest)) (*Response, error)

// RoleMappingDeleteRequest configures the RoleMapping Create API request.
type RoleMappingDeleteRequest struct {
	Role string

	Body io.Reader

	MasterTimeout         time.Duration
	ClusterManagerTimeout time.Duration
	Timeout               time.Duration
	WaitForActiveShards   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r RoleMappingDeleteRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "DELETE"

	path.Grow(37 + len(r.Role))
	path.WriteString("/_plugins/_security/api/rolesmapping/")
	path.WriteString(r.Role)

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f RoleMappingDelete) WithContext(v context.Context) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.ctx = v
	}
}

// WithBody - The configuration for the role (`settings` and `mappings`).
func (f RoleMappingDelete) WithBody(v io.Reader) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.Body = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to cluster-manager node.
//
// Deprecated: To promote inclusive language, use WithClusterManagerTimeout instead.
func (f RoleMappingDelete) WithMasterTimeout(v time.Duration) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.MasterTimeout = v
	}
}

// WithClusterManagerTimeout - explicit operation timeout for connection to cluster-manager node.
func (f RoleMappingDelete) WithClusterManagerTimeout(v time.Duration) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.ClusterManagerTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f RoleMappingDelete) WithTimeout(v time.Duration) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.Timeout = v
	}
}

// WithWaitForActiveShards - set the number of active shards to wait for before the operation returns..
func (f RoleMappingDelete) WithWaitForActiveShards(v string) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f RoleMappingDelete) WithPretty() func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f RoleMappingDelete) WithHuman() func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f RoleMappingDelete) WithErrorTrace() func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f RoleMappingDelete) WithFilterPath(v ...string) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f RoleMappingDelete) WithHeader(h map[string]string) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f RoleMappingDelete) WithOpaqueID(s string) func(*RoleMappingDeleteRequest) {
	return func(r *RoleMappingDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
