package imageprocessor

type Processor struct {
	config Config
	storage Storage
}

func New (config Config , storage Storage) Processor {
	return Processor{
		config: config,
		storage: storage,
	}
}