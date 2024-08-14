package core

import (
	"sync"

	cd "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"

	engine "github.com/muidea/magicEngine/http"

	_ "github.com/muidea/quickModbus/internal/core/kernel/master"
	_ "github.com/muidea/quickModbus/internal/core/kernel/slave"
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

	httpServer        engine.HTTPServer
	eventHub          event.Hub
	backgroundRoutine task.BackgroundRoutine
}

// Startup 启动
func (s *Core) Startup(eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) (err *cd.Result) {
	routeRegistry := engine.NewRouteRegistry()

	s.httpServer = engine.NewHTTPServer(s.listenPort)
	s.httpServer.Bind(routeRegistry)

	modules := module.GetModules()
	for _, val := range modules {
		module.BindRegistry(val, routeRegistry)
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
