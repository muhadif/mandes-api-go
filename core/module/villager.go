package module

import (
	"context"
	"github.com/muhadif/mandes/config"
	"github.com/muhadif/mandes/core/entity"
)

type VillagerModule interface {
	GetVillagerByQuery(ctx context.Context, req *entity.GetVillagerByQueryRequest) (*entity.GetVillagerByQueryResponse, error)
	CreateVillagerBatch(ctx context.Context, req *entity.CreateVillagerBatchRequest) (*entity.CreateVillagerBatchResponse, error)
	UpdateVillager(ctx context.Context, req *entity.UpdateVillagerRequest) (*entity.UpdateVillagerResponse, error)
	DeleteVillager(ctx context.Context, req *entity.DeleteVillagerRequest) (*entity.DeleteVillagerResponse, error)
}

type villagerModule struct {
	cfg config.Config
}

func NewVillagerModule(cfg config.Config) VillagerModule {
	return &villagerModule{
		cfg: cfg,
	}
}

func (v *villagerModule) GetVillagerByQuery(ctx context.Context, req *entity.GetVillagerByQueryRequest) (*entity.GetVillagerByQueryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerModule) CreateVillagerBatch(ctx context.Context, req *entity.CreateVillagerBatchRequest) (*entity.CreateVillagerBatchResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerModule) UpdateVillager(ctx context.Context, req *entity.UpdateVillagerRequest) (*entity.UpdateVillagerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerModule) DeleteVillager(ctx context.Context, req *entity.DeleteVillagerRequest) (*entity.DeleteVillagerResponse, error) {
	//TODO implement me
	panic("implement me")
}
