package resGNode

// Group Node all-in-one struct for Vue front-end
type ResGNode struct {
	ID       int64       `json:"id"`
	Label    string      `json:"label"`
	Children []*ResGNode `json:"children,omitempty"`
}
