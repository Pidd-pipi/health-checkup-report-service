package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newRegRepoTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.Registration{}, &model.Examinee{}, &model.Package{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestRegistrationRepoListCtxCancel(t *testing.T) {
	db := newRegRepoTestDB(t)
	repo := NewRegistrationRepository(db)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err := repo.List(ctx, "", 1, 10)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("List err = %v, want context.Canceled", err)
	}
}
