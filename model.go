package main


// SnowflakNextResponse
type SnowflakNextResponse struct {
	Id int64 `json:"id"`
}

// SnowflakBatchResponse
type SnowflakBatchResponse struct {
	Ids []int64 `json:"ids"`
}