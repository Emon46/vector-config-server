package main

import (
	"github.com/Emon46/vector-config-server/api"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := api.LoadConfig("/")
	if err != nil {
		log.Fatal(err)
	}
	kubeConfig, err := restclient.InClusterConfig()
	if err != nil {
		klog.Fatalln(err)
	}
	clientcmd.Fix(kubeConfig)
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		klog.Fatalln(err)
	}
	server, err := api.NewServer(config, kubeClient)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
