package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
	"github.com/aren55555/grpc-example/parking_server/storage"
)

// Our implementation of the parking service
type handler struct {
	storage storage.Storer
}

func NewHandler(storage storage.Storer) http.Handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	switch req.Method {
	case http.MethodGet:
		h.list(res, req)
	case http.MethodDelete:
		h.delete(res, req)
	case http.MethodPost:
		h.create(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *handler) list(res http.ResponseWriter, req *http.Request) {
	m := h.storage.GetMap()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(m)
}

func (h *handler) create(res http.ResponseWriter, req *http.Request) {
	car := &pb.Car{}

	if err := json.NewDecoder(req.Body).Decode(car); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(Error{Message: fmt.Sprintf("invalid request body: %v", err)})
		res.Write(data)
		return
	}

	cd, err := h.storage.AddCar(car)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(Error{Message: err.Error()})
		res.Write(data)
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(cd)
}

func (h *handler) delete(res http.ResponseWriter, req *http.Request) {
	// TODO: parse query param remove Car
	id := req.URL.Query().Get("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(Error{Message: "id query param is required"})
		res.Write(data)
		return
	}

	cd, err := h.storage.RemoveCar(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(Error{Message: err.Error()})
		res.Write(data)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(cd)
}
