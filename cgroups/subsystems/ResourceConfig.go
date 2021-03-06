package subsystems

type ResourceConfig struct {
	MemoryLimit string
}

type Subsystem interface {
	Name() string
	Set(res *ResourceConfig) error
	Apply(pid string) error
	Remove() error
}

const (
	ResourceName = "mydocker"
)