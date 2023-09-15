package service

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
	Service struct {
		Contract     []string `yaml:"contract,omitempty"`
		Dependencies []string `yaml:"dependencies,omitempty"`
		Database     struct {
			DB       string `yaml:"db,omitempty"`
			Name     string `yaml:"name,omitempty"`
			Host     string `yaml:"host,omitempty"`
			Port     string `yaml:"port,omitempty"`
			Username string `yaml:"username,omitempty"`
			Password string `yaml:"password,omitempty"`
		} `yaml:"database,omitempty"`
	} `yaml:"service,omitempty"`
}

func ReadServiceConfigFile(service string) (*Service, error) {
	var serviceFile *os.File
	var err error

	if service != "" {
		path := fmt.Sprintf("%s/services/%s/service.yaml", strings.TrimSpace(os.Getenv("WORKSPACE")), strings.TrimSpace(service))
		serviceFile, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	}

	if service == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		serviceFile, err = os.Open(fmt.Sprintf("%s/service.yaml", cwd))
		if err != nil {
			return nil, err
		}
	}

	serviceFileB, err := io.ReadAll(serviceFile)
	if err != nil {
		return nil, err
	}

	var svc Service
	yaml.Unmarshal(serviceFileB, &svc)

	return &svc, nil
}
