package migrations

import "context"

// MockRun mocks Run function
func MockRun(err error) Run {
	return func(ctx context.Context, migrationsDir string) error {
		return err
	}
}
