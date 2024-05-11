package grpc

import (
	"context"

	"github.com/RizqiSugiarto/GaleryService/internal/entity"
	"github.com/RizqiSugiarto/GaleryService/internal/usecase"
	proto "github.com/RizqiSugiarto/GaleryService/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GalleryGrpcHandler struct {
	galleryService usecase.Gallery
	proto.UnimplementedGaleryServiceServer
}

func NewGrpcService(grpcServer *grpc.Server, galleryService usecase.Gallery) {
	grpcHandler := &GalleryGrpcHandler{
		galleryService: galleryService,
	}
	proto.RegisterGaleryServiceServer(grpcServer, grpcHandler)
}

func(g *GalleryGrpcHandler) SaveLink(ctx context.Context, req *proto.SaveLinkRequest) (*proto.CommonResponse, error) {
	userData := entity.Gallery{
		Link: req.Link,
		UserId: req.UserId,
	}
	
	err := g.galleryService.CreateLink(ctx, userData)

	if err != nil {
		return &proto.CommonResponse{}, err
	}

	res := &proto.CommonResponse{
		Message: "Save Link successfuly",
	}

	return res, nil
}

func(g *GalleryGrpcHandler) UpdateLink(ctx context.Context, req *proto.UpdateLinkRequest) (*proto.CommonResponse, error) {
	userData := entity.Gallery{
		Link: req.Galery.Link,
		UserId: req.Galery.UserId,
	}

	err := g.galleryService.UpdateLink(ctx, req.UserId, userData)

	if err != nil {
		return &proto.CommonResponse{}, err
	}

	res := &proto.CommonResponse{
		Message: "Update Link successfuly",
	}
	return res, nil
}

func(g *GalleryGrpcHandler) GetLinkByUserId(ctx context.Context, req *proto.GetLinkByUserIdRequest) (*proto.GetLinkByUserIdResponse, error) {
	userData, err := g.galleryService.GetLinkByUserId(ctx, req.UserId)

	if err != nil {
		return &proto.GetLinkByUserIdResponse{}, err
	}

	res := &proto.GetLinkByUserIdResponse{
		Galery: &proto.Galery{
			Id: userData.Id,
			Link: userData.Link,
			UserId: userData.UserId,
			CreatedAt: timestamppb.New(userData.Created_at),
			UpdatedAt: timestamppb.New(userData.Updated_at),
		},
	}

	return res, nil
}

func(g *GalleryGrpcHandler) DeleteLink(ctx context.Context, req *proto.DeleteLinkRequest) (*proto.CommonResponse, error) {
	err := g.galleryService.DeleteLink(ctx, req.UserId)

	if err != nil {
		return &proto.CommonResponse{}, err
	}

	res := &proto.CommonResponse{
		Message: "Delete Successfuly",
	}

	return res, nil
}