package main

import (
    "code.google.com/p/go.net/websocket"
    "flag"
    "fmt"
    "encoding/json"
)

func init() {
    flag.Parse()
}

//=======================
type LoginReply struct {
    Op string `json:"op"`
    Message string `json:"message"`
}

type Trade struct {
    Amount float64 `json:"amount"`
    AmountInt string `json:"amount_int"`
    Date int `json:"date"`
    Item string `json:"item"`
    Price float64 `json:"price"`
    PriceCurrency string `json:"price_currency"`
    PriceInt string `json:"price_int"`
    Primary string `json:"primary"`
    Properties string `json:"properties"`
    Tid string `json:"tid"`
    TradeType string `json:"bid"`
    Type string `json:"type"`
}

type TradeFeed struct {
    Channel string `json:"channel"`
    Op string `json:"op"`
    Origin string `json:"origin"`
    Private string `json:"private"`
    Trade `json:"trade"`
}

//=======================

func main() {
    src := "ws://localhost/"
    dst := "ws://websocket.mtgox.com:80"
    
    ws, err := websocket.Dial(dst, "ws", src)
    
    if err != nil {
        panic(err.Error())
    }
    
    decoder := json.NewDecoder(ws)
    encoder := json.NewEncoder(ws)
    
    remark := &LoginReply{}
    decoder.Decode(remark)
    fmt.Println(remark, remark.Op, remark.Message)
    
    // subscribe
    subscription := map[string]string{"op": "mtgox.subscribe", "type": "trades"}
    err = encoder.Encode(&subscription)
    
    if err != nil {
        panic(err.Error())
    }
    
    msg := &TradeFeed{}
    
    for {
        err := decoder.Decode(&msg)
        
        if err != nil {
            panic(err.Error())
        }
        
        fmt.Println(msg.Trade, msg.Trade.Amount, msg.Trade.Price, msg.Trade.Item)
    }
}