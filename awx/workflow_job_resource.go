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

type WorkflowJobResource struct {
	Resource
}

func NewWorkflowJobResource(connection *Connection, path string) Getter {
	resource := new(WorkflowJobResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobResource) Get() Sender {
	request := new(WorkflowJobGetRequest)
	request.resource = &r.Resource
	return request
}

type WorkflowJobGetRequest struct {
	Request
}

func (r *WorkflowJobGetRequest) Send() (response interface{}, err error) {
	response = new(WorkflowJobGetResponse)
	err = r.get(response)
	if err != nil {
		return nil, err
	}
	return
}

func (r *WorkflowJobGetRequest) Filter(name string, value interface{}) Sender {
	r.addFilter(name, value)
	return r
}
