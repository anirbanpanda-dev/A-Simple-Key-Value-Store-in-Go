package main

type KVStoreMgr struct{
	storageHandler[] StorageHandler
	clnHandler CleanupHandler
	decider StorageDecider
}

func (kvStoreMgr KVStoreMgr) init(){
	go kvStoreMgr.clnHandler.cleanUp()
}


func (kvStoreMgr KVStoreMgr) getValue(key string) string {
	var strHandler StorageHandler = kvStoreMgr.storageHandler[kvStoreMgr.decider.DecideStorage(key)]
	return strHandler.getValue(key)
}

func (kvStoreMgr KVStoreMgr) setValue(key string, value string){
	var strHandler StorageHandler = kvStoreMgr.storageHandler[kvStoreMgr.decider.DecideStorage(key)]
	strHandler.setValue(key, value)
}