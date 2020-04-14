package metrick

import (
	"hash"
)

type Node interface {
	GetIndex() int
	Parent() Node
	Add(hash.Hash64)
	Contains(hash.Hash64) bool
}
