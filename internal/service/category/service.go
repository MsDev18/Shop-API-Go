package category

import "shop/internal/pkg/imageprocessor"

type Service struct {
	repository     Repository
	imageProcessor imageprocessor.Processor
}

type Repository interface {
}

func New(repository Repository, imageprocessor imageprocessor.Processor) Service {
	return Service{
		repository: repository,
		imageProcessor: imageprocessor,
	}
}
