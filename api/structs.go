package api

// Configuration for a source
type Source struct {
	URL        string   `json:"url"`
	Checksum   string   `json:"checksum"`
	Type       string   `json:"type"`
	Commit     string   `json:"commit"`
	Tag        string   `json:"tag"`
	Branch     string   `json:"branch"`
	Packages   []string `json:"packages"`
	Path       string   `json:"path"`
	OnlyArches []string `json:"only-arches" mapstructure:"only-arches"`
}

// Configuration for a recipe
type Recipe struct {
	Name          string
	Id            string
	Vibversion    string
	Stages        []Stage
	Path          string
	ParentPath    string
	DownloadsPath string
	SourcesPath   string
	IncludesPath  string
	PluginPath    string
	Containerfile string
	Finalize      []interface{}
}

// Configuration for a stage in the recipe
type Stage struct {
	Id          string            `json:"id"`
	Base        string            `json:"base"`
	Copy        []Copy            `json:"copy"`
	Addincludes bool              `json:"addincludes"`
	Labels      map[string]string `json:"labels"`
	Env         map[string]string `json:"env"`
	Adds        []Add             `json:"adds"`
	Args        map[string]string `json:"args"`
	Runs        Run               `json:"runs"`
	Expose      map[string]string `json:"expose"`
	Cmd         Cmd               `json:"cmd"`
	Modules     []interface{}     `json:"modules"`
	Entrypoint  Entrypoint
}

type PluginType int

const (
	BuildPlugin PluginType = iota
	FinalizePlugin
)

// Information about a plugin
type PluginInfo struct {
	Name             string
	Type             PluginType
	UseContainerCmds bool
}

// Configuration for copying files or directories in a stage
type Copy struct {
	From    string
	SrcDst  map[string]string
	Workdir string
}

// Configuration for adding files or directories in a stage
type Add struct {
	SrcDst  map[string]string
	Workdir string
}

// Configuration for the entrypoint of a container
type Entrypoint struct {
	Exec    []string
	Workdir string
}

// Configuration for a command to run in the container
type Cmd struct {
	Exec    []string
	Workdir string
}

// Configuration for commands to run in the container
type Run struct {
	Commands []string
	Workdir  string
}
