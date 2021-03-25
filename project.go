package ganboard

import (
	"encoding/json"
)

// CreateProject https://docs.kanboard.org/en/latest/api/project_procedures.html#createprojects
func (c *Client) CreateProject(params ProjectParams) (int, error) {
	query := request{
		Client: c,
		Method: "createProject",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// GetAllProjects https://docs.kanboard.org/en/latest/api/project_procedures.html#getallprojects
func (c *Client) GetAllProjects() ([]Project, error) {
	query := request{
		Client: c,
		Method: "getAllProjects",
	}
	response, err := query.decodeProjects()
	return response, err
}

// GetProjectByID https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyid
func (c *Client) GetProjectByID(id int) (Project, error) {
	query := request{
		Client: c,
		Method: "getProjectById",
		Params: ProjectParams{
			ProjectID: id,
		},
	}
	response, err := query.decodeProject()
	return response, err
}

// GetProjectByName https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByName(name string) (Project, error) {
	query := request{
		Client: c,
		Method: "getProjectByName",
		Params: ProjectParams{
			Name: name,
		},
	}
	response, err := query.decodeProject()
	return response, err
}

// GetProjectByIdentifier https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByIdentifier(identifier string) (Project, error) {
	query := request{
		Client: c,
		Method: "getProjectByIdentifier",
		Params: ProjectParams{
			Identifier: identifier,
		},
	}
	response, err := query.decodeProject()
	return response, err
}

// UpdateProject https://docs.kanboard.org/en/latest/api/project_procedures.html#updateproject
func (c *Client) UpdateProject(params ProjectParams) (bool, error) {
	query := request{
		Client: c,
		Method: "updateProject",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveProject https://docs.kanboard.org/en/latest/api/project_procedures.html#removeproject
func (c *Client) RemoveProject(projectID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeProject",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// EnableProject https://docs.kanboard.org/en/latest/api/project_procedures.html#enableproject
func (c *Client) EnableProject(projectID int) (bool, error) {
	query := request{
		Client: c,
		Method: "enableProject",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// DisableProject https://docs.kanboard.org/en/latest/api/project_procedures.html#enableproject
func (c *Client) DisableProject(projectID int) (bool, error) {
	query := request{
		Client: c,
		Method: "disableProject",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetProjectActivity https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectactivity
// FIXME describe Activity type
func (c *Client) GetProjectActivity(projectID int) (interface{}, error) {
	query := request{
		Client: c,
		Method: "getProjectActivity",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}
	response, err := query.decodeInterface()
	return response, err
}

// GetProjectActivities https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectactivity
// FIXME describe Activity type
func (c *Client) GetProjectActivities(projectID []int) (interface{}, error) {
	query := request{
		Client: c,
		Method: "getProjectActivity",
		Params: map[string][]int{
			"project_ids": projectID,
		},
	}
	response, err := query.decodeInterface()
	return response, err
}

// TODO: create Event type to decode Json in usable structs. Missing data type documentation.

// ProjectParams parameters for UpdateProject()
type ProjectParams struct {
	ProjectID   int    `json:"project_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
	Email       string `json:"email,omitempty"`
	OwnerID     int    `json:"owner_id,omitempty"`
	Description string `json:"decription,omitempty"`
}

// Project reflects getAllProjects method
type Project struct {
	ID                  json.Number `json:"id,string"`
	Name                string      `json:"name"`
	IsActive            json.Number `json:"is_active,string"`
	Token               string      `json:"token"`
	LastModified        json.Number `json:"last_modified,string"`
	IsPublic            json.Number `json:"is_public,string"`
	IsPrivate           json.Number `json:"is_private,string"`
	DefaultSwimlane     string      `json:"default_swimlane"`
	ShowDefaultSwimlane json.Number `json:"show_default_swimlane,string"`
	Description         string      `json:"description"`
	Identifier          string      `json:"identifier"`
	URL                 struct {
		Board    string `json:"board"`
		Calendar string `json:"calendar"`
		List     string `json:"list"`
	} `json:"url"`
}

func (r *request) decodeProjects() ([]Project, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string    `json:"jsonrpc"`
		ID      FlexInt   `json:"id"`
		Result  []Project `json:"result"`
	}{}
	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeProject() (Project, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Project{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  Project `json:"result"`
	}{}
	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
