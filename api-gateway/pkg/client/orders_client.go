package client

import (
	"api-gateway/pkg/models"
	"context"

	pb "github.com/suyundykovv/AP1_ASS2/proto"
	"google.golang.org/grpc"
)

type OrderClient struct {
	conn   *grpc.ClientConn
	client pb.OrderServiceClient
}

func NewOrderClient(address string) (*OrderClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // Use WithTransportCredentials for production
	if err != nil {
		return nil, err
	}

	return &OrderClient{
		conn:   conn,
		client: pb.NewOrderServiceClient(conn),
	}, nil
}

func (c *OrderClient) CreateOrder(order *models.Order) (*models.Order, error) {
	pbReq := &pb.CreateOrderRequest{
		UserId: int32(order.UserID),
		Items:  order.Items,
		Status: order.Status,
		Total:  order.Total,
	}

	pbResp, err := c.client.CreateOrder(context.Background(), pbReq)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID:     pbResp.OrderId,
		UserID: int(pbResp.UserId),
		Items:  pbResp.Items,
		Status: pbResp.Status,
		Total:  pbResp.Total,
	}, nil
}

func (c *OrderClient) GetOrder(orderID int32) (*models.Order, error) {
	pbResp, err := c.client.GetOrder(context.Background(), &pb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID:     pbResp.OrderId,
		UserID: int(pbResp.UserId),
		Items:  pbResp.Items,
		Status: pbResp.Status,
		Total:  pbResp.Total,
	}, nil
}

func (c *OrderClient) UpdateOrder(order *models.Order) (*models.Order, error) {
	pbReq := &pb.UpdateOrderRequest{
		OrderId: order.ID,
		UserId:  int32(order.UserID),
		Items:   order.Items,
		Status:  order.Status,
		Total:   order.Total,
	}

	pbResp, err := c.client.UpdateOrder(context.Background(), pbReq)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID:     pbResp.OrderId,
		UserID: int(pbResp.UserId),
		Items:  pbResp.Items,
		Status: pbResp.Status,
		Total:  pbResp.Total,
	}, nil
}

func (c *OrderClient) DeleteOrder(orderID int32) error {
	_, err := c.client.DeleteOrder(context.Background(), &pb.DeleteOrderRequest{OrderId: orderID})
	return err
}

func (c *OrderClient) ListOrders(userID int) ([]*models.Order, error) {
	// Implement if your proto has ListOrders RPC, or simulate with multiple GetOrder calls
	return nil, nil
}

func (c *OrderClient) Close() error {
	return c.conn.Close()
}
