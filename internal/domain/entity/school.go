package entity

type School struct {
	ID    string
	Name  string
	Style string
	Enemy string
}

func NewSchool(id, name, style, enemy string) *School {
	return &School{
		ID:    id,
		Name:  name,
		Style: style,
		Enemy: enemy,
	}
}