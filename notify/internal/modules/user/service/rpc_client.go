package service

import (
	"context"
	"fmt"
	"projects/LDmitryLD/hugoproxy-microservices/notify/config"
	"projects/LDmitryLD/hugoproxy-microservices/notify/internal/models"

	pb "github.com/LDmitryLD/protos/gen/usergrpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Userer interface {
	GetByID(ctx context.Context, id int) (models.User, error)
}

type UserGRPCClient struct {
	client pb.UsererClient
}

func NewUserGRPCClient(conf config.UserRPC) (Userer, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.Host, conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewUsererClient(conn)
	return &UserGRPCClient{client: client}, nil
}

func (u UserGRPCClient) GetByID(ctx context.Context, id int) (models.User, error) {
	res, err := u.client.GetByID(ctx, &pb.GetByIDRequest{Id: uint32(id)})
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email: res.GetEmail(),
		Phone: res.GetPhone(),
	}

	return user, nil
}
