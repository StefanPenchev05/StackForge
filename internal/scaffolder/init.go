package scaffolder

type InitOptions struct {
	ProjectName  string
	WithFrontend bool
	WithAuth     bool
	Lang         string // "node" or "go"
	WithDocker   bool
	WithCI       bool
}

func InitProject(opts InitOptions) error {
	return nil
}
