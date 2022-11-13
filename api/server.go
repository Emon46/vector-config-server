package api

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/tools/clientcmd"
)

type Server struct {
	config     Config
	kubeConfig *restclient.Config
	kubeClient *kubernetes.Clientset
	router     *gin.Engine
}

// NewServer it will create a new gin api and setup routing for all the api call
func NewServer(config Config) (*Server, error) {
	kubeConfig, err := restclient.InClusterConfig()
	if err != nil {
		klog.Fatalln(err)
	}
	clientcmd.Fix(kubeConfig)
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		klog.Fatalln(err)
	}

	server := &Server{
		config:     config,
		kubeConfig: kubeConfig,
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
