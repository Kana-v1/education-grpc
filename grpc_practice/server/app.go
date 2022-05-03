package server

import (
	context "context"
	"errors"
	"fmt"
	"io"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func Start() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("certificates/localhost.crt", "certificates/localhost.key")
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)

	RegisterEmployeeServiceServer(s, new(employeeService))

	fmt.Println("start listening on port " + port)
	s.Serve(listener)
}

type employeeService struct{}

func (es *employeeService) GetByBadgeNumber(ctx context.Context,
	req *GetByBadgeNumberRequest) (*EmployeeResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("metadata received: %v\n", md)
	}

	for _, e := range employees {
		if e.BadgeNumber == req.BadgeNumber {
			return &EmployeeResponse{Employee: &e}, nil
		}
	}

	return nil, errors.New("employee not found")
}

func (es *employeeService) GetAll(req *GetAllRequest, stream EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&EmployeeResponse{Employee: &e})
	}

	return nil
}

func (es *employeeService) Save(ctx context.Context, stream *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, nil
}

func (es *employeeService) SaveAll(stream EmployeeService_SaveAllServer) error {
	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		employees = append(employees, *emp.Employee)
		stream.Send(&EmployeeResponse{Employee: emp.Employee})
	}

	for _, e := range employees {
		fmt.Println(e)
	}

	return nil
}

func (es *employeeService) AddPhoto(stream EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("receiving photo from badge number %v\n", md["badge_number"][0])
	}

	imgData := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("file received with length: %v\n", len(imgData))
			return stream.SendAndClose(&AddPhotoResponse{IsOK: true})
		}

		if err != nil {
			return err
		}

		fmt.Printf("Received %v bytes\n", len(data.Data))
		imgData = append(imgData, data.Data...)
	}
}

func (es *employeeService) mustEmbedUnimplementedEmployeeServiceServer() {}
