package cachedb

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockDatabase "test-payment-system/internal/app/payment/database/mock"
	"test-payment-system/internal/app/payment/database/model"
	"testing"
	"time"
)

func TestCacheDB_GetWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mockDatabase.NewMockDatabase(ctrl)

	cache := New(dbMock)

	ctx := context.Background()
	walletID := uint(123)

	//error db
	t.Run("error_from_db", func(t *testing.T) {
		wantError := errors.New("error 1")
		dbMock.EXPECT().GetWallet(ctx, walletID).Return(nil, wantError)
		_, err := cache.GetWallet(ctx, walletID)
		assert.EqualError(t, err, wantError.Error())
		key := getKeyWallet(walletID)
		_, ok := cache.cache.Get(key)
		assert.False(t, ok)
	})
	t.Run("from_cache", func(t *testing.T) {
		wantWallet := &model.Wallet{
			ID: walletID,
		}
		key := getKeyWallet(walletID)
		cache.cache.Set(key, wantWallet, time.Millisecond*500)
		defer cache.cache.Delete(key)
		wallet, err := cache.GetWallet(ctx, walletID)
		assert.NoError(t, err)
		assert.Equal(t, wantWallet, wallet)
	})
	t.Run("set_to_cache", func(t *testing.T) {
		wantWallet := &model.Wallet{
			ID: walletID,
		}
		dbMock.EXPECT().GetWallet(ctx, walletID).Return(wantWallet, nil)

		wallet, err := cache.GetWallet(ctx, walletID)
		assert.Equal(t, wantWallet, wallet)
		assert.NoError(t, err)

		key := getKeyWallet(walletID)
		gotWallet, ok := cache.cache.Get(key)
		assert.True(t, ok)
		assert.Equal(t, wantWallet, gotWallet)
	})
}
