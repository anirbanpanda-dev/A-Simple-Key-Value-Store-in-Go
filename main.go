package main


import(
	"fmt"
	"time"
)

func test(kvStore KVStoreMgr){
	kvStore.setValue("anirban", "10101")
	fmt.Println("time : 00:00 -> ", kvStore.getValue("anirban"))
	fmt.Println("time : 00:00 -> ", kvStore.getValue("panda"))
	time.Sleep(1 * time.Minute)
	kvStore.setValue("panda", "011")
	fmt.Println("time : 01:00 -> ", kvStore.getValue("anirban"))
	fmt.Println("time : 00:00 -> ", kvStore.getValue("panda"))
	time.Sleep(1 * time.Minute)
	fmt.Println("time : 02:00 -> ", kvStore.getValue("anirban"))
	fmt.Println("time : 00:00 -> ", kvStore.getValue("panda"))
	time.Sleep(1 * time.Minute)
	fmt.Println("time : 03:00 -> ", kvStore.getValue("anirban"))
	fmt.Println("time : 00:00 -> ", kvStore.getValue("panda"))
}



func main(){
	fmt.Println("Starting ...");
	var strHandler []StorageHandler
	for i :=0 ; i<numberOfStorageHandler; i++ {
		strHandler = append(strHandler, StorageHandler{StorageMap : make(map[string]StorageInformation)})
	}

	var cln CleanupHandler = CleanupHandler{checkInterval : cleanUpInterval, 
							storageHandler : strHandler}
	var _decider StorageDecider  = StorageDecider{dividerBit : 10}
	var kvStore KVStoreMgr = KVStoreMgr{storageHandler : strHandler, clnHandler : cln, decider : _decider}
	kvStore.init()
	test(kvStore)
}