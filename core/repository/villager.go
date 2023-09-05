package repository

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
)

type VillagerRepository interface {
	GetVillagerByQuery(ctx context.Context, req *entity.GetVillagerByQueryRequest) (*entity.GetVillagerByQueryResponse, error)
	CreateVillagerBatch(ctx context.Context, req []*entity.Villager) error
	UpdateVillager(ctx context.Context, req *entity.UpdateVillagerRequest) (*entity.UpdateVillagerResponse, error)
	DeleteVillager(ctx context.Context, villagerSerial string) error
}
