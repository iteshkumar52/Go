package router

import (
	"store-product-restapi/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/newstore", middleware.CreateStore).Methods("POST", "OPTIONS")
	router.HandleFunc("/store", middleware.GetStores).Methods("GET", "OPTIONS")
	r:=router.PathPrefix("/store").Subrouter()
	r.HandleFunc("/{sid}", middleware.GetStoreById).Methods("GET", "OPTIONS")
	r.HandleFunc("/{sname}", middleware.GetStoreByName).Methods("GET", "OPTIONS")
	router.HandleFunc("/updatestore/{sid}", middleware.UpdateStore).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deletestore/{sid}", middleware.DeleteStore).Methods("DELETE", "OPTIONS")


	router.HandleFunc("/newproduct", middleware.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/product", middleware.GetProducts).Methods("GET", "OPTIONS")
	r = router.PathPrefix("/product").Subrouter()
	r.HandleFunc("/{pid}", middleware.GetProductById).Methods("GET", "OPTIONS")
	r.HandleFunc("/{pname}", middleware.GetProductByName).Methods("GET", "OPTIONS")
	r.HandleFunc("/{exp}", middleware.GetProductByDate).Methods("GET", "OPTIONS")
	router.HandleFunc("/updateproduct/{pid}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteproduct/{pid}", middleware.DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}