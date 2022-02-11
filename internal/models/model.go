package models

type Member struct {
	ID          int    `json:"id,omitempty"`
	email       string `json:"email,omitempty"`
	displayName string `json:"display_name,omitempty"`
	description string `json:"description,omitempty"`
	lvl         int    `json:"lvl,omitempty"`

	// TODO: continue the rest of table model
}
