package test

import (
	"log"
	"ordersystem/common"
	"testing"
)

func TestFilePath(t *testing.T) {
	rootPath := common.RootPath()
	log.Println("TestFilePath.rootPath:", rootPath)
	// assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}

func TestGetCurrentAbPath(t *testing.T) {
	rootPath := common.GetCurrentAbPath()
	log.Println("TestGetCurrentAbPath.rootPath:", rootPath)
}
