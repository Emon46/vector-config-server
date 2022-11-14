package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core_util "kmodules.xyz/client-go/core/v1"
	"net/http"
)

type getVectorConfigRequest struct {
	ConfigMapName      string `json:"configMapName" binding:"required"`
	ConfigMapNameSpace string `json:"configMapNameSpace" binding:"required"`
}

func (server *Server) GetVectorConfig(ctx *gin.Context) {
	var req getVectorConfigRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	configMap, err := server.kubeClient.CoreV1().ConfigMaps(req.ConfigMapNameSpace).Get(context.TODO(), req.ConfigMapName, meta_v1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	vectorConfig := VectorConfig{}
	err = yaml.Unmarshal([]byte(configMap.Data[vectorConfigFileName]), &vectorConfig)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vectorConfig)

}

type updateVectorConfigRequest struct {
	ConfigMapName      string                 `json:"configMapName" binding:"required"`
	ConfigMapNameSpace string                 `json:"configMapNameSpace" binding:"required"`
	Transforms         map[string]interface{} `json:"transforms"`
	Sinks              map[string]interface{} `json:"sinks"`
	Sources            map[string]interface{} `json:"sources"`
}

func (server *Server) AddNewTransformInConfig(ctx *gin.Context) {
	var req updateVectorConfigRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get the current vector config by getting the config map
	configMap, err := server.kubeClient.CoreV1().ConfigMaps(req.ConfigMapNameSpace).Get(context.TODO(), req.ConfigMapName, meta_v1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	vectorConfigDataYaml, vectorConfig, err := UpdateVectorConfigWithRequestedConfig(configMap.Data[vectorConfigFileName], req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// now patch the configmap with new updated vector config
	_, _, err = core_util.PatchConfigMap(context.TODO(), server.kubeClient, configMap, func(in *core_v1.ConfigMap) *core_v1.ConfigMap {
		in.Data[vectorConfigFileName] = string(vectorConfigDataYaml)
		return in
	}, meta_v1.PatchOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vectorConfig)
}
