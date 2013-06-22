package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

const (
	SleepInterval = time.Second
)

type Client struct {
	feeds       chan *common.Feed
	tickerFeeds bool
	orderBookFeeds   bool
}

func NewClient() *Client {
	return &Client{make(chan *common.Feed), false, false}
}

func (c *Client) ToggleTickerFeeds() {
	c.tickerFeeds = !c.tickerFeeds
}

func (c *Client) ToggleOrderBookFeeds() {
    c.orderBookFeeds = !c.orderBookFeeds
}

func (c *Client) ToggleAsync() {
	go c.async()
}

func (c *Client) Channel() <-chan *common.Feed {
	return c.feeds
}

func (c *Client) async() {
	for {
		time.Sleep(SleepInterval)
		if c.tickerFeeds {
			c.pullTickerFeed()
		}
	}
}

func (c *Client) pullTickerFeed() {
	ticker := GetTicker()
	feed := &common.Feed{
		common.TickerFeed,
		ticker.GetTimestamp(),
		ticker}
	c.feeds <- feed
}
