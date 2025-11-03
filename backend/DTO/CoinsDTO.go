package dto

type CombinedStreamMessage struct {
    Stream string `json:"stream"`
    Data   BinanceTicker `json:"data"` 
}