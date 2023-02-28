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

func newRoleMappingCreateFunc(t Transport) RoleMappingCreate {
	return func(role string, o ...func(*RoleMappingCreateRequest)) (*Response, error) {
		var r = RoleMappingCreateRequest{Role: role}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// RoleMappingCreate creates an role with optional settings and mappings.
type RoleMappingCreate func(role string, o ...func(*RoleMappingCreateRequest)) (*Response, error)

// RoleMappingCreateRequest configures the RoleMapping Create API request.
type RoleMappingCreateRequest struct {
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
func (r RoleMappingCreateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(37 + len(r.Role))
	path.WriteString("/_plugins/_security/api/rolesmapping/")
	path.WriteString(r.Role)

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.ClusterManagerTimeout != 0 {
		params["cluster_manager_timeout"] = formatDuration(r.ClusterManagerTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.WaitForActiveShards != "" {
		params["wait_for_active_shards"] = r.WaitForActiveShards
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
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
func (f RoleMappingCreate) WithContext(v context.Context) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.ctx = v
	}
}

// WithBody - The configuration for the role (`settings` and `mappings`).
func (f RoleMappingCreate) WithBody(v io.Reader) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.Body = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to cluster-manager node.
//
// Deprecated: To promote inclusive language, use WithClusterManagerTimeout instead.
func (f RoleMappingCreate) WithMasterTimeout(v time.Duration) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.MasterTimeout = v
	}
}

// WithClusterManagerTimeout - explicit operation timeout for connection to cluster-manager node.
func (f RoleMappingCreate) WithClusterManagerTimeout(v time.Duration) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.ClusterManagerTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f RoleMappingCreate) WithTimeout(v time.Duration) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.Timeout = v
	}
}

// WithWaitForActiveShards - set the number of active shards to wait for before the operation returns..
func (f RoleMappingCreate) WithWaitForActiveShards(v string) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f RoleMappingCreate) WithPretty() func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f RoleMappingCreate) WithHuman() func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f RoleMappingCreate) WithErrorTrace() func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f RoleMappingCreate) WithFilterPath(v ...string) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f RoleMappingCreate) WithHeader(h map[string]string) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f RoleMappingCreate) WithOpaqueID(s string) func(*RoleMappingCreateRequest) {
	return func(r *RoleMappingCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
