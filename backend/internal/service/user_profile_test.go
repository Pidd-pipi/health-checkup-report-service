package service

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbcheckup/internal/repository"
	"github.com/blueship581/gbcheckup/internal/util"
)

func TestUpdateProfileMissingUser404(t *testing.T) {
	db := newTestDB(t)
	svc := NewUserService(repository.NewUserRepository(db), "secret", 24, testLogger())
	_, err := svc.UpdateProfile(context.Background(), 999, "新名字", "", "")
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != 404 {
		t.Fatalf("err = %v, want NotFound(404)", err)
	}
}

func TestLoginMissingPhoneReturnsUnauthorized(t *testing.T) {
	db := newTestDB(t)
	svc := NewUserService(repository.NewUserRepository(db), "secret", 24, testLogger())
	_, _, err := svc.Login(context.Background(), "13700000001", "pass123")
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != 401 {
		t.Fatalf("err = %v, want Unauthorized(401)", err)
	}
}
