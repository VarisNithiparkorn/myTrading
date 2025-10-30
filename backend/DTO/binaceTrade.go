package dto

type CoinPriceResponse struct {
    Symbol              string  `json:"symbol"`
    LastPrice           float64 `json:"last_price"` 
    PriceChange         float64 `json:"price_change"` 
    PriceChangePercent  float64 `json:"price_change_percent"` 
}

type ResponseData struct {

    TotalCoins int                     `json:"total_coins"`

    Data       []CoinPriceResponse     `json:"data"`

}
type BinanceTicker struct {
    // ข้อมูล Event และ Symbol
    EventType string `json:"e"` // e: Event type ("24hrTicker")
    EventTime int64  `json:"E"` // E: Event Time (Timestamp in ms)
    Symbol    string `json:"s"` // s: Symbol (e.g., "SEIUSDC")
    
    // ราคาและสถิติการเปลี่ยนแปลง
    PriceChange        string `json:"p"` // p: Price change (ราคาเปลี่ยนแปลง)
    PriceChangePercent string `json:"P"` // P: Price change percent (%)
    WeightedAvgPrice   string `json:"w"` // w: Weighted average price
    
    // ราคาปิดล่าสุด (Last Price)
    // นี่คือ Field ที่ใช้ในการอัปเดตแคชหลัก
    LastPrice          string `json:"c"` // c: Last Price
    
    // ปริมาณการซื้อขายล่าสุด
    LastQuantity       string `json:"Q"` // Q: Last trade quantity (ปริมาณซื้อขายล่าสุด)
    
    // ข้อมูล Bid/Ask (Best Price)
    BestBidPrice       string `json:"b"` // b: Best Bid Price
    BestBidQty         string `json:"B"` // B: Best Bid Quantity
    BestAskPrice       string `json:"a"` // a: Best Ask Price
    BestAskQty         string `json:"A"` // A: Best Ask Quantity
    
    // ราคา Open/High/Low
    OpenPrice          string `json:"o"` // o: Open Price (ราคาเปิด)
    HighPrice          string `json:"h"` // h: 24hr High Price
    LowPrice           string `json:"l"` // l: 24hr Low Price
    
    // ปริมาณการซื้อขาย
    BaseVolume         string `json:"v"` // v: Total traded base asset volume
    QuoteVolume        string `json:"q"` // q: Total traded quote asset volume
    
    // ข้อมูล Timestamp และ Count
    OpenTime           int64  `json:"O"` // O: Open Time (Timestamp)
    CloseTime          int64  `json:"C"` // C: Close Time (Timestamp)
    FirstTradeID       int64  `json:"F"` // F: First trade ID
    LastTradeID        int64  `json:"L"` // L: Last trade ID
    TradeCount         int64  `json:"n"` // n: Number of trades
}