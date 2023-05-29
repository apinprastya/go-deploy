package action

type IAction interface {
	Name() string
	Setup(jsonConfig []byte) error
	Start() error
	Stop() error
	Shutdown() error
	Result() (int, error)
	IsRunning() bool
}
