package object

func (object *Object) HaveChildren(childrenName string) bool {

	for index := range object.Children {

		if object.Children[index].Name == childrenName {
			return true
		}
	}

	return false
}
