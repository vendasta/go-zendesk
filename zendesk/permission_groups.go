package zendesk

import (
	"fmt"
	"time"
)

// PermissionGroup represents a Zendesk Permission Group
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/permission_groups/#json-format
type PermissionGroup struct {
	ID        *int64     `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Edit      []int64    `json:"edit,omitempty"`
	Publish   []int64    `json:"publish,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	BuiltIn   *bool      `json:"built_in,omitempty"`
}

// ShowPermissionGroup shows a Zendesk help permission group for a given permission group id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/permission_groups/#show-permission-group
func (c *client) ShowPermissionGroup(id int64) (*PermissionGroup, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/guide/permission_groups/%d.json", id), out)
	return out.PermissionGroup, err
}

// CreatePermissionGroup will create a new Zendesk Help Center permission group
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/guide/help-center-api/permission_groups/#create-permission-group
func (c *client) CreatePermissionGroup(permissionGroup *PermissionGroup) (*PermissionGroup, error) {
	in := &APIPayload{PermissionGroup: permissionGroup}
	out := new(APIPayload)
	err := c.post("/api/v2/guide/permission_groups.json", in, out)
	return out.PermissionGroup, err
}

// UpdatePermissionGroup will update a Zendesk Help Center permission group
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/permission_groups/#update-permission-group
func (c *client) UpdatePermissionGroup(id int64, permissionGroup *PermissionGroup) (*PermissionGroup, error) {
	in := &APIPayload{PermissionGroup: permissionGroup}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/guide/permission_groups/%d.json", id), in, out)
	return out.PermissionGroup, err
}

// DeletePermissionGroup deletes a Zendesk help permission group for a given permissionGroup id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/permission_groups/#delete-permission-group
func (c *client) DeletePermissionGroup(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/guide/permission_groups/%d.json", id), nil)
}
