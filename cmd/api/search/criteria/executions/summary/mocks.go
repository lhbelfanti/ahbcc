package summary

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// MockSelectMonthlyTweetsCountsByYearByCriteriaID mocks SelectMonthlyTweetsCountsByYearByCriteriaID function
func MockSelectMonthlyTweetsCountsByYearByCriteriaID(daos []DAO, err error) SelectMonthlyTweetsCountsByYearByCriteriaID {
	return func(ctx context.Context, criteriaID int) ([]DAO, error) {
		return daos, err
	}
}

// MockSelectAll mocks SelectAll function
func MockSelectAll(daos []DAO, err error) SelectAll {
	return func(ctx context.Context) ([]DAO, error) {
		return daos, err
	}
}

// MockDeleteAll mocks DeleteAll function
func MockDeleteAll(err error) DeleteAll {
	return func(ctx context.Context) error {
		return err
	}
}

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(tx pgx.Tx, ctx context.Context, dao DAO) (int, error) {
		return 1, err
	}
}

// MockExecutionSummaryDAO mocks a summary.DAO
func MockExecutionSummaryDAO(searchCriteriaID, year, month, total int) DAO {
	return DAO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Total:            total,
	}
}

// MockExecutionsSummaryDAOSlice mocks a []summary.DAO
func MockExecutionsSummaryDAOSlice() []DAO {
	return []DAO{
		MockExecutionSummaryDAO(2, 2025, 5, 105),
		MockExecutionSummaryDAO(2, 2025, 2, 333),
		MockExecutionSummaryDAO(1, 2025, 2, 500),
		MockExecutionSummaryDAO(1, 2024, 9, 350),
		MockExecutionSummaryDAO(1, 2025, 1, 1000),
	}
}
