package metrick

import (
	"github.com/segmentio/fasthash/fnv1a"
)

func (self *Tree) Label(label string) *Label {

	h := hashableUint64(fnv1a.HashString64(label))

	self.Lock()
	defer self.Unlock()

	obj := self.Labels[label]
	if obj == nil {
		obj = &Label{
			tree: self,
			Label: label,
		}
		obj.hashableUint64 = h
		self.Labels[label] = obj
	}
	return obj
}

type Label struct {
	tree *Tree
	hashableUint64
	Label string
}
