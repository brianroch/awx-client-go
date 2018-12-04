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

type WorkflowJobsResource struct {
	Resource
}

func NewWorkflowJobsResource(connection AwxConnection, path string) IdGetter {
	resource := new(WorkflowJobsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobsResource) Get() Sender {
	request := new(WorkflowJobsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobsResource) Id(id int) Getter {
	return NewWorkflowJobResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type WorkflowJobsGetRequest struct {
	Request
}

func (r *WorkflowJobsGetRequest) Filter(name string, value interface{}) Sender {
	r.addFilter(name, value)
	return r
}

func (r *WorkflowJobsGetRequest) Send() (response interface{}, err error) {
	response = new(WorkflowJobsGetResponse)
	err = r.get(response)
	if err != nil {
		return
	}
	return
}
