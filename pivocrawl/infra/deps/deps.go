package deps

import (
	"github.com/melphi/pivocrawl/infra/api"
	"github.com/melphi/pivocrawl/infra/config"
)

type Deps struct {
	ApiService *api.Service
}

func Init() *Deps {
	cfg := config.New()
	apiService := NewApiService(cfg)
	return &Deps{
		ApiService: apiService,
	}
}

func NewApiService(cfg *config.Config) *api.Service {
	return api.NewService(cfg.ApiPort)
}
