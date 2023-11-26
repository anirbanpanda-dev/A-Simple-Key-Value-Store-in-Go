package main

import(
	"time"
)

type CleanupHandler struct{
	checkInterval time.Duration
	storageHandler []StorageHandler
}

func (cln CleanupHandler)cleanUp(){
	for true{
		time.Sleep(cln.checkInterval)
		for i :=0; i<numberOfStorageHandler; i++{
			for key , info := range cln.storageHandler[i].StorageMap{
				if time.Now().Sub(info.startTime) >= storageTime{
					delete(cln.storageHandler[i].StorageMap, key)
				}
			}
		}
		
	}
}