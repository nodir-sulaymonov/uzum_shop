package server

import (
	"log"
	"net/http"

	"github.com/Shemistan/uzum_shop/internal/service/shop1"
)

type Server struct {
	Adress  string
	Service shop1.IShopSystemService
}

func New(address string, service shop1.IShopSystemService) *Server {
	return &Server{
		Adress:  address,
		Service: service,
	}
}

func (s *Server) Run() {
	log.Printf("server is running on port %s", s.Adress)
	http.ListenAndServe(s.Adress, nil)
}

func NewAllUsersHandler(service shop1.IShopSystemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// all the logic goes here

		// users := service.GetAllUser()
		// service.

		// var users []User

		// _, err := json.Encode(&users)
		// if err != nil {
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// }

		// w.Write(users)
	}
}
