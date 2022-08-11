package zendesk

import (
	"fmt"
	"time"
)

// Section represents a Zendesk Section
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/sections/#json-format
type Section struct {
	ID              *int64     `json:"id,omitempty"`
	CategoryID      *int64     `json:"category_id,omitempty"`
	Description     *string    `json:"description,omitempty"`
	HtmlURL         *string    `json:"html_url,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Outdated        *bool      `json:"outdated,omitempty"`
	ParentSectionID *string    `json:"parentSectionID,omitempty"`
	Position        *int64     `json:"position,omitempty"`
	SourceLocale    *string    `json:"sourceLocale,omitempty"`
	ThemeTemplate   *string    `json:"themeTemplate,omitempty"`
	URL             *string    `json:"URL,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
}

// ShowSection shows a Zendesk help section for a given section id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/sections/#show-section
func (c *client) ShowSection(id int64) (*Section, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/help_center/sections/%d.json", id), out)
	return out.Section, err
}

// CreateSection will create a new Zendesk Help Center section
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/sections/#create-section
func (c *client) CreateSection(categoryID int64, section *Section) (*Section, error) {
	in := &APIPayload{Section: section}
	out := new(APIPayload)
	err := c.post(fmt.Sprintf("/api/v2/help_center/categories/%d/sections.json", categoryID), in, out)
	return out.Section, err
}

// UpdateSection will update a Zendesk Help Center section
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/sections/#update-section
func (c *client) UpdateSection(id int64, section *Section) (*Section, error) {
	in := &APIPayload{Section: section}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/help_center/sections/%d.json", id), in, out)
	return out.Section, err
}

// DeleteSection deletes a Zendesk help section for a given section id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/sections/#delete-section
func (c *client) DeleteSection(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/help_center/sections/%d.json", id), nil)
}
