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

func newRoleDeleteFunc(t Transport) RoleDelete {
	return func(role string, o ...func(*RoleDeleteRequest)) (*Response, error) {
		var r = RoleDeleteRequest{Role: role}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// RoleDelete creates an role with optional settings and mappings.
type RoleDelete func(role string, o ...func(*RoleDeleteRequest)) (*Response, error)

// RoleDeleteRequest configures the Role Create API request.
type RoleDeleteRequest struct {
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
func (r RoleDeleteRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "DELETE"

	path.Grow(30 + len(r.Role))
	path.WriteString("/_plugins/_security/api/roles/")
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
func (f RoleDelete) WithContext(v context.Context) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.ctx = v
	}
}

// WithBody - The configuration for the role (`settings` and `mappings`).
func (f RoleDelete) WithBody(v io.Reader) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.Body = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to cluster-manager node.
//
// Deprecated: To promote inclusive language, use WithClusterManagerTimeout instead.
func (f RoleDelete) WithMasterTimeout(v time.Duration) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.MasterTimeout = v
	}
}

// WithClusterManagerTimeout - explicit operation timeout for connection to cluster-manager node.
func (f RoleDelete) WithClusterManagerTimeout(v time.Duration) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.ClusterManagerTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f RoleDelete) WithTimeout(v time.Duration) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.Timeout = v
	}
}

// WithWaitForActiveShards - set the number of active shards to wait for before the operation returns..
func (f RoleDelete) WithWaitForActiveShards(v string) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f RoleDelete) WithPretty() func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f RoleDelete) WithHuman() func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f RoleDelete) WithErrorTrace() func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f RoleDelete) WithFilterPath(v ...string) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f RoleDelete) WithHeader(h map[string]string) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f RoleDelete) WithOpaqueID(s string) func(*RoleDeleteRequest) {
	return func(r *RoleDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
