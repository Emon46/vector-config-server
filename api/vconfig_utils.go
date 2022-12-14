package api

import (
	"gopkg.in/yaml.v3"
)

type VectorConfig struct {
	DataDir    string                 `yaml:"data_dir"`
	Sinks      map[string]interface{} `yaml:"sinks"`
	Transforms map[string]interface{} `yaml:"transforms"`
	Sources    map[string]interface{} `yaml:"sources"`
}

func getUpdatedConfigWithRequestedConfig(currentConfig string, req updateVectorConfigRequest) (vectorDataYaml []byte, vectorConfig VectorConfig, err error) {
	err = yaml.Unmarshal([]byte(currentConfig), &vectorConfig)
	if err != nil {
		return nil, vectorConfig, err
	}

	// add new source configs
	for key, data := range req.Sources {
		vectorConfig.Sources[key] = data
	}
	// add new transforms configs
	for key, data := range req.Transforms {
		vectorConfig.Transforms[key] = data
	}
	// add new sinks configs
	for key, data := range req.Sinks {
		vectorConfig.Sinks[key] = data
	}

	vectorDataYaml, err = yaml.Marshal(&vectorConfig)
	if err != nil {
		return nil, vectorConfig, err
	}

	return vectorDataYaml, vectorConfig, nil
}
