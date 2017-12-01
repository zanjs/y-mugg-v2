package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v2/app/models"
)

// AllWarerooms is get all warerooms
func AllWarerooms(c echo.Context) error {
	var (
		warerooms []models.Wareroom
		err       error
	)
	warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, warerooms)
}

// ShowWareroom is get one wareroom
func ShowWareroom(c echo.Context) error {
	var (
		wareroom models.Wareroom
		err      error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	wareroom, err = models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, wareroom)
}

// CreateWareroom wareroom
func CreateWareroom(c echo.Context) error {
	wareroom := new(models.Wareroom)
	wareroom.Title = c.FormValue("title")
	wareroom.Numbering = c.FormValue("numbering")

	err := models.CreateWareroom(wareroom)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, wareroom)
}

// UpdateWareroom is update wareroom
func UpdateWareroom(c echo.Context) error {
	// Parse the content
	wareroom := new(models.Wareroom)

	wareroom.Title = c.FormValue("title")
	wareroom.Numbering = c.FormValue("numbering")

	sortV := c.FormValue("sort")
	sort, _ := strconv.Atoi(sortV)

	wareroom.Sort = sort

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update wareroom data
	err = m.UpdateWareroom(wareroom)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

//delete wareroom
func DeleteWareroom(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetWareroomById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteWareroom()
	return c.JSON(http.StatusNoContent, err)
}
