package service

import (
	"context"
	"github.com/kotayz/grpc/internal/database"
	"github.com/kotayz/grpc/internal/pb"
)

type CategoryService {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}