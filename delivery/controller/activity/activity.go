package activity

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"todo/delivery/common"
	"todo/entity"
	"todo/repository/activity"
)

type ActivityController struct {
	ActivityRepo activity.ActivityRepository
}

func NewActivityController(activitytRepo activity.ActivityRepository) *ActivityController {
	return &ActivityController{activitytRepo}
}

func (ac ActivityController) CreateActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &CreateActivityRequest{}

		c.Bind(&data)

		if data.Title == "" {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Bad Request", "title cannot be null"))
		}

		var activity = &entity.Activity{Title: data.Title, Email: data.Email}

		result, err := ac.ActivityRepo.CreateActivity(activity)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Create Activity", err.Error()))
		}

		response := &DataResponse{
			Id:        result.ID,
			Title:     result.Title,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, common.SuccessResponse(response))
	}
}

func (ac ActivityController) GetActivityById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Id must be a number", err.Error()))
		}

		result, err := ac.ActivityRepo.GetActivityByID(int64(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.ActivityNotFoundResponse(uint(id)))
		}

		response := &DataResponse{
			Id:        result.ID,
			Title:     result.Title,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac ActivityController) GetAllActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := ac.ActivityRepo.GetAllActivity()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Get All Activity", err.Error()))
		}

		var response = make([]*DataResponse, 0)
		for _, v := range result {
			response = append(response, &DataResponse{
				Id:        v.ID,
				Title:     v.Title,
				Email:     v.Email,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac ActivityController) UpdateActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &UpdateActivityRequest{}

		c.Bind(&data)

		if data.Title == "" {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Bad Request", "title cannot be null"))
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.ActivityNotFoundResponse(uint(id)))
		}

		var activity = &entity.Activity{Title: data.Title}
		activity.ID = uint(id)
		activity.UpdatedAt = time.Now()

		result, err := ac.ActivityRepo.UpdateActivity(activity)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, common.ActivityNotFoundResponse(uint(id)))
			}
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Update Activity", err.Error()))
		}

		response := &DataResponse{
			Id:        result.ID,
			Title:     result.Title,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac ActivityController) DeleteActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Id must be a number", err.Error()))
		}

		err = ac.ActivityRepo.DeleteActivity(int64(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, common.ActivityNotFoundResponse(uint(id)))
			}
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Delete Activity", err.Error()))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(make(map[string]interface{})))
	}
}
