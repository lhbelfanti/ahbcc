package summary

import (
	"context"
	
	"github.com/jackc/pgx/v5"
)

// MockSelectIDBySearchCriteriaIDYearAndMonth mocks SelectIDBySearchCriteriaIDYearAndMonth function
func MockSelectIDBySearchCriteriaIDYearAndMonth(id int, err error) SelectIDBySearchCriteriaIDYearAndMonth {
	return func(ctx context.Context, searchCriteriaID, year, month int) (int, error) {
		return id, err
	}
}

// MockSelectMonthlyTweetsCountsByYearByCriteriaID mocks SelectMonthlyTweetsCountsByYearByCriteriaID function
func MockSelectMonthlyTweetsCountsByYearByCriteriaID(daos []DAO, err error) SelectMonthlyTweetsCountsByYearByCriteriaID {
	return func(ctx context.Context, criteriaID int) ([]DAO, error) {
		return daos, err
	}
}

// MockInsert mocks Insert function
func MockInsert(insertedRowID int, err error) Insert {
	return func(tx pgx.Tx, ctx context.Context, dao DAO) (int, error) {
		return insertedRowID, err
	}
}

// MockUpdateTotalTweets mocks UpdateTotalTweets function
func MockUpdateTotalTweets(err error) UpdateTotalTweets {
	return func(tx pgx.Tx, ctx context.Context, id, totalTweets int) error {
		return err
	}
}

// MockUpsert mocks Upsert function
func MockUpsert(err error) Upsert {
	return func(ctx context.Context, tx pgx.Tx, executionSummary DAO) error {
		return err
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
		MockExecutionSummaryDAO(1, 2024, 9, 350),
		MockExecutionSummaryDAO(1, 2025, 1, 1000),
		MockExecutionSummaryDAO(1, 2025, 2, 500),
	}
}
