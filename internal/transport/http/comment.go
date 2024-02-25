package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	coment "github.com/instinctG/Rest-API/internal/comment"
	"net/http"
	"strconv"
)

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint id", err)
		return
	}

	comment, err := h.Service.GetComment(context.Background(), uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Error in retrieving comment by id", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments(context.Background())
	if err != nil {
		sendErrorResponse(w, "Error of retrieving all comments:", err)
		return
	}

	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var cmt coment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	cmt, err := h.Service.PostComment(context.Background(), cmt)
	if err != nil {
		sendErrorResponse(w, "Failed to post a new comment", err)
		return
	}

	if err := sendOkResponse(w, cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment coment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse id to uint", err)
		return
	}

	comment, err = h.Service.UpdateComment(context.Background(), uint(commentID), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error in parsing id to uint", err)
		return
	}

	err = h.Service.DeleteComment(context.Background(), uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by comment ID", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
}
