package main



func mirroredQuery() string {
	response := make(chan string,3)
	go func() { response <- request("asia.gopl.io") }()
	go func() { response <- request("europe.gopl.io") }()
    go func() { response <- request("americas.gopl.io") }()
	return <- response
}

func main()  {
	mirroredQuery()
}