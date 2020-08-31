package resNodeTree

func TreeToJson(t Tree) (string, error) {
	return GroupToJson(t.RootGroup)
}

func TreeToJsonIndent(t Tree) (string, error) {
	return GroupToJsonIndent(t.RootGroup)
}
