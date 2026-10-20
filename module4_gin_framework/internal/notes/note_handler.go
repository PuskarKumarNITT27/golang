package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler{
	return &Handler{repo:repo}
}

func (h *Handler) CreateNote(c *gin.Context) {
	var req CreateNoteRequest

	if err:= c.ShouldBindJSON(&req) ; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid JSON",
		})
	}

	now := time.Now().UTC()

	note := Note{
		ID: bson.NewObjectID(),

		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,

		CreatedAt: now,
		UpdatedAt: now,
	}

	created,err := h.repo.Create(c.Request.Context(),note)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to create Note",
		})
	}

	c.JSON(http.StatusCreated,created)
}


func (h *Handler) ListNotes (c *gin.Context){
	notes ,err := h.repo.List(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to fetch all notes",
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"notes":notes,
	})
}

func (h *Handler) GetNoteById(c *gin.Context){
	idStr := c.Param("id")

	objID, err := bson.ObjectIDFromHex(idStr) // convert to mongodb object id type
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid ID",
		})
		return
	}

	note ,err :=  h.repo.GetById(c.Request.Context(),objID)
	if err != nil {
		if errors.Is(err,mongo.ErrNoDocuments){
			c.JSON(http.StatusNotFound,gin.H{
				"error":"Note Not found",
			})
			return 
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Internal Server Error",
		})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"note":note,
	})
}

func (h *Handler) UpdateNoteById(c *gin.Context){
	idStr := c.Param("id")

	objID, err := bson.ObjectIDFromHex(idStr) // convert to mongodb object id type
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid ID",
		})
		return
	}

	var req UpdateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid json format",
		})
		return 
	}

	updated,err := h.repo.UpdateById(c.Request.Context(),objID,req)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Couldn't update the Note",
		})
		return 
	}
	
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"updated": updated,
	})
}

func (h *Handler) DeleteNoteByID(c *gin.Context){
	idStr := c.Param("id")

	objID, err := bson.ObjectIDFromHex(idStr) // convert to mongodb object id type
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid ID",
		})
		return
	}

	deleted, err := h.repo.DeleteById(c.Request.Context(),objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to delete",
		})
		return 
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"Note Not found",
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
	})
}