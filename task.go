package ganboard

import (
	"encoding/json"
	"time"
)

// CreateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#createtask
func (c *Client) CreateTask(params TaskParams) (int, error) {
	query := request{
		Client: c,
		Method: "createTask",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// GetTask https://docs.kanboard.org/en/latest/api/task_procedures.html#gettask
func (c *Client) GetTask(taskID int) (Task, error) {
	query := request{
		Client: c,
		Method: "getTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeTask()
	return response, err
}

// GetTaskByReference https://docs.kanboard.org/en/latest/api/task_procedures.html#gettaskbyreference
func (c *Client) GetTaskByReference(projectID int, reference string) (Task, error) {
	query := request{
		Client: c,
		Method: "getTaskByReference",
		Params: struct {
			ProjectID int    `json:"project_id"`
			Reference string `json:"reference"`
		}{
			ProjectID: projectID,
			Reference: reference,
		},
	}
	response, err := query.decodeTask()
	return response, err
}

// GetAllTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getalltasks
func (c *Client) GetAllTasks(projectID int, statusID int) ([]Task, error) {
	query := request{
		Client: c,
		Method: "getAllTasks",
		Params: map[string]int{
			"project_id": projectID,
			"status_id":  statusID,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// GetOverdueTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasks
func (c *Client) GetOverdueTasks() ([]Task, error) {
	query := request{
		Client: c,
		Method: "getOverdueTasks",
	}
	response, err := query.decodeTasks()
	return response, err
}

// GetOverdueTasksByProject https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasksbyproject
func (c *Client) GetOverdueTasksByProject(projectID int) ([]Task, error) {
	query := request{
		Client: c,
		Method: "getOverdueTasksByProject",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// UpdateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#updatetask
func (c *Client) UpdateTask(params TaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "updateTask",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// OpenTask https://docs.kanboard.org/en/latest/api/task_procedures.html#opentask
func (c *Client) OpenTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "openTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// CloseTask https://docs.kanboard.org/en/latest/api/task_procedures.html#closetask
func (c *Client) CloseTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "closeTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveTask https://docs.kanboard.org/en/latest/api/task_procedures.html#removetask
func (c *Client) RemoveTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// MoveTaskPosition https://docs.kanboard.org/en/latest/api/task_procedures.html#movetaskposition
func (c *Client) MoveTaskPosition(projectID int, taskID int, columnID int, position int, swimlaneID int) (bool, error) {
	query := request{
		Client: c,
		Method: "moveTaskPosition",
		Params: map[string]int{
			"project_id":  projectID,
			"task_id":     taskID,
			"column_id":   columnID,
			"position":    position,
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// MoveTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#movetasktoproject
func (c *Client) MoveTaskToProject(params MoveTaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "moveTaskToProject",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// DuplicateTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#duplicatetasktoproject
func (c *Client) DuplicateTaskToProject(params MoveTaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "duplicateTaskToProject",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// SearchTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#searchtasks
func (c *Client) SearchTasks(projectID int, queryString string) ([]Task, error) {
	query := request{
		Client: c,
		Method: "searchTasks",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Query     string `json:"query"`
		}{
			ProjectID: projectID,
			Query:     queryString,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// TaskParams input for CreateTask
type TaskParams struct {
	ID                  int        `json:"id,omitempty"`
	Title               string     `json:"title"`
	ProjectID           int        `json:"project_id"`
	ColorID             string     `json:"color_id,omitempty"`
	ColumnID            int        `json:"column_id,string,omitempty"`
	OwnerID             int        `json:"owner_id,string,omitempty"`
	CreatorID           int        `json:"creator_id,omitempty"`
	DateDue             *time.Time `json:"date_due,omitempty"`
	Description         string     `json:"description,omitempty"`
	CategoryID          int        `json:"category_id,string,omitempty"`
	Score               int        `json:"score,string,omitempty"`
	SwimlaneID          int        `json:"swimlane_id,string,omitempty"`
	Priority            int        `json:"priority,omitempty"`
	RecurrenceStatus    int        `json:"recurrence_status,string,omitempty"`
	RecurrenceTrigger   int        `json:"recurrence_trigger,string,omitempty"`
	RecurrenceFactor    int        `json:"recurrence_factor,string,omitempty"`
	RecurrenceTimeframe int        `json:"recurrence_timeframe,string,omitempty"`
	RecurrenceBaseDate  int        `json:"recurrence_basedate,string,omitempty"`
	Reference           string     `json:"reference"`
	Tags                []string   `json:"tags,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
}

// MoveTaskParams input for MoveTaskToProject
type MoveTaskParams struct {
	TaskID     int `json:"task_id,string"`
	ProjectID  int `json:"project_id,string"`
	SwimlaneID int `json:"swimlane_id,string,omitempty"`
	ColumnID   int `json:"column_id,string,omitempty"`
	CategoryID int `json:"category_id,string,omitempty"`
	OwnerID    int `json:"owner_id,string,omitempty"`
}

// Task type
type Task struct {
	ID                  json.Number `json:"id,string,omitempty"`
	ProjectID           json.Number `json:"project_id,string"`
	ColorID             string      `json:"color_id,omitempty"`
	ColumnID            json.Number `json:"column_id,string,omitempty"`
	OwnerID             json.Number `json:"owner_id,string,omitempty"`
	CreatorID           json.Number `json:"creator_id,omitempty"`
	DateDue             json.Number `json:"date_due,string,omitempty"`
	Reference           string      `json:"reference"`
	Title               string      `json:"title"`
	Description         string      `json:"description"`
	DateCreation        json.Number `json:"date_creation,string"`
	DateCompleted       json.Number `json:"date_completed,string"`
	DateModification    json.Number `json:"date_modification,string"`
	DateStarted         json.Number `json:"date_started,string"`
	TimeEstimated       json.Number `json:"time_estimated,string"`
	TimeSpent           json.Number `json:"time_spend,string"`
	Position            json.Number `json:"position,string"`
	IsActive            json.Number `json:"is_active,string"`
	Score               json.Number `json:"score,string"`
	CategoryID          json.Number `json:"category_id,string"`
	SwimlaneID          json.Number `json:"swimlane_id,string,omitempty"`
	DateMoved           json.Number `json:"date_moved"`
	RecurrenceStatus    json.Number `json:"recurrence_status,string"`
	RecurrenceTrigger   json.Number `json:"recurrence_trigger,string"`
	RecurrenceFactor    json.Number `json:"recurrence_factor,string"`
	RecurrenceTimeframe json.Number `json:"recurrence_timeframe,string"`
	RecurrenceBaseDate  json.Number `json:"recurrence_basedate,string"`
	RecurrenceParent    json.Number `json:"recurrence_parent,string"`
	RecurrenceChild     json.Number `json:"recurrence_child,string"`
	CategoryName        string      `json:"category_name"`
	ProjectName         string      `json:"project_name"`
	DefaultSwimlane     string      `json:"default_swimlane"`
	ColumnTitle         string      `json:"column_title"`
	AssigneeUsername    string      `json:"assignee_username"`
	AssigneeName        string      `json:"assignee_name"`
	CreatorUsername     string      `json:"creator_username"`
	CreatorName         string      `json:"creator_name"`
	NbComments          json.Number `json:"nb_comments,string"`
	NbFiles             json.Number `json:"nb_files,string"`
	NbSubtasks          json.Number `json:"nb_subtasks,string"`
	NbCompletedSubtasks json.Number `json:"nb_completed_subtasks,string"`
	NbLinks             json.Number `json:"nb_links,string"`
	Color               Color       `json:"color"`
}

func (r *request) decodeTasks() ([]Task, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  []Task  `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeTask() (Task, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Task{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  Task    `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
