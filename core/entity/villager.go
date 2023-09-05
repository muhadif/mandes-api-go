package entity

import "time"

type GetVillagerByQueryRequest struct {
	Pagination    *Pagination
	VillageSerial string
	VillagerName  string
}

type Villager struct {
	ID        string
	Serial    string
	IdKTP     string
	IdKK      string
	FullName  string
	Address   string
	RT        string
	RW        string
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type GetVillagerByQueryResponse struct {
	Villagers  []*Villager
	Pagination *Pagination
}

type CreateVillagerBatchRequest struct {
	Villagers []*Villager
}

type CreateVillagerBatchResponse struct {
}

type UpdateVillagerRequest struct {
	Villager Villager
}

type UpdateVillagerResponse struct {
	Serial string
}

type DeleteVillagerRequest struct {
	serial string
}

type DeleteVillagerResponse struct {
}
