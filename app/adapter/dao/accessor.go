package dao

type Accessor interface {
  Create(interface{}) (*Accessor)
  Save(interface{}) (*Accessor)
  First(interface{}, ...interface{}) (*Accessor)
  LastError() error
}
