package modules

type Scanner struct {
	ip         string
	ports      string
	scanModule string
	execParam  string
	check      bool
	c          map[string]string
	filePath   string
	Results    string
	targets    []string
}
