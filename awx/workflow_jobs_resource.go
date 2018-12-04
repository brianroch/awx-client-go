/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awx

import (
	"fmt"
)

type WorkFlowJobsResource struct {
	Resource
}

func NewWorkFlowJobsResource(connection AwxConnection, path string) IdGetter {
	resource := new(WorkFlowJobsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkFlowJobsResource) Get() Sender {
	request := new(WorkFlowJobsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkFlowJobsResource) Id(id int) Getter {
	return NewWorkFlowJobResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type WorkFlowJobsGetRequest struct {
	Request
}

func (r *WorkFlowJobsGetRequest) Filter(name string, value interface{}) Sender {
	r.addFilter(name, value)
	return r
}

func (r *WorkFlowJobsGetRequest) Send() (response interface{}, err error) {
	response = new(WorkFlowJobsGetResponse)
	err = r.get(response)
	if err != nil {
		return
	}
	return
}
