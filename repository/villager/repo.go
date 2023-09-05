package villager

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
	"github.com/muhadif/mandes/core/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func NewVillagerRepository(db *gorm.DB) repository.VillagerRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetVillagerByQuery(ctx context.Context, request *entity.GetVillagerByQueryRequest) (*entity.GetVillagerByQueryResponse, error) {
	var villagers []*entity.Villager

	// Construct the SQL query dynamically based on the provided query parameters.
	query := "SELECT * FROM villager WHERE 1=1"
	args := []interface{}{}

	if request.VillageSerial != "" {
		query += " AND serial = ?"
		args = append(args, request.VillageSerial)
	}

	if request.VillagerName != "" {
		query += " AND full_name = ?"
		args = append(args, request.VillagerName)
	}

	// Handle pagination
	if request.Pagination != nil {
		var totalData int64
		r.db.Raw(query, args...).Count(&totalData)
		query += " LIMIT ? OFFSET ?"
		request.Pagination.Total = totalData
		request.Pagination.ValidatePagination()
		args = append(args, request.Pagination.PageSize, (request.Pagination.Page-1)*request.Pagination.PageSize)
	}

	rows, err := r.db.Raw(query, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var villager entity.Villager
		err := rows.Scan(
			&villager.ID,
			&villager.Serial,
			&villager.IdKTP,
			&villager.FullName,
			&villager.IdKK,
			&villager.Address,
			&villager.RT,
			&villager.RW,
			&villager.Status,
			&villager.CreatedAt,
			&villager.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		villagers = append(villagers, &villager)
	}

	return &entity.GetVillagerByQueryResponse{Villagers: villagers, Pagination: request.Pagination}, nil

}

func (r repo) CreateVillagerBatch(ctx context.Context, villagers []*entity.Villager) error {
	for _, villager := range villagers {
		// Use the Clauses method to specify the ON DUPLICATE KEY UPDATE clause.
		result := r.db.Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"full_name", "id_kk", "address", "RT", "RW", "status"}),
		}).Create(villager)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (r repo) UpdateVillager(ctx context.Context, req *entity.UpdateVillagerRequest) (*entity.UpdateVillagerResponse, error) {
	if err := r.db.Save(&req.Villager).Error; err != nil {
		return nil, err
	}

	return &entity.UpdateVillagerResponse{Serial: req.Villager.Serial}, nil
}

func (r repo) DeleteVillager(ctx context.Context, villagerSerial string) error {
	if err := r.db.Table("villager").Where("serial = ?", villagerSerial).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
