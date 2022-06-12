package helper

import (
	"sudutkampus/gorestfulapi/model/domain"
	"sudutkampus/gorestfulapi/model/web"
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
