package main

type server struct {
}

func (s *server) Init() {

}

func main() {
	(&server{}).Init()
}
