package repository

import (
	"errors"

	"github.com/blueship581/gbcheckup/internal/model"
	"gorm.io/gorm"
)

// PackageItemRepository 检查项目仓储。
type PackageItemRepository struct{ db *gorm.DB }

// NewPackageItemRepository 构造检查项目仓储。
func NewPackageItemRepository(db *gorm.DB) *PackageItemRepository { return &PackageItemRepository{db: db} }

func (r *PackageItemRepository) Create(item *model.PackageItem) error { return r.db.Create(item).Error }

func (r *PackageItemRepository) CreateBatch(items []model.PackageItem) error {
	if len(items) == 0 {
		return nil
	}
	return r.db.Create(&items).Error
}

func (r *PackageItemRepository) ListByPackage(packageID uint) ([]model.PackageItem, error) {
	var items []model.PackageItem
	err := r.db.Where("package_id = ?", packageID).Order("sort_order asc").Find(&items).Error
	return items, err
}

func (r *PackageItemRepository) FindByID(id uint) (*model.PackageItem, error) {
	var item model.PackageItem
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PackageItemRepository) Update(item *model.PackageItem) error { return r.db.Save(item).Error }

func (r *PackageItemRepository) Delete(id uint) error { return r.db.Delete(&model.PackageItem{}, id).Error }

// CountGroupByDepartment 按科室统计工作量。
func (r *PackageItemRepository) CountGroupByDepartment() ([]model.NameCount, error) {
	var rows []model.NameCount
	err := r.db.Model(&model.PackageItem{}).Select("department as name, count(*) as count").Group("department").Scan(&rows).Error
	return rows, err
}

// BulkUpdate 批量更新检查项目。
func (r *PackageItemRepository) BulkUpdate(items []model.PackageItem) (err error) {
	tx := r.db.Begin()
	defer func() {
		err = tx.Commit().Error
	}()
	for _, it := range items {
		if it.ItemName == "" {
			return errors.New("package item name is empty")
		}
		if err = tx.Save(&it).Error; err != nil {
			return err
		}
	}
	return nil
}

// BulkDelete 批量删除检查项目。
func (r *PackageItemRepository) BulkDelete(ids []uint) (err error) {
	tx := r.db.Begin()
	defer func() {
		err = tx.Commit().Error
	}()
	for _, id := range ids {
		if id == 0 {
			return errors.New("package item id is empty")
		}
		if err = tx.Delete(&model.PackageItem{}, id).Error; err != nil {
			return err
		}
	}
	return nil
}
