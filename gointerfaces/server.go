package main

import "fmt"

type Server interface {
	Find(id int) string
	Save(id int, name string) bool
}

type NewServer struct {
	router string
	port   int
}

func GetServer() *NewServer {
	return &NewServer{
		router: "localhost",
		port:   8080,
	}
}

func (s *NewServer) Find(id int) string {
	return "Find"
}

func main() {
	s := GetServer()
	fmt.Println(s.Find(1))
	fmt.Println(s.Find(2))
}
