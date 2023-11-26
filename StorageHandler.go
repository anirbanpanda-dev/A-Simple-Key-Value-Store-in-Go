package main

import(
	"time"
	"sync"
)

var storageMtx sync.Mutex

type StorageInformation struct{
	entryID int
	value string
	startTime time.Time
}

type StorageHandler struct {
	StorageMap map[string]StorageInformation
}

func (handler StorageHandler) getValue(key string) string {
	return handler.StorageMap[key].value
}

func (handler StorageHandler) setValue(key string, value string){
	storageMtx.Lock()
	handler.StorageMap[key] = StorageInformation{ entryID: 0, value: value, startTime: time.Now()}
	storageMtx.Unlock()
}




