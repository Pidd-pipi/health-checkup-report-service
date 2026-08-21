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

func newMetricService(t *testing.T) (*AbnormalMetricService, *repository.AbnormalMetricRepository) {
	db := newTestDB(t)
	repo := repository.NewAbnormalMetricRepository(db)
	return NewAbnormalMetricService(repo, testLogger()), repo
}

func TestUpdateFollowUpToProcessing(t *testing.T) {
	svc, repo := newMetricService(t)
	m := model.AbnormalMetric{ExamineeID: 1, PackageItemID: 1, AbnormalLevel: "mild", FollowUpStatus: constants.FollowUpPending}
	if err := repo.Create(&m); err != nil {
		t.Fatal(err)
	}
	got, err := svc.UpdateFollowUp(context.Background(), m.ID, constants.FollowUpProcessing, "正在跟进")
	if err != nil {
		t.Fatalf("UpdateFollowUp(processing) err = %v", err)
	}
	if got.FollowUpStatus != constants.FollowUpProcessing {
		t.Fatalf("status = %s, want processing", got.FollowUpStatus)
	}
}

func TestUpdateFollowUpDoneIsTerminal(t *testing.T) {
	svc, repo := newMetricService(t)
	m := model.AbnormalMetric{ExamineeID: 1, PackageItemID: 1, AbnormalLevel: "mild", FollowUpStatus: constants.FollowUpDone}
	if err := repo.Create(&m); err != nil {
		t.Fatal(err)
	}
	_, err := svc.UpdateFollowUp(context.Background(), m.ID, constants.FollowUpPending, "")
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusConflict {
		t.Fatalf("err = %v, want Conflict(409)", err)
	}
}

func TestUpdateFollowUpKeepsAdvice(t *testing.T) {
	svc, repo := newMetricService(t)
	m := model.AbnormalMetric{ExamineeID: 1, PackageItemID: 1, AbnormalLevel: "mild", FollowUpStatus: constants.FollowUpPending}
	if err := repo.Create(&m); err != nil {
		t.Fatal(err)
	}
	got, err := svc.UpdateFollowUp(context.Background(), m.ID, constants.FollowUpDone, "建议复查")
	if err != nil {
		t.Fatal(err)
	}
	if got.SpecialistAdvice != "建议复查" {
		t.Fatalf("specialist_advice = %q, want 建议复查", got.SpecialistAdvice)
	}
}

func TestListIncludesProcessingMetric(t *testing.T) {
	_, repo := newMetricService(t)
	for _, st := range []string{constants.FollowUpPending, constants.FollowUpProcessing, constants.FollowUpDone} {
		m := model.AbnormalMetric{ExamineeID: 1, PackageItemID: 1, AbnormalLevel: "mild", FollowUpStatus: st}
		if err := repo.Create(&m); err != nil {
			t.Fatal(err)
		}
	}
	items, total, err := repo.List(1, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 3 || total != 3 {
		t.Fatalf("items=%d total=%d, want 3/3", len(items), total)
	}
}
