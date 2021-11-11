package service

// ProcessInventoryEvent Process Inventory Event related operations.
type ProcessInventoryEvent interface {
	CountTodaysRanProcessByCompanyId(companyId string) int64
}
