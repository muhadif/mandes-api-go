package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/mandes/core/module"
)

type VillagerHandler interface {
	GetVillagerByQuery(ctx *gin.Context)
	CreateVillagerBatch(ctx *gin.Context)
	UpdateVillager(ctx *gin.Context)
	DeleteVillager(ctx *gin.Context)
}

func NewVillagerHandler(authModule module.AuthModule) VillagerHandler {
	return &villagerHandler{}
}

type villagerHandler struct {
}

func (v *villagerHandler) GetVillagerByQuery(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerHandler) CreateVillagerBatch(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerHandler) UpdateVillager(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *villagerHandler) DeleteVillager(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
