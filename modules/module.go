package modules

type Scanner struct {
	scanModule int
	execParam  string
	check      bool
	c          map[string]string
	results    string
	outPutPath string
}
type params struct {
	Type int    `json:"type"`
	Args string `json:"args"`
}
