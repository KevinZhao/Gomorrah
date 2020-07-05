package main

import (
    "Gomorrah/db_operation"
    "Gomorrah/data_source/hkex"
)

func main()  {

    db_operation.InitDB()
  
  	hkex.Tbd()
}
