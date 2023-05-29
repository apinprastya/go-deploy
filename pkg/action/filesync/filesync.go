package filesync

type FileSync struct {
}

func New() *FileSync {
	return &FileSync{}
}

func (fs *FileSync) Name() string {
	return "filesync"
}

func (fs *FileSync) Setup(jsonConfig []byte) error {
	return nil
}

func (fs *FileSync) Start() error {
	return nil
}

func (fs *FileSync) Stop() error {
	return nil
}

func (fs *FileSync) Shutdown() error {
	return nil
}

func (fs *FileSync) Result() (int, error) {
	return 0, nil
}

func (fs *FileSync) IsRunning() bool {
	return false
}
