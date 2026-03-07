package rules

type Rule interface {
	ID() string
	Validate(params map[string]any) error
}
