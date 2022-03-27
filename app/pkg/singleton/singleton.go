package singleton

import (
  "sync"
)

type Singleton struct {
  instance interface{}
}

func (s *Singleton) GetInstance() interface{} {
  return s.instance
}

func (s *Singleton) SetInstance(once sync.Once, ins interface{}) {
  once.Do(func (){
    s.instance = ins
  })
}
