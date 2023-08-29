package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"main/internal/service"
	"main/pkg/models"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{Service: srv}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("login")
	password := r.Header.Get("password")
	activeUser, id := h.Service.CheckUser(login, password)
	//activeUser, id := service.CheckUser(login, password, h.Service.Db)
	if activeUser == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var NewContent models.Note
	err := json.NewDecoder(r.Body).Decode(&NewContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Service.AddNoteToDb(id, NewContent)
	//err = service.CreateNote(id, NewContent, h.Service.Db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("login")
	password := r.Header.Get("password")
	active, userId := h.Service.CheckUser(login, password)
	if active == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	note, err := h.Service.Read(userId, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(note, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("login")
	password := r.Header.Get("password")

	active, userId := h.Service.CheckUser(login, password)
	if active == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)

	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var NewContent models.Note
	err = json.NewDecoder(r.Body).Decode(&NewContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	updNote, err := h.Service.UpdateNote(userId, id, NewContent)

	bytes, err := json.MarshalIndent(updNote, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write(bytes)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("login")
	password := r.Header.Get("password")

	active, userId := h.Service.CheckUser(login, password)
	if active == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteNote(userId, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusGone)
}
