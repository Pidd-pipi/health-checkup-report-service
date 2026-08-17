package repository

import (
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

// CountExamResults 统计引用该检查项目的检查结果数量。
func (r *PackageItemRepository) CountExamResults(itemID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.ExamResult{}).Where("package_item_id = ?", itemID).Count(&count).Error
	return count, err
}

// CountAbnormalMetrics 统计引用该检查项目的异常指标数量。
func (r *PackageItemRepository) CountAbnormalMetrics(itemID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.AbnormalMetric{}).Where("package_item_id = ?", itemID).Count(&count).Error
	return count, err
}

// CountGroupByDepartment 按科室统计工作量。
func (r *PackageItemRepository) CountGroupByDepartment() ([]model.NameCount, error) {
	var rows []model.NameCount
	err := r.db.Model(&model.PackageItem{}).Select("department as name, count(*) as count").Group("department").Scan(&rows).Error
	return rows, err
}
