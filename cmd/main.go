package main

import s "github.com/Noskine/connection/config/server"

func main() {
	Httpi := s.NewHttpI()

	Httpi.Listen(4030, "Server is running 🚀")
}