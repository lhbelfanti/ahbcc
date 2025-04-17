package user

import (
	"context"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/internal/log"
)

// Information returns all the search criteria executions information for a given user ID. It includes
// the number of tweets retrieved and the number of tweets analyzed by the user, ordered by month and year.
type Information func(ctx context.Context, userID int) ([]criteria.SearchCriteriaInformation, error)

// MakeInformation creates a new Information
func MakeInformation(selectAllCriteriaExecutionsSummaries summary.SelectAll, selectAllSearchCriteria criteria.SelectAll, selectAllCategorizedTweets categorized.SelectAllByUserID) Information {
	return func(ctx context.Context, userID int) ([]criteria.SearchCriteriaInformation, error) {
		criteriaExecutionsSummaries, err := selectAllCriteriaExecutionsSummaries(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, criteria.FailedToRetrieveSearchCriteriaExecutionsSummaries
		}

		searchCriteria, err := selectAllSearchCriteria(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, criteria.FailedToRetrieveSearchCriteria
		}

		categorizedTweets, err := selectAllCategorizedTweets(ctx, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, criteria.FailedToRetrieveCategorizedTweetsByUserID
		}

		executionsMap := make(map[int][]summary.DAO)
		for _, executionSummary := range criteriaExecutionsSummaries {
			executionsMap[executionSummary.SearchCriteriaID] = append(executionsMap[executionSummary.SearchCriteriaID], executionSummary)
		}

		information := make([]criteria.SearchCriteriaInformation, 0, len(criteriaExecutionsSummaries))
		for searchCriteriaID, executionsSummaries := range executionsMap {
			yearMap := make(map[int][]summary.DAO)
			for _, executionSummary := range executionsSummaries {
				yearMap[executionSummary.Year] = append(yearMap[executionSummary.Year], executionSummary)
			}

			years := make([]criteria.YearData, 0, len(yearMap))
			for year, summaries := range yearMap {
				monthMap := make(map[int]summary.DAO)
				for _, s := range summaries {
					monthMap[s.Month] = s
				}

				months := make([]criteria.MonthData, 0, len(monthMap))
				for month, s := range monthMap {
					analyzedTweets := 0
					for _, ct := range categorizedTweets {
						if ct.SearchCriteriaID == s.SearchCriteriaID && ct.Year == year && ct.Month == month {
							analyzedTweets = ct.Analyzed
							break
						}
					}

					months = append(months, criteria.MonthData{
						Month:          month,
						AnalyzedTweets: analyzedTweets,
						TotalTweets:    s.Total,
					})
				}

				years = append(years, criteria.YearData{
					Year:   year,
					Months: months,
				})
			}

			var criteriaName string
			for _, criteria := range searchCriteria {
				if criteria.ID == searchCriteriaID {
					criteriaName = criteria.Name
				}
			}

			information = append(information, criteria.SearchCriteriaInformation{
				Name:  criteriaName,
				ID:    searchCriteriaID,
				Years: years,
			})
		}

		return information, nil
	}
}
