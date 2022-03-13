package lib

import (
	"sync"
)

type Singleton struct {
	instance interface{}
}

func (s *Singleton) GetInstance() interface{} {
	return s.instance
}

// Thread safe
func (s *Singleton) SetInstance(once sync.Once, ins interface{}) {
	once.Do(func (){
		s.instance = ins
	})
}
