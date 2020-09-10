package client

import (
	"context"
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"time"

	pb "github.com/dkzhang/RmsGo/ResourceSM/resScheduling/grpc"
	"google.golang.org/grpc"
)

type ResSchedulingClient struct {
	host    string
	port    int
	address string
	pdm     projectDM.ProjectDM
}

func NewResSchedulingClient(host string, port int, pdm projectDM.ProjectDM) ResSchedulingClient {
	return ResSchedulingClient{
		host:    host,
		port:    port,
		address: fmt.Sprintf("%s:%d", host, port),
		pdm:     pdm,
	}
}

func (rsc ResSchedulingClient) SchedulingCGpu(projectID int, cgpuType int, nodesAfter []int64,
	ctrlID int, ctrlCN string) (err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rsc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////

	_, err = c.SchedulingCGpu(ctx, &pb.SchedulingCGpuRequest{
		ProjectID:  int64(projectID),
		CgpuType:   int64(cgpuType),
		NodesAfter: nodesAfter,
		CtrlID:     int64(ctrlID),
		CtrlCN:     ctrlCN,
	})
	if err != nil {
		return fmt.Errorf("grpc call SchedulingCGpu error: %v", err)
	}
	return nil
}

func (rsc ResSchedulingClient) SchedulingStorage(projectID int,
	storageSizeAfter int, storageAllocInfoAfter string, ctrlID int, ctrlCN string) (err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rsc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
}
