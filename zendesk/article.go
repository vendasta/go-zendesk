package zendesk

import (
	"fmt"
	"time"
)

// Article represents a Zendesk Article
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/#json-format
type Article struct {
	ID                *int64     `json:"id,omitempty"`
	AuthorID          *int64     `json:"author_id,omitempty"`
	Body              *string    `json:"body,omitempty"`
	CommentsDisabled  *bool      `json:"comments_disabled,omitempty"`
	Draft             *bool      `json:"draft,omitempty"`
	HtmlURL           *string    `json:"html_url,omitempty"`
	LabelNames        []string   `json:"label_names,omitempty"`
	Locale            *string    `json:"locale,omitempty"`
	Outdated          *bool      `json:"outdated,omitempty"`
	OutdatedLocales   []string   `json:"outdated_locales,omitempty"`
	PermissionGroupID *int64     `json:"permission_group_id,omitempty"`
	Position          *int64     `json:"position,omitempty"`
	Promoted          *bool      `json:"promoted,omitempty"`
	SectionID         *int64     `json:"section_id,omitempty"`
	SourceLocale      *string    `json:"source_locale,omitempty"`
	Title             *string    `json:"title,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	URL               *string    `json:"url,omitempty"`
	UserSegmentID     *int64     `json:"user_segment_id"`
	VoteCount         *int64     `json:"vote_count,omitempty"`
	VoteSum           *int64     `json:"vote_sum,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	EditedAt          *time.Time `json:"edited_at,omitempty"`
}

// ShowArticle shows a Zendesk help article for a given article id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/#show-article
func (c *client) ShowArticle(id int64) (*Article, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/help_center/articles/%d.json", id), out)
	return out.Article, err
}

// CreateArticle will create a new Zendesk Help Center article
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/#create-article
func (c *client) CreateArticle(sectionID int64, article *Article) (*Article, error) {
	in := &APIPayload{Article: article}
	out := new(APIPayload)
	err := c.post(fmt.Sprintf("/api/v2/help_center/sections/%d/articles.json", sectionID), in, out)
	return out.Article, err
}

// UpdateArticle will update a Zendesk Help Center article
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/#update-article
func (c *client) UpdateArticle(id int64, article *Article) (*Article, error) {
	in := &APIPayload{Article: article}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/help_center/articles/%d.json", id), in, out)
	return out.Article, err
}

// DeleteArticle archives a Zendesk help article for a given article id
//
// Zendesk Help Center API docs: https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/#archive-article
func (c *client) DeleteArticle(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/help_center/articles/%d.json", id), nil)
}
