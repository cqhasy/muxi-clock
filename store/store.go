package store

type Type string

const (
	MapStore Type = "MapStore"
)

type Entity struct {
	Key  string
	Data interface{}
}

type Store interface {
	Type() Type
	Connect(addr string, args ...interface{}) (interface{}, error)
	InitTables(name ...string) error
	Get(key string) (Entity, error)
	Update(key string, val Entity) (Entity, error)
	Create(val Entity) error
	Delete(key string) error
	Custom(fn func() (interface{}, error)) (interface{}, error)
}
