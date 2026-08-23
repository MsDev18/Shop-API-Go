package product

import "shop/internal/pkg/imageprocessor"

type Service struct {
	repository Repository
	imageProcessor imageprocessor.Processor
}

type Repository interface {

}

func New(repository Repository, imageProcessor imageprocessor.Processor) Service {
	return Service{
		repository: repository,
		imageProcessor: imageProcessor,
	}
}