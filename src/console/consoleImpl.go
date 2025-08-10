package console

import "github.com/labstack/gommon/log"

type Modules struct {
}

type consoleImpl struct {
	modules Modules
}

func Provide() *consoleImpl {
	// TODO: implement modules initialization
	return &consoleImpl{
		modules: Modules{},
	}
}

func (s *consoleImpl) Start() {
	log.Info("started")
}
