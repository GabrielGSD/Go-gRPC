package service

import (
	"context"

	"github.com/GabrielGSD/Go-gRPC/internal/database"
	"github.com/GabrielGSD/Go-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDb,
	}
}

// ctx -> contexto, dados de Header, autenticação, etc
// in -> entrada, request que estamos recebendo
func (c *CategoryService) Create(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) List(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{
		Categories: categoriesResponse,
	}, nil
}
