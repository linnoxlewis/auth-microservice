package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetDB(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	_, err = gorm.Open("postgres", db)

	assert.NoError(t, err)
}
