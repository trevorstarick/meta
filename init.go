package meta

func init() {
	tree = make([]node, 0xff)

	buildLUT()
}
