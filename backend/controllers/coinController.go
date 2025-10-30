package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	dto "github.com/VarisNithiparkorn/cryptoGraph/backend/DTO"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
type CombinedStreamMessage struct {
    Stream string `json:"stream"`
    Data   dto.BinanceTicker `json:"data"` // 💡 รับ BinanceTicker Object เดี่ยว
}

var (

	globalCoinData = make(map[string]dto.CoinPriceResponse) 

	cacheMutex sync.RWMutex 
)

//const wsURL = "wss://stream.binance.com:9443/ws/!ticker@arr"
const wsURL = "wss://stream.binance.com/stream?streams=btcusdt@ticker/ethusdt@ticker/solusdt@ticker"



func RunWebSocketClient() {
	log.Printf("Starting WebSocket connection to: %s", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("WebSocket Dial failed:", err)
		return
	}
	defer conn.Close()
	for {

        _, message, err := conn.ReadMessage() 
        if err != nil {
            log.Println("Read Message failed:", err); break 
        }

        var msg CombinedStreamMessage 
        

        if err := json.Unmarshal(message, &msg); err == nil {
            ticket := msg.Data 

            cacheMutex.Lock()
            
            lastPrice, err := strconv.ParseFloat(ticket.LastPrice, 64)
            if err != nil { log.Printf("Parse price error for %s: %v", ticket.Symbol, err); lastPrice = 0.0 }
            
            priceChange, err := strconv.ParseFloat(ticket.PriceChange, 64)
            if err != nil { priceChange = 0.0 }
            
            priceChangePercent, err := strconv.ParseFloat(ticket.PriceChangePercent, 64)
            if err != nil { priceChangePercent = 0.0 }
                
            globalCoinData[ticket.Symbol] = dto.CoinPriceResponse{
                Symbol:             ticket.Symbol,
                LastPrice:          lastPrice,
                PriceChange:        priceChange,
                PriceChangePercent: priceChangePercent,
            }
            
            log.Printf("Updated: %s = %.2f", ticket.Symbol, lastPrice)
            cacheMutex.Unlock()
            
        } else {
            // Log Error ถ้า Unmarshal ล้มเหลว
            log.Printf("Unmarshal Failed: %v, RAW JSON: %s", err, string(message))
        }
    }
    log.Println("WebSocket client terminated. Attempting reconnect...")
}
func HandleGetAllCoins(c *gin.Context) {

	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
    c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    c.Writer.Header().Set("Pragma", "no-cache")
    c.Writer.Header().Set("Expires", "0")

	var coinList []dto.CoinPriceResponse
	for _, coin := range globalCoinData {
		coinList = append(coinList, coin)
	}

	finalResponse := dto.ResponseData{
		TotalCoins: len(coinList),
		Data: coinList,
	}


	c.JSON(http.StatusOK, finalResponse)
}
	
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true 
    },
}

func HandleWebSocketPrices(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("WebSocket Upgrade failed:", err)
        return
    }
    defer conn.Close()

    for {
        cacheMutex.RLock()
        
        var coinList []dto.CoinPriceResponse
        for _, coin := range globalCoinData {
            coinList = append(coinList, coin)
        }
        
        finalResponse := dto.ResponseData{
            TotalCoins: len(coinList),
            Data: coinList,
        }
        
        cacheMutex.RUnlock() 

        
        if err := conn.WriteJSON(finalResponse); err != nil {
            log.Println("WebSocket Write failed, closing connection:", err)
            break 
        }

        time.Sleep(1 * time.Second) 
    }
}