package main

import (

   "log"
   "loremipsumbytes/bootstrap"
 )

func main() {
  
	app := bootstrap.NewApplication()
  log.Fatal(app.Listen(":5001"))

 
}

 