package executions

import "context"

// MockInsertExecution mocks InsertExecution function
func MockInsertExecution(criteriaID int, err error) InsertExecution {
	return func(ctx context.Context, searchCriteriaID int, forced bool) (int, error) {
		return criteriaID, err
	}
}

// MockUpdateExecution mocks UpdateExecution function
func MockUpdateExecution(err error) UpdateExecution {
	return func(ctx context.Context, executionID int, status string) error {
		return err
	}
}

// MockInsertExecutionDay mocks InsertExecutionDay function
func MockInsertExecutionDay(err error) InsertExecutionDay {
	return func(ctx context.Context, executionDay ExecutionDayDTO) error {
		return err
	}
}

// MockSelectExecutionsByStatuses mocks SelectExecutionsByStatuses function
func MockSelectExecutionsByStatuses(executionsDAO []ExecutionDAO, err error) SelectExecutionsByStatuses {
	return func(ctx context.Context, statuses []string) ([]ExecutionDAO, error) {
		return executionsDAO, err
	}
}

// MockSelectLastDayExecutedByCriteriaID mocks SelectLastDayExecutedByCriteriaID function
func MockSelectLastDayExecutedByCriteriaID(lastDayExecuted ExecutionDayDAO, err error) SelectLastDayExecutedByCriteriaID {
	return func(ctx context.Context, id int) (ExecutionDayDAO, error) {
		return lastDayExecuted, err
	}
}

// MockSelectExecutionByID mocks SelectExecutionByID function
func MockSelectExecutionByID(executionDAO ExecutionDAO, err error) SelectExecutionByID {
	return func(ctx context.Context, id int) (ExecutionDAO, error) {
		return executionDAO, err
	}
}

// MockExecutionDAOValues mocks the properties of ExecutionDAO to be used in the Scan function
func MockExecutionDAOValues(dao ExecutionDAO) []any {
	return []any{
		dao.ID,
		dao.Status,
		dao.SearchCriteriaID,
	}
}

// MockExecutionDAO mocks an ExecutionDAO
func MockExecutionDAO() ExecutionDAO {
	return ExecutionDAO{
		ID:               1,
		Status:           DoneStatus,
		SearchCriteriaID: 2,
	}
}

// MockExecutionsDAO mocks a slice of ExecutionDAO
func MockExecutionsDAO() []ExecutionDAO {
	return []ExecutionDAO{
		{
			ID:               1,
			Status:           PendingStatus,
			SearchCriteriaID: 2,
		},
		{
			ID:               2,
			Status:           InProgressStatus,
			SearchCriteriaID: 4,
		},
	}
}

// MockExecutionDayDTO mocks an ExecutionDayDTO
func MockExecutionDayDTO(errorReason *string) ExecutionDayDTO {
	return ExecutionDayDTO{
		ExecutionDate:             "2006-01-01",
		TweetsQuantity:            20,
		ErrorReason:               errorReason,
		SearchCriteriaExecutionID: 5,
	}
}

// MockExecutionDTO mocks an ExecutionDTO
func MockExecutionDTO() ExecutionDTO {
	return ExecutionDTO{
		Status:           "DONE",
		SearchCriteriaID: 0,
	}
}
