package v1

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/romanchechyotkin/avito_test_task/internal/controller/v1/middleware"
	"github.com/romanchechyotkin/avito_test_task/internal/controller/v1/request"
	"github.com/romanchechyotkin/avito_test_task/internal/controller/v1/response"
	"github.com/romanchechyotkin/avito_test_task/internal/service"
	"github.com/romanchechyotkin/avito_test_task/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type flatRoutes struct {
	log *slog.Logger

	flatService service.Flat
}

func newFlatRoutes(log *slog.Logger, g *gin.RouterGroup, flatService service.Flat, authMiddleware *middleware.AuthMiddleware) {
	log = log.With(slog.String("component", "flat routes"))

	r := &flatRoutes{
		log:         log,
		flatService: flatService,
	}

	g.POST("/create", authMiddleware.AuthOnly(), r.createFlat)
	g.PATCH("/update", authMiddleware.ModeratorsOnly(), r.updateFlat)
}

// @Summary Create Flat
// @Description Create Flat
// @Tags flat
// @Accept json
// @Produce json
// @Param input body request.CreateFlat true "input"
// @Success 201 {object} response.Flat
// @Security JWT
// @Router /v1/flat/create [post]
func (r *flatRoutes) createFlat(c *gin.Context) {
	var req request.CreateFlat

	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error("failed to read request data", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := validator.New().Struct(req); err != nil {
		r.log.Error("failed to validate request data", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	flat, err := r.flatService.CreateFlat(c, &service.FlatCreateInput{
		Number:      req.Number,
		HouseID:     req.HouseID,
		Price:       req.Price,
		RoomsAmount: req.RoomsAmount,
	})
	if err != nil {
		if errors.Is(err, service.ErrHouseNotFound) || errors.Is(err, service.ErrFlatExists) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		r.log.Error("failed to create flat in database", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, response.BuildFlat(flat))
}

// @Summary Update Flat
// @Description Update Flat
// @Tags flat
// @Accept json
// @Produce json
// @Param input body request.UpdateFlat true "input"
// @Success 200 {object} response.Flat
// @Security JWT
// @Router /v1/flat/update [post]
func (r *flatRoutes) updateFlat(c *gin.Context) {
	var req request.UpdateFlat

	userID, ok := c.Get("userID")
	if !ok {
		r.log.Error("failed to get key from context", slog.String("key", "userType"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get key from context",
		})

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error("failed to read request data", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := validator.New().Struct(req); err != nil {
		r.log.Error("failed to validate request data", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	flat, err := r.flatService.UpdateFlat(c, &service.FlatUpdateInput{
		ID:          req.ID,
		Status:      req.Status,
		ModeratorID: userID.(string),
	})
	if err != nil {
		if errors.Is(err, service.ErrFlatNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		r.log.Error("failed to update flat status", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.BuildFlat(flat))
}
