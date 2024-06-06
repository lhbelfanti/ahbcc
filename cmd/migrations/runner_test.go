package migrations_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/migrations"
)

const migrationsTestDir string = "./migrations_test"

func TestMain(m *testing.M) {
	setupTestEnvironment()
	code := m.Run()
	teardownTestEnvironment()
	os.Exit(code)
}

func setupTestEnvironment() {
	os.Mkdir(migrationsTestDir, 0766)
}

func teardownTestEnvironment() {
	os.RemoveAll(migrationsTestDir)
}

func TestRun_success(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer os.Remove(migrationFile)
	mockConn := migrations.MockPgxConnStruct(migrations.MockExecFunc("EXECUTE", nil))

	run := migrations.MakeRun(mockConn)

	got := run(context.Background(), migrationsTestDir)

	assert.NoError(t, got)
}

func TestRun_failsWhenExecuteSQLFromFileThrowsUnableToReadFileError(t *testing.T) {
	invalidFile := filepath.Join(migrationsTestDir, "invalid.sql")
	err := os.WriteFile(invalidFile, []byte("CREATE TABLE test (id INT);"), 0000)
	assert.NoError(t, err)
	defer os.Remove(invalidFile)
	mockConn := migrations.MockPgxConnStruct(migrations.MockExecFunc("EXECUTE", nil))

	run := migrations.MakeRun(mockConn)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
}

func TestRun_failsWhenExecuteSQLFromFileThrowsUnableToExecuteSQLError(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer os.Remove(migrationFile)
	mockConn := migrations.MockPgxConnStruct(migrations.MockExecFunc("", migrations.UnableToExecuteSQL))

	run := migrations.MakeRun(mockConn)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
}
