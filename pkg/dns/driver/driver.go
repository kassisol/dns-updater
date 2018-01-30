package driver

type Storager interface {
	List(canonical string) (string, []string, error)
	Add(name, canonical string) error
}
