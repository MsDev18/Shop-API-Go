package imageprocessor

type Processor struct {
	config Config
}

func New(config Config) Processor {
	return Processor{
		config: config,
	}
}
