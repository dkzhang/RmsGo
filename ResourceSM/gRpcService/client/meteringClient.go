package client

import (
	"context"
	"fmt"
	pb "github.com/dkzhang/RmsGo/ResourceSM/gRpcService/grpc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"google.golang.org/grpc"
	"time"
)

type MeteringClient struct {
	host    string
	port    int
	address string
}

func NewMeteringClient(host string, port int) MeteringClient {
	return MeteringClient{
		host:    host,
		port:    port,
		address: fmt.Sprintf("%s:%d", host, port),
	}
}

func (mc MeteringClient) QueryMetering(projectID int,
	meteringType int, typeInfo string) (ms metering.Statement, err error) {
	/////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(mc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return metering.Statement{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMeteringServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.QueryMetering(ctx, &pb.QmRequest{
		ProjectID:    int64(projectID),
		MeteringType: int64(meteringType),
		TypeInfo:     typeInfo,
	})
	if err != nil {
		return metering.Statement{},
			fmt.Errorf("grpc call SchedulingStorage error: %v", err)
	}

	return QmReplyToStatement(reply), nil
}

func (mc MeteringClient) ComputeMetering(projectID int,
	meteringType int, typeInfo string) (ms metering.Statement, err error) {
	/////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(mc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return metering.Statement{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMeteringServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.ComputeMetering(ctx, &pb.QmRequest{
		ProjectID:    int64(projectID),
		MeteringType: int64(meteringType),
		TypeInfo:     typeInfo,
	})
	if err != nil {
		return metering.Statement{},
			fmt.Errorf("grpc call SchedulingStorage error: %v", err)
	}

	return QmReplyToStatement(reply), nil
}

func (mc MeteringClient) RegisterCronMeteringTask(meteringType int) (msg string, err error) {
	/////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(mc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "",
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMeteringServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.RegisterCronMeteringTask(ctx, &pb.RegCmtRequest{
		MeteringType: int64(meteringType),
	})
	if err != nil {
		return "",
			fmt.Errorf("grpc call SchedulingStorage error: %v", err)
	}

	return reply.Msg, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func QmReplyToStatement(reply *pb.QmReply) (ms metering.Statement) {
	ms = metering.Statement{
		MeteringID:           0,
		ProjectID:            int(reply.ProjectID),
		MeteringType:         int(reply.MeteringType),
		MeteringTypeInfo:     reply.MeteringTypeInfo,
		CpuAmountInDays:      int(reply.CpuAmountInDays),
		GpuAmountInDays:      int(reply.GpuAmountInDays),
		StorageAmountInDays:  int(reply.StorageAmountInDays),
		CpuAmountInHours:     int(reply.CpuAmountInHours),
		GpuAmountInHours:     int(reply.GpuAmountInHours),
		StorageAmountInHours: int(reply.StorageAmountInHours),
		CpuNodeMeteringJson:  reply.CpuNodeMeteringJson,
		GpuNodeMeteringJson:  reply.GpuNodeMeteringJson,
		StorageMeteringJson:  reply.StorageMeteringJson,
	}

	// ignore time.Parse error
	ms.CreatedAt, _ = time.Parse(time.RFC3339Nano, reply.CreatedAt)

	return ms
}
