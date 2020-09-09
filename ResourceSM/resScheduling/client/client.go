package client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/dkzhang/RmsGo/ResourceSM/resScheduling/grpc"
	"google.golang.org/grpc"
)

type ResScheduling struct {
	host    string
	port    int
	address string
}

func NewResScheduling(host string, port int) ResScheduling {
	return ResScheduling{
		host:    host,
		port:    port,
		address: fmt.Sprintf("%s:%d", host, port),
	}
}

func (rs ResScheduling) SchedulingCGpu(projectID int, cgpuType int, nodesAfter []int64,
	ctrlID int, ctrlCN string) (err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rs.address, grpc.WithInsecure(), grpc.WithBlock())
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
