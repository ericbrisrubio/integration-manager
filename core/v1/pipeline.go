package v1


type Pipeline struct {
	ApiVersion string            `json:"api_version" yaml:"api_version"`
	Name       string            `json:"name"  yaml:"name"`
	ProcessId  string            `json:"process_id" yaml:"process_id"`
	Label      map[string]string `json:"label" yaml:"label"`
	Steps      [] Step           `json:"steps" yaml:"steps"`
}

func(pipeline Pipeline)Validate()error{

	return nil
}