package repository

import (
	"echo-hello/domain/mock"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFindUser_shouldFound(t *testing.T) {
	t.Parallel()
	sqlDB, db, mock := mock.MockDB(t)
	defer sqlDB.Close()

	ur := NewUserRepository(db)
	user := sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "foo@mail.com", "password", "bar", time.Now(), time.Now(), nil)

	expectedSQL := `SELECT (.+) FROM "users" WHERE "users"."id" = (.+) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(user)
	_, err := ur.FindByID(1)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUser_shouldNotFound(t *testing.T) {
	t.Parallel()
	sqlDB, db, mock := mock.MockDB(t)
	defer sqlDB.Close()

	ur := NewUserRepository(db)
	user := sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at", "deleted_at"})

	expectedSQL := `SELECT (.+) FROM "users" WHERE "users"."id" = (.+) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(user)
	_, err := ur.FindByID(1)

	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUserByEmail_shouldFound(t *testing.T) {
	t.Parallel()
	sqlDB, db, mock := mock.MockDB(t)
	defer sqlDB.Close()

	ur := NewUserRepository(db)
	user := sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "foo@mail.com", "password", "bar", time.Now(), time.Now(), nil)

	expectedSQL := `SELECT (.+) FROM "users" WHERE email = (.+) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	mock.ExpectQuery(expectedSQL).WithArgs("foo@mail.com").WillReturnRows(user)
	_, err := ur.FindByEmail("foo@mail.com")

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUserByEmail_shouldNotFound(t *testing.T) {
	t.Parallel()
	sqlDB, db, mock := mock.MockDB(t)
	defer sqlDB.Close()

	ur := NewUserRepository(db)
	user := sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at", "deleted_at"})

	expectedSQL := `SELECT (.+) FROM "users" WHERE email = (.+) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	mock.ExpectQuery(expectedSQL).WithArgs("foo@mail.com").WillReturnRows(user)
	_, err := ur.FindByEmail("foo@mail.com")

	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUserAll(t *testing.T) {
	t.Parallel()
	sqlDB, db, mock := mock.MockDB(t)
	defer sqlDB.Close()

	ur := NewUserRepository(db)
	users := sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "foo@mail.com", "password", "bar", time.Now(), time.Now(), nil).
		AddRow(2, "haaa@mail.com", "password", "max", time.Now(), time.Now(), nil)

	expectedSQL := `SELECT (.+) FROM "users" WHERE "users"."deleted_at" IS NULL`

	mock.ExpectQuery(expectedSQL).WillReturnRows(users)
	res, err := ur.FindAll()

	assert.Equal(t, 2, len(res))
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
