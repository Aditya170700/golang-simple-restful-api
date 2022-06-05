package helper

import (
	"github.com/Aditya170700/gorestfulapi/model/domain"
	"github.com/Aditya170700/gorestfulapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponses []web.CategoryResponse
	for _, category := range categories {
		CategoryResponses = append(CategoryResponses, ToCategoryResponse(category))
	}

	return CategoryResponses
}
