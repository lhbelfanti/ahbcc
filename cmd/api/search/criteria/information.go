package criteria

import (
	"context"
	"sort"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/categorized"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// Information returns all the search criteria executions information for a given user ID. It includes
	// the number of tweets retrieved and the number of tweets analyzed by the user, ordered by month and year.
	Information func(ctx context.Context, token string) (InformationDTOs, error)

	// SummarizedInformation returns a search criteria execution information filtering by a given month and year. It includes
	// the number of tweets retrieved and the number of tweets analyzed by the user, the criteria name and the criteria ID
	SummarizedInformation func(ctx context.Context, token string, criteriaID int, year int, month int) (SummarizedInformationDTO, error)
)

// MakeInformation creates a new Information
func MakeInformation(selectUserIDByToken session.SelectUserIDByToken, selectAllCriteriaExecutionsSummaries summary.SelectAll, selectAllSearchCriteria SelectAll, selectAllCategorizedTweets categorized.SelectAllByUserID) Information {
	return func(ctx context.Context, token string) (InformationDTOs, error) {
		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveUserID
		}
		ctx = log.With(ctx, log.Param("user_id", userID))

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

// MakeSummarizedInformation creates a new SummarizedInformation
func MakeSummarizedInformation(selectUserIDByToken session.SelectUserIDByToken, selectCriteriaByID SelectByID, selectAllSummarized summary.SelectAll, selectAllCategorized categorized.SelectAllByUserID) SummarizedInformation {
	return func(ctx context.Context, token string, criteriaID int, year int, month int) (SummarizedInformationDTO, error) {
		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return SummarizedInformationDTO{}, FailedToRetrieveUserID
		}
		ctx = log.With(ctx, log.Param("user_id", userID))

		criteriaInfo, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return SummarizedInformationDTO{}, FailedToRetrieveSearchCriteria
		}

		summarizedTweets, err := selectAllSummarized(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return SummarizedInformationDTO{}, FailedToRetrieveSearchCriteriaExecutionsSummaries
		}

		categorizedTweets, err := selectAllCategorized(ctx, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return SummarizedInformationDTO{}, FailedToRetrieveCategorizedTweetsByUserID
		}

		var totalTweets int
		for _, current := range summarizedTweets {
			if current.SearchCriteriaID != criteriaID {
				continue
			}

			if (year == 0 || current.Year == year) && (month == 0 || current.Month == month) {
				totalTweets += current.Total
			}
		}

		var analyzedTweets int
		for _, current := range categorizedTweets {
			if current.SearchCriteriaID != criteriaID {
				continue
			}

			if (year == 0 || current.Year == year) && (month == 0 || current.Month == month) {
				analyzedTweets += current.Analyzed
			}
		}

		return SummarizedInformationDTO{
			ID:             criteriaID,
			Name:           criteriaInfo.Name,
			Year:           year,
			Month:          month,
			AnalyzedTweets: analyzedTweets,
			TotalTweets:    totalTweets,
		}, nil
	}
}
