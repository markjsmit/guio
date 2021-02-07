package element

type Element struct {
	Tag      string
	Attr     map[string]string
	Children []*Element
	Parent   *Element
}
