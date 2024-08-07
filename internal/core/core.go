package core

import (
	"sync"

	cd "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/session"
	"github.com/muidea/magicCommon/task"

	engine "github.com/muidea/magicEngine/http"
)

// New 新建Core
func New(endpointName, listenPort string) (ret *Core, err *cd.Result) {
	core := &Core{
		endpointName: endpointName,
		listenPort:   listenPort,
	}

	ret = core
	return
}

// Core Core对象
type Core struct {
	endpointName string
	listenPort   string

	sessionRegistry   session.Registry
	httpServer        engine.HTTPServer
	eventHub          event.Hub
	backgroundRoutine task.BackgroundRoutine
}

// Startup 启动
func (s *Core) Startup(eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) (err *cd.Result) {
	router := engine.NewRouter()
	sessionRegistry := session.CreateRegistry()
	s.sessionRegistry = sessionRegistry

	s.httpServer = engine.NewHTTPServer(s.listenPort)
	s.httpServer.Bind(router)

	modules := module.GetModules()
	for _, val := range modules {
		module.Setup(val, s.endpointName, eventHub, backgroundRoutine)
	}

	s.eventHub = eventHub
	s.backgroundRoutine = backgroundRoutine
	return
}

func (s *Core) Run() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		modules := module.GetModules()
		for _, val := range modules {
			module.Run(val)
		}
	}()

	wg.Wait()
	s.httpServer.Run()
}

// Shutdown 销毁
func (s *Core) Shutdown() {
	modules := module.GetModules()
	totalSize := len(modules)
	for idx := range modules {
		module.Teardown(modules[totalSize-idx-1])
	}
}
