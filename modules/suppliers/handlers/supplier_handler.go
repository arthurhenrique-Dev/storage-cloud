package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"storage-api/modules/suppliers/models"
	"storage-api/modules/suppliers/services"
)

type SupplierHandler struct {
	service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

func (h *SupplierHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /suppliers", h.Create)
	mux.HandleFunc("GET /suppliers", h.List)
	mux.HandleFunc("GET /suppliers/{id}", h.GetByID)
	mux.HandleFunc("PUT /suppliers/{id}", h.Update)
	mux.HandleFunc("DELETE /suppliers/{id}", h.Delete)
}

// POST /suppliers
func (h *SupplierHandler) Create(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

// GET /suppliers/{id}
func (h *SupplierHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	supplier, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(supplier)
}

// GET /suppliers
func (h *SupplierHandler) List(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suppliers)
}

// PUT /suppliers/{id}
func (h *SupplierHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var supplier models.Supplier
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	supplier.ID = id
	if err := h.service.Update(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(supplier)
}

// DELETE /suppliers/{id}
func (h *SupplierHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
