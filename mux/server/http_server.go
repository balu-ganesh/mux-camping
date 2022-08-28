package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router  http.Handler
	isLocal bool
}

func New(options ...func(*HTTPServer)) (h *HTTPServer) {
	h = &HTTPServer{}
	for _, option := range options {
		option(h)
	}

	router := mux.NewRouter()
	// router.NotFoundHandler = http.HandlerFunc(h.notFound)

	//list all available slots
	router.HandleFunc("/api/slots", h.GetAllAvailableSlotsByDate).Methods("GET")
	//Reserve slot for an user
	router.HandleFunc("/api/user/{userid}/slots/{slotid}/book", h.ReserveSlot).Methods("POST")

	//GET ALL SLOTS BY USER ID
	router.HandleFunc("/api/user/{id}/slot", h.GetSlotsByUser).Methods("GET")

	//Manage a user slot
	router.HandleFunc("/api/user/{userid}/slots/{slotid}/reserve", h.ManageSlot).Methods("PUT")

	//Cancel the selected slot
	router.HandleFunc("/api/user/{userid}/slots/{slotid}/cancel", h.CancelBooking).Methods("PATCH")

	h.router = router
	return h
}

func Local(isLocal bool) func(h *HTTPServer) {
	return func(h *HTTPServer) {
		h.isLocal = true
	}
}

func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type errorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func doWriteHTTPError(res http.ResponseWriter, err error, code int) {
	info := &errorInfo{Message: err.Error(), Code: code}
	b, _ := json.Marshal(info)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(b)
}

func writeHTTPError(res http.ResponseWriter, err interface{}) {
	var statusCode int
	var httpError error

	defer func() {
		doWriteHTTPError(res, httpError, statusCode)
	}()

	if brerr, ok := err.(*BadRequestError); ok {
		statusCode = http.StatusBadRequest
		httpError = brerr
		return
	}
	statusCode = http.StatusInternalServerError
	httpError = fmt.Errorf("%v", err)
}

func writeHTTPResponse(res http.ResponseWriter, val interface{}) {
	jsonData, err := json.Marshal(val)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
