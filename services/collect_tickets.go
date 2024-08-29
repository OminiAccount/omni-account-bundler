package services

import (
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

type CollectTicketsChannel struct {
	db      ethdb.Database
	ticker  *time.Ticker
	channel chan bool
	logger  log.Logger
}
