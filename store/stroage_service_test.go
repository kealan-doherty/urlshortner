package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStorageService = &StorageService{}

func init() {
	testStorageService = InitializeStorageService()
}

// basic test to check the redis clinet has been initalized properly and is not nil
func TestStoreInit(t *testing.T) {
	assert.True(t, testStorageService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	initallink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUid := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "abc123"

	SaveUrlMapping(shortUrl, initallink, userUUid)

	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initallink, retrievedUrl)
}
