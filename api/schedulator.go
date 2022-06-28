package api

type Rule struct {
	ID         int64  `db:"id" json:"id,omitempty"`
	Name       string `db:"name" json:"name,omitempty"`
	Expression string `db:"expression" json:"expression,omitempty"`
	Type       string `db:"rule_type" json:"type,omitempty"`
	Data       string `db:"data" json:"data,omitempty"`
}

type CreateRule struct {
	Name       string `db:"name" json:"name,omitempty"`
	Expression string `db:"expression" json:"expression,omitempty"`
	Type       string `db:"rule_type" json:"type,omitempty"`
	Data       string `db:"data" json:"data,omitempty"`
}

type UpdateRule struct {
	Name       string `db:"name" json:"name,omitempty"`
	Expression string `db:"expression" json:"expression,omitempty"`
	Type       string `db:"rule_type" json:"type,omitempty"`
	Data       string `db:"data" json:"data,omitempty"`
}
