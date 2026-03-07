package rules

type Registry struct {
	rules map[string]Rule
}

func NewRegistry() *Registry {
	return &Registry{rules: make(map[string]Rule)}
}

func (r *Registry) Register(rule Rule) {
	r.rules[rule.ID()] = rule
}
