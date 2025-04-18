package criteria

import (
	"context"
	"sort"

	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/log"
)

// Information returns all the search criteria executions information for a given user ID. It includes
// the number of tweets retrieved and the number of tweets analyzed by the user, ordered by month and year.
type Information func(ctx context.Context, token string) (InformationDTOs, error)

// MakeInformation creates a new Information
func MakeInformation(selectUserIDByToken session.SelectUserIDByToken, selectAllCriteriaExecutionsSummaries summary.SelectAll, selectAllSearchCriteria SelectAll, selectAllCategorizedTweets categorized.SelectAllByUserID) Information {
	return func(ctx context.Context, token string) (InformationDTOs, error) {
		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveUserID
		}

		criteriaExecutionsSummaries, err := selectAllCriteriaExecutionsSummaries(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveSearchCriteriaExecutionsSummaries
		}

		searchCriteria, err := selectAllSearchCriteria(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveSearchCriteria
		}

		categorizedTweets, err := selectAllCategorizedTweets(ctx, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveCategorizedTweetsByUserID
		}

		executionsMap := make(map[int][]summary.DAO)
		for _, executionSummary := range criteriaExecutionsSummaries {
			executionsMap[executionSummary.SearchCriteriaID] = append(executionsMap[executionSummary.SearchCriteriaID], executionSummary)
		}

		information := make(InformationDTOs, 0, len(criteriaExecutionsSummaries))
		for searchCriteriaID, executionsSummaries := range executionsMap {
			yearMap := make(map[int][]summary.DAO)
			for _, executionSummary := range executionsSummaries {
				yearMap[executionSummary.Year] = append(yearMap[executionSummary.Year], executionSummary)
			}

			years := make(YearDataDTOs, 0, len(yearMap))
			for year, summaries := range yearMap {
				monthMap := make(map[int]summary.DAO)
				for _, s := range summaries {
					monthMap[s.Month] = s
				}

				months := make(MonthDataDTOs, 0, len(monthMap))
				for month, s := range monthMap {
					analyzedTweets := 0
					for _, ct := range categorizedTweets {
						if ct.SearchCriteriaID == s.SearchCriteriaID && ct.Year == year && ct.Month == month {
							analyzedTweets = ct.Analyzed
							break
						}
					}

					months = append(months, MonthDataDTO{
						Month:          month,
						AnalyzedTweets: analyzedTweets,
						TotalTweets:    s.Total,
					})
				}

				sort.Sort(months)

				years = append(years, YearDataDTO{
					Year:   year,
					Months: months,
				})
			}

			var criteriaName string
			for _, crit := range searchCriteria {
				if crit.ID == searchCriteriaID {
					criteriaName = crit.Name
				}
			}

			sort.Sort(years)

			information = append(information, InformationDTO{
				Name:  criteriaName,
				ID:    searchCriteriaID,
				Years: years,
			})
		}

		sort.Sort(information)

		return information, nil
	}
}
