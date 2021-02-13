package controllers

import (
	"net/http"

	"github.com/End-S/spooky_kingdom/controllers/requests"
	"github.com/End-S/spooky_kingdom/controllers/responses"
	"github.com/End-S/spooky_kingdom/models"
	"github.com/End-S/spooky_kingdom/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ArticleController struct
type ArticleController struct {
	articleModel *models.ArticleModel
}

// NewActicleController creates a new article controller instance
func NewActicleController(am *models.ArticleModel) *ArticleController {
	return &ArticleController{
		articleModel: am,
	}
}

// GetArticles function handles a request for articles
func (ac *ArticleController) GetArticles(c echo.Context) error {
	r := new(requests.GetArticlesReq)

	// bind query params to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ValidationError(err, requests.GetArticlesReq{}))
	}

	// if user want to retrieve articles that are pending review check they are admin
	if r.AssessPending == "true" {
		if err := utils.VerifyAuthHeader(c, "admin"); err != nil {
			return c.JSON(http.StatusForbidden, responses.NewErrorResponse("Not permitted for this resource"))
		}
	} else {
		r.AssessPending = "false"
	}

	articles, count, err := ac.articleModel.List(r, r.AssessPending == "true")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to retrieve articles"))
	}

	return c.JSON(http.StatusOK, responses.NewListArticlesResponse(articles, count))
}

// UpdateArticle function handles the updation of an article
func (ac *ArticleController) UpdateArticle(c echo.Context) error {
	r := new(requests.UpdateArticleReq)

	// bind json to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ValidationError(err, requests.UpdateArticleReq{}))
	}

	article, err := ac.articleModel.Update(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to update article"))
	}

	return c.JSON(http.StatusOK, responses.NewUpdateArticlesResponse(article))
}

// DeleteArticle function handles a delete request for an article
func (ac *ArticleController) DeleteArticle(c echo.Context) error {
	// extract param from url
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			responses.NewErrorResponse("Unrecognised UUID"))
	}

	count, err := ac.articleModel.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			responses.NewErrorResponse("Server error, unable to delete article"))
	}

	if count <= 0 {
		return c.JSON(http.StatusBadRequest,
			responses.NewErrorResponse("Article does not exist"))
	}

	return c.JSON(http.StatusOK, responses.NewDeleteResponse(true))
}

// ADMIN POST article changes

// ADMIN POST new article (actioned by web crawler)

// To impove pagination we should have indexes for fields, which we are using in ORDER BY
// index date pub, title ASC/DSC and type