package db

import (
	"github.com/zerjioang/gotools/common"
	"github.com/zerjioang/gotools/lib/stack"
)

type DaoInterface interface {
	Key() []byte
	Value(serializer common.Serializer) []byte
	// creates new instance
	// to allow concurrent access, etc
	NewDao() DaoInterface
	Decode(data []byte) (DaoInterface, stack.Error)
	Update(o DaoInterface) (DaoInterface, stack.Error)
}
