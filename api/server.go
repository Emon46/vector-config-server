package api

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

type Server struct {
	config     Config
	kubeClient kubernetes.Interface
	router     *gin.Engine
}

// NewServer it will create a new gin api and setup routing for all the api call
func NewServer(config Config, kubeClient kubernetes.Interface) (*Server, error) {

	server := &Server{
		config:     config,
		kubeClient: kubeClient,
	}

	server.setUpRouterWithSubUrl()
	return server, nil
}

func (server *Server) setUpRouterWithSubUrl() {
	router := gin.Default()
	router.GET("/config", server.GetVectorConfig)
	router.POST("/config", server.AddNewTransformInConfig)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
