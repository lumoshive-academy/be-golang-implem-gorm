package vendor

import (
	"go-33/internal/dto"
	"time"
)

func VonderA(phone string) dto.ResponValidatePhone {
	time.Sleep(2 * time.Second)
	data := dto.ResponValidatePhone{
		Status: "valid",
		Phone:  phone,
		Vendor: "vendor A",
	}
	return data
}

func VonderB(phone string) dto.ResponValidatePhone {
	time.Sleep(1 * time.Second)
	data := dto.ResponValidatePhone{
		Status: "valid",
		Phone:  phone,
		Vendor: "vendor B",
	}
	return data
}

func VonderC(phone string) dto.ResponValidatePhone {
	time.Sleep(4 * time.Second)
	data := dto.ResponValidatePhone{
		Status: "valid",
		Phone:  phone,
		Vendor: "vendor C",
	}
	return data
}
