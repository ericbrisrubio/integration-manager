package service

type ProcessInventoryEvent interface {
	CountTodaysRanProcessByCompanyId(companyId string)int64
}

