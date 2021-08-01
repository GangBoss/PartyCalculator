package patterns

import (
	"syscall"
)

type Repository interface {
	Find(syscall.GUID) interface{}
	Add(interface{}) interface{}
	Delete(syscall.GUID) bool
	Update(syscall.GUID, interface{}) interface{}
}
