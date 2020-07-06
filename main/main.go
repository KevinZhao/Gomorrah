package main

import (
    "fmt"
    "time"
    "Gomorrah/db_operation"
    "Gomorrah/data_source/hkex"
    "github.com/go-co-op/gocron"
)

func main()  {

    //1. init Block
    db_operation.InitDB()

    //2. task scheduler Block
    localtime, _ := time.LoadLocation("Asia/Chongqing")
    s1 := gocron.NewScheduler(localtime)

    //Do hkex holding mining for every day at 4:00 AM
    s1.Every(1).Day().At("4:00").Do(hkex_holding_task)

    //Test
    s1.StartImmediately().Do(hkex_holding_task)
    
    //Do hkex top10
    //s1.Every(1).Day().At("4:00").Do(hkex_holding_task)

    s1.StartBlocking()
}

func hkex_holding_task(){

	fmt.Println("hkex_holding_task executed")

    hkex.GetHolding()
}