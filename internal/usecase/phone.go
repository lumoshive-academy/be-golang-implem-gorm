package usecase

import (
	"go-33/internal/dto"
	"go-33/pkg/vendor"
)

type UsecasePhone struct {
}

func NewUsecasePhone() UsecasePhone {
	return UsecasePhone{}
}

func (p *UsecasePhone) Validate(phone string) *dto.ResponValidatePhone {
	ch1 := make(chan dto.ResponValidatePhone)
	ch2 := make(chan dto.ResponValidatePhone)
	ch3 := make(chan dto.ResponValidatePhone)

	go vendor.VonderA(phone)
	go vendor.VonderB(phone)
	go vendor.VonderC(phone)

	select {
	case data := <-ch1:
		return &data
	case data := <-ch2:
		return &data
	case data := <-ch3:
		return &data
	default:
		return nil
	}
}
