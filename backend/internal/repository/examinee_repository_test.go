package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newRepoTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.Examinee{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestExamineeRepoCreateBatchCancelledContext(t *testing.T) {
	db := newRepoTestDB(t)
	repo := NewExamineeRepository(db)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := repo.CreateBatch(ctx, []model.Examinee{{Name: "张三", IDCardNo: "110101199001011234"}})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("CreateBatch err = %v, want context.Canceled", err)
	}
}
