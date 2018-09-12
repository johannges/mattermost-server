// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	GroupTypeLdap             = "ldap"
	GroupNameMaxLength        = 64
	GroupTypeMaxLength        = 64
	GroupDisplayNameMaxLength = 128
	GroupDescriptionMaxLength = 1024
)

var groupTypes = []string{
	GroupTypeLdap,
}

type Group struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	DisplayName string     `json:"display_name"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Props       GroupProps `json:"props"`
	CreateAt    int64      `json:"create_at"`
	UpdateAt    int64      `json:"update_at"`
	DeleteAt    int64      `json:"delete_at"`
}

type GroupProps map[string]interface{}

func (groupProps GroupProps) Value() (driver.Value, error) {
	j, err := json.Marshal(groupProps)
	return j, err
}

func (groupProps *GroupProps) Scan(src interface{}) error {
	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("Incompatible type for GroupProps")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	var ok bool
	*groupProps, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

func (group *Group) IsValidForCreate() *AppError {
	if l := len(group.Name); l == 0 || l > GroupNameMaxLength {
		return NewAppError("Group.IsValidForCreate", "model.group.name.app_error", map[string]interface{}{"GroupNameMaxLength": GroupNameMaxLength}, "", http.StatusBadRequest)
	}

	if l := len(group.DisplayName); l == 0 || l > GroupDisplayNameMaxLength {
		return NewAppError("Group.IsValidForCreate", "model.group.display_name.app_error", map[string]interface{}{"GroupDisplayNameMaxLength": GroupDisplayNameMaxLength}, "", http.StatusBadRequest)
	}

	if len(group.Description) > GroupDescriptionMaxLength {
		return NewAppError("Group.IsValidForCreate", "model.group.description.app_error", map[string]interface{}{"GroupDescriptionMaxLength": GroupDescriptionMaxLength}, "", http.StatusBadRequest)
	}

	isValidType := false
	for _, groupType := range groupTypes {
		if group.Type == groupType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		return NewAppError("Group.IsValidForCreate", "model.group.type.app_error", map[string]interface{}{"ValidGroupTypes": strings.Join(groupTypes, ", ")}, "", http.StatusBadRequest)
	}

	if group.Props == nil {
		return NewAppError("Group.IsValidForCreate", "model.group.props.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (group *Group) IsValidForUpdate() *AppError {
	if len(group.Id) != 26 {
		return NewAppError("Group.IsValidForUpdate", "model.group.id.app_error", nil, "", http.StatusBadRequest)
	}
	if group.CreateAt == 0 {
		return NewAppError("Group.IsValidForCreate", "model.group.create_at.app_error", nil, "", http.StatusBadRequest)
	}
	if group.UpdateAt == 0 {
		return NewAppError("Group.IsValidForCreate", "model.group.update_at.app_error", nil, "", http.StatusBadRequest)
	}
	if err := group.IsValidForCreate(); err != nil {
		return err
	}
	return nil
}
