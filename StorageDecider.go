package main


import(
	"fmt"
)

type StorageDecider struct{
	dividerBit byte 

}

func (decider StorageDecider) DecideStorage(key string) int{
	fmt.Println(key)
	if key[0]-'a' < decider.dividerBit{
		fmt.Println("Server 0...")
		return 0
	} else {
		fmt.Println("Server 1...")
		return 1 
	}
}