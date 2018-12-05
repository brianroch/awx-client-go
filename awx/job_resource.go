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

// This file contains the implementation of the resource that manages launching of jobs from job
// templates.

package awx

type JobResource struct {
	Resource
}

func NewJobResource(connection *Connection, path string) Getter {
	resource := new(JobResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *JobResource) Get() Sender {
	request := new(JobGetRequest)
	request.resource = &r.Resource
	return request
}

type JobGetRequest struct {
	Request
}

func (r *JobGetRequest) Send() (response interface{}, err error) {
	response = new(JobGetResponse)
	err = r.get(response)
	if err != nil {
		return nil, err
	}
	return
}

func (r *JobGetRequest) Filter(name string, value interface{}) Sender {
	r.addFilter(name, value)
	return r
}
