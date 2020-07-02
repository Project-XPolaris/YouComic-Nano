package datasource

type DataFilter interface {
	filter(data interface{}) (bool, error)
}
type DataReader interface {
	Read() (interface{}, int, error)
}

