package repository

import (
	"auth-microservice/src/helpers"
	"auth-microservice/src/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var user = &models.User{
	gorm.Model{ID: 12313},
	"test@mail.ru",
	 "66bf6557ba05ce87c5a0b641edf0615ce8e29a6c9e840b529d6b1e39a53d81d0",
	 helpers.ACTIVE,
}

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return gdb, mock
}

func TestGetUserByEmail(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer db.Close()

	query := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = $1)) ORDER BY "users"."id" ASC LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(user.ID, user.Email))
	userResult,err := repo.GetUserByEmail(user.Email)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NotNil(t, userResult)
	assert.NoError(t, err)
	assert.Equal(t, user.ID,userResult.ID)
}

func TestGetUserById(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer db.Close()

	id := strconv.FormatUint(uint64(user.ID), 10)
	query := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND (("users"."id" = ` + id  +`)) ORDER BY "users"."id" ASC LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(user.ID, user.Email))
	userResult,err := repo.GetUserById(user.ID)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NotNil(t, userResult)
	assert.NoError(t, err)
	assert.Equal(t, user.Email,userResult.Email)
}

func TestCreateUser(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer db.Close()

	mock.ExpectBegin()
	id := strconv.FormatUint(uint64(user.ID), 10)
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","email","password","status") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "users"."id"`)).
		WithArgs(time.Now(),time.Now(),time.Now(),user.Email, user.Password,user.Status).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))
	mock.ExpectCommit()

	usr,err := repo.CreateUser(user.Email,user.Password)
	assert.NotNil(t, usr)
	assert.NoError(t, err)
}