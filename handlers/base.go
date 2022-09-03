package handlers

import (
	"github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	motor2 "github.com/sonr-io/sonr/third_party/types/motor"
	"walkwithme-backend/search"
)

type Server struct {
	Mtr      motor.MotorNode
	Requests []search.Request
}

func NewServer() (*Server, error) {
	s := &Server{}
	var err error

	s.Mtr, err = motor.EmptyMotor(&motor2.InitializeRequest{
		DeviceId: "unique_device_id",
	}, common.DefaultCallback())
	if err != nil {
		return nil, err
	}

	return s, nil
}
