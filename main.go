package main

import "mycoin-cli-geth/cmd"

func main() {
	cmd.Execute()
}

//func main() {
//	//trace.Start(os.Stderr)
//	//defer trace.Stop()
//	service,err:=services.GetService()
//	if err != nil {
//		panic(err)
//	}
//	contractAddress:="0x2b8e5b3D4BC8C2cAcF02F3a3E0e9C359405F99e5"
//	toAddress:="0x3C2661a7D8ec0782691d55e3fd6292e82067C2Ef"
//	err=service.MyCoin.Transfer(context.Background(),toAddress,contractAddress,10000000)
//	if err != nil {
//		panic(err)
//	}
//	//manager,err:=contractevents.NewMyCoinSubManager()
//	////defer manager.Close()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//var wg sync.WaitGroup
//	//wg.Add(1)
//	//go func() {
//	//	err:=manager.Run()
//	//	if err != nil {
//	//		fmt.Println(err)
//	//	}
//	//	wg.Done()
//	//}()
//	//address:="0x2b8e5b3D4BC8C2cAcF02F3a3E0e9C359405F99e5"
//	//err=manager.Subscribe(address)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//time.Sleep(time.Second*5)
//	//manager.Close()
//	//wg.Wait()
//}
