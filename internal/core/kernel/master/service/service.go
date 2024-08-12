package service

import (
	"context"
	"net/http"

	engine "github.com/muidea/magicEngine/http"
)

type Master struct {
	routeRegistry engine.Router
}

func (s *Master) BindRegistry(router engine.Router) {
	s.routeRegistry = router
}

func (s *Master) ConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) DisConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}
