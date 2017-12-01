package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v2/app/models"
)

// ArticlesController is
type ArticlesController struct {
	Controller
}

// NewArticlesController is
func NewArticlesController() ArticlesController {
	return ArticlesController{}
}

// GetAll is get all articles
func (ctl ArticlesController) GetAll(c echo.Context) error {

	var (
		articles    models.ArticlePage
		queryparams models.QueryParams
		err         error
	)

	qps := c.QueryParams()

	limitq := c.QueryParam("limit")
	offsetq := c.QueryParam("offset")
	startTimeq := c.QueryParam("start_time")
	endTime := c.QueryParam("end_time")

	limit, _ := strconv.Atoi(limitq)
	offset, _ := strconv.Atoi(offsetq)
	fmt.Println(qps)
	fmt.Println(limit)
	fmt.Println(offset)

	if limit == 0 {
		limit = 10
	}

	queryparams.Limit = limit
	queryparams.Offset = offset
	queryparams.StartTime = startTimeq
	queryparams.EndTime = endTime

	queryparams2 := c.Get("queryparams")

	fmt.Println("queryparams2")
	fmt.Println(queryparams2)

	articles, err = models.GetArticles(queryparams)
	if err != nil {
		return ctl.ErrorResponse(c, http.StatusForbidden, err.Error())
	}

	return ctl.SuccessResponse(c, articles)
}

// Get is get one article
func (ctl ArticlesController) Get(c echo.Context) error {
	var (
		article models.Article
		err     error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	article, err = models.GetArticleById(id)
	if err != nil {
		return ctl.ErrorResponse(c, http.StatusForbidden, err.Error())
	}
	return ctl.SuccessResponse(c, article)
}

// Create is create article
func (ctl ArticlesController) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))
	article := new(models.Article)

	article.UserID = userID
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	err := models.CreateArticle(article)

	if err != nil {
		return ctl.ErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
	}

	return ctl.SuccessResponse(c, article)
}

// Update is update article
func (ctl ArticlesController) Update(c echo.Context) error {
	// Parse the content
	article := new(models.Article)

	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update article data
	err = m.UpdateArticle(article)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

// Delete is delete article
func (ctl ArticlesController) Delete(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteArticle()
	return c.JSON(http.StatusNoContent, err)
}
