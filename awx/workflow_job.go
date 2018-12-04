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

// This file contains the data structures used for sending and receiving jobs.

package awx

type WorkflowJob struct {
	Id          int     `json:"id,omitempty"`
	Status      string  `json:"status,omitempty"`
	Name        string  `json:"name,omitempty"`
	Elapsed     float32 `json:"elapsed,omitempty"`
	Finished    string  `json:"finished,omitempty"`
	Started     string  `json:"started,omitempty"`
	Failed      bool    `json:"failed,omitempty"`
	Description string  `json:"description,omitempty"`
}

type WorkflowJobGetResponse struct {
	WorkflowJob
}
