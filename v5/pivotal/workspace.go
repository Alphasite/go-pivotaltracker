// Copyright (c) 2018 Salsita Software
// Use of this source code is governed by the MIT License.
// The license can be found in the LICENSE file.

package pivotal

import (
	"fmt"
	"net/http"
)

// Workspace is the primary data object for the epic service.
type Workspace struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PersonID   int    `json:"person_id,omitempty"`
	ProjectIDs []int  `json:"project_ids,omitempty"`
	Kind       string `json:"kind,omitempty"`
}

// WorkspaceService wraps the client context to do actions.
type WorkspaceService struct {
	client *Client
}

func newWorkspaceService(client *Client) *WorkspaceService {
	return &WorkspaceService{client}
}

// List returns all workspaces availible to the user
func (service *WorkspaceService) List() ([]*Workspace, *http.Response, error) {
	req, err := service.client.NewRequest("GET", "my/workspaces", nil)
	if err != nil {
		return nil, nil, err
	}

	var workspaces []*Workspace
	resp, err := service.client.Do(req, &workspaces)
	if err != nil {
		return nil, resp, err
	}

	return workspaces, resp, err
}

func newWorkspaceRequestFunc(client *Client) func() *http.Request {
	return func() *http.Request {
		u := fmt.Sprintf("my/workspaces")
		req, _ := client.NewRequest("GET", u, nil)
		return req
	}
}
