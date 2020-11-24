package cachedb

import (
	"context"
	"test-payment-system/internal/app/payment/database"
	"test-payment-system/internal/app/payment/database/model"
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheDB in memory cache for Database
// This is for example. The in memory cache was chosen so as not to deploy an external service Redis, Memcached, etc.
// Memory cache should not be used in such a case,
// since when scaling horizontally, the cache will not work in other instances.
// Also, using this template, you can implement logging, open tracing, repetitions, etc.
type CacheDB struct {
	database.Database
	cache *cache.Cache
}

var _ database.Database = (*CacheDB)(nil)

const (
	keyWallet = "wallet"
)

const (
	expirationWallet = time.Minute * 2
)

// New returns new CacheDB
func New(db database.Database) *CacheDB {
	defaultExpiration, cleanupInterval := time.Minute*5, time.Minute*10
	newCache := &CacheDB{
		Database: db,
		cache:    cache.New(defaultExpiration, cleanupInterval),
	}
	return newCache
}

func (c *CacheDB) GetWallet(ctx context.Context, id uint) (*model.Wallet, error) {
	if wallet, ok := c.cache.Get(keyWallet); ok {
		return wallet.(*model.Wallet), nil
	}

	wallet, err := c.Database.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	c.cache.Set(keyWallet, wallet, expirationWallet)

	return wallet, nil
}