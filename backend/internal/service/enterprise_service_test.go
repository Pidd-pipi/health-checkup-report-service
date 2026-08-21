package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/blueship581/gbcheckup/internal/constants"
	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/repository"
	"github.com/blueship581/gbcheckup/internal/util"
)

func TestDeliverReportsMissingOrder404(t *testing.T) {
	db := newTestDB(t)
	svc := NewEnterpriseService(repository.NewEnterpriseRepository(db), repository.NewGroupOrderRepository(db), repository.NewPackageRepository(db), testLogger())
	_, err := svc.DeliverReports(context.Background(), 999)
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusNotFound {
		t.Fatalf("err = %v, want NotFound(404)", err)
	}
}

func TestDeliverReportsAlreadyDone409(t *testing.T) {
	db := newTestDB(t)
	ent := model.Enterprise{Name: "华信科技"}
	if err := repository.NewEnterpriseRepository(db).Create(&ent); err != nil {
		t.Fatal(err)
	}
	pkg := model.Package{Name: "入职体检", PackageType: "entry", Price: 300, Status: "active"}
	if err := repository.NewPackageRepository(db).Create(&pkg); err != nil {
		t.Fatal(err)
	}
	order := model.GroupOrder{EnterpriseID: ent.ID, PackageID: pkg.ID, ExamineeCount: 10, Status: constants.GroupOrderDone}
	if err := repository.NewGroupOrderRepository(db).Create(&order); err != nil {
		t.Fatal(err)
	}
	svc := NewEnterpriseService(repository.NewEnterpriseRepository(db), repository.NewGroupOrderRepository(db), repository.NewPackageRepository(db), testLogger())
	_, err := svc.DeliverReports(context.Background(), order.ID)
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusConflict {
		t.Fatalf("err = %v, want Conflict(409)", err)
	}
}

func TestCreateOrderMissingPackage404(t *testing.T) {
	db := newTestDB(t)
	ent := model.Enterprise{Name: "华信科技"}
	if err := repository.NewEnterpriseRepository(db).Create(&ent); err != nil {
		t.Fatal(err)
	}
	svc := NewEnterpriseService(repository.NewEnterpriseRepository(db), repository.NewGroupOrderRepository(db), repository.NewPackageRepository(db), testLogger())
	_, err := svc.CreateOrder(context.Background(), ent.ID, 999, 10)
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusNotFound {
		t.Fatalf("err = %v, want NotFound(404)", err)
	}
}
