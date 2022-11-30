package v1

import (
	"net/http"
	"strconv"

	"github.com/MuhammadyusufAdhamov/note/api/models"
	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /note [post]
// @Summary Create a note
// @Description Create a note
// @Tags note
// @Accept json
// @Produce json
// @Param note body models.CreateNoteRequest true "Note"
// @Success 201 {object} models.Note
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateNote(c *gin.Context) {
	var (
		req models.CreateNoteRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Note().CreateNote(&repo.Note{
		UserId: req.UserId,
		Title: req.Title,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, parseNoteModel(resp))
}

func parseNoteModel(note *repo.Note) models.Note {
	return models.Note{
		ID: note.ID,
		UserId: note.UserId,
		Title: note.Title,
		Description: note.Description,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		DeletedAt: note.DeletedAt,
	}
}

// @Router /note/{id} [get]
// @Summary Get note by id
// @Description Get note by id
// @Tags note
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.Note
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Note().GetNote(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseNoteModel(resp))
}

// @Router /note/{id} [put]
// @Summary Update a note
// @Description Update a note
// @Tags note
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param note body models.CreateNoteRequest true "Note"
// @Success 201 {object} models.Note
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateNote(c *gin.Context) {
	var result models.Note

	err := c.ShouldBindJSON(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	result.ID = id

	note, err := h.storage.Note().UpdateNote((*repo.Note)(&result))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, note)
}

// @Router /note/{id} [delete]
// @Summary delete note by id
// @Description Delete note by id
// @Tags note
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.Note
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Note().DeleteNote(&repo.Note{
		ID: int64(id),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseNoteModel(resp))
}

// @Router /notes [get]
// @Summary Get all notes
// @Description Get all notes
// @Tags note
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 201 {object} models.GetAllNotesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllnotes(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Note().GetAllNotes(&repo.GetAllNotesParams{
		Page: req.Page,
		Limit: req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getNotesResponse(result))
}

func getNotesResponse(data *repo.GetAllNotesResults) *models.GetAllNotesResponse {
	response := models.GetAllNotesResponse{
		Notes: make([]*models.Note, 0),
		Count: data.Count,
	}

	for _, note := range data.Notes {
		u := parseNoteModel(note)
		response.Notes = append(response.Notes, &u)
	}

	return &response
}