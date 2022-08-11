package zendesk

import (
	"fmt"
	"time"
)

// Category represents a Zendesk Category
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/categories/#json-format
type Category struct {
	ID           *int64     `json:"id,omitempty"`
	HtmlURL      *string    `json:"html_url,omitempty"`
	Locale       *string    `json:"locale,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Outdated     *bool      `json:"outdated,omitempty"`
	Position     *int64     `json:"position,omitempty"`
	SourceLocale *string    `json:"sourceLocale,omitempty"`
	URL          *string    `json:"URL,omitempty"`
	Description  *string    `json:"description,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}

// ShowCategory shows a Zendesk help category for a given category id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/categories/#show-category
func (c *client) ShowCategory(id int64) (*Category, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/help_center/categories/%d.json", id), out)
	return out.Category, err
}

// CreateCategory will create a new Zendesk Help Center category
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/categories/#create-category
func (c *client) CreateCategory(category *Category) (*Category, error) {
	in := &APIPayload{Category: category}
	out := new(APIPayload)
	err := c.post("/api/v2/help_center/categories.json", in, out)
	return out.Category, err
}

// UpdateCategory will update a Zendesk Help Center category
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/categories/#update-category
func (c *client) UpdateCategory(id int64, category *Category) (*Category, error) {
	in := &APIPayload{Category: category}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/help_center/categories/%d.json", id), in, out)
	return out.Category, err
}

// DeleteCategory deletes a Zendesk help category for a given category id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/categories/#delete-category
func (c *client) DeleteCategory(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/help_center/categories/%d.json", id), nil)
}
