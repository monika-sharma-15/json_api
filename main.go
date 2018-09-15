package main

import "os"

func main() {
    a := App{}
    a.Initialize( "postgres","appointy","dbname"  )

    a.Run(":8080")
}