package content

type Profile struct {
	ID       string                 `yaml:"id"`
	Name     string                 `yaml:"name"`
	Extends  []string               `yaml:"extends"`
	Vars     map[string]any         `yaml:"vars"`
	Rules    []RuleInvocation       `yaml:"rules"`
}

type RuleInvocation struct {
	ID     string                 `yaml:"id"`
	Params map[string]any         `yaml:"params"`
}
