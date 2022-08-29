package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pffreitas/optimizely-terraform-provider/optimizely/attribute"
)

type Post struct {
	Archived    bool   `json:"archived"`
	Description string `json:"description"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	ProjectId   int    `json:"project_id"`
}

func (c OptimizelyClient) DeleteAttribute(id int) error {
	url := fmt.Sprintf("v2/attributes/%d", id)

	if _, err := c.sendHttpRequest("DELETE", url, nil); err != nil {
		return err
	}

	return nil
}

func (c OptimizelyClient) CreateAttribute(archived bool, description, key, name string, project_id int) (string, error) {
	instance := new(Post)
	instance.Archived = archived
	instance.Description = description
	instance.Key = key
	instance.Name = name
	instance.ProjectId = project_id

	body, err := json.Marshal(instance)

	if err != nil {
		return "", err
	}

	res, err := c.sendHttpRequest("POST", "v2/attributes", bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	var result attribute.Attribute
	json.Unmarshal(res, &result)

	return strconv.FormatInt(result.ID, 10), nil
}

func (c OptimizelyClient) GetAttribute(id int) (*attribute.Attribute, error) {
	url := fmt.Sprintf("v2/attributes/%d", id)
	res, err := c.sendHttpRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	var instance attribute.Attribute
	json.Unmarshal(res, &instance)

	return &instance, nil
}

func (c OptimizelyClient) UpdateAttribute(id int, archived bool, description, key, name string) error {
	instance := new(attribute.Attribute)
	instance.Archived = archived
	instance.Description = description
	instance.Key = key
	instance.Name = name

	body, err := json.Marshal(instance)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("v2/attributes/%d", id)
	if _, err := c.sendHttpRequest("PATCH", url, bytes.NewBuffer(body)); err != nil {
		return err
	} else {
		return nil
	}
}
