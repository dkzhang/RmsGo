package client

import (
	"context"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"time"

	pb "github.com/dkzhang/RmsGo/ResourceSM/gRpcService/grpc"
	"google.golang.org/grpc"
)

type SchedulingClient struct {
	host    string
	port    string
	address string
}

func NewSchedulingClient(host string, port string) SchedulingClient {
	return SchedulingClient{
		host:    host,
		port:    port,
		address: fmt.Sprintf("%s:%s", host, port),
	}
}

func (sc SchedulingClient) SchedulingCGpu(projectID int, cgpuType int, nodesAfter []int64,
	ctrlID int, ctrlCN string) (allocInfo project.AllocInfo, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(sc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return project.AllocInfo{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////

	reply, err := c.SchedulingCGpu(ctx, &pb.SchedulingCGpuRequest{
		ProjectID:  int64(projectID),
		CgpuType:   int64(cgpuType),
		NodesAfter: nodesAfter,
		CtrlID:     int64(ctrlID),
		CtrlCN:     ctrlCN,
	})
	if err != nil {
		return project.AllocInfo{},
			fmt.Errorf("grpc call SchedulingCGpu error: %v", err)
	}

	allocInfo = project.AllocInfo{
		ProjectID:           projectID,
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
	}

	return allocInfo, nil
}

func (sc SchedulingClient) SchedulingStorage(projectID int,
	storageSizeAfter int, storageAllocInfoAfter string, ctrlID int, ctrlCN string) (allocInfo project.AllocInfo, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(sc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return project.AllocInfo{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.SchedulingStorage(ctx, &pb.SchedulingStorageRequest{
		ProjectID:             int64(projectID),
		StorageSizeAfter:      int64(storageSizeAfter),
		StorageAllocInfoAfter: storageAllocInfoAfter,
		CtrlID:                int64(ctrlID),
		CtrlCN:                ctrlCN,
	})
	if err != nil {
		return project.AllocInfo{},
			fmt.Errorf("grpc call SchedulingStorage error: %v", err)
	}

	allocInfo = project.AllocInfo{
		ProjectID:           projectID,
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
	}

	return allocInfo, nil
}

func (sc SchedulingClient) QueryCGpuTree(projectID int, cgpuType int, QueryType int) (jsonTree string, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(sc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "",
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.QueryCGpuTree(ctx, &pb.QueryTreeRequest{
		ProjectID: int64(projectID),
		CgpuType:  int64(cgpuType),
		QueryType: int64(QueryType),
	})

	if err != nil {
		return "", fmt.Errorf("grpc call QueryCGpuTree error: %v", err)
	}

	return reply.JsonTree, nil
}

func (sc SchedulingClient) QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(sc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return projectRes.ResInfo{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.QueryProjectRes(ctx, &pb.QueryProjectResRequest{
		ProjectID: int64(projectID),
	})

	if err != nil {
		return projectRes.ResInfo{}, fmt.Errorf("grpc call QueryProjectResRequest error: %v", err)
	}

	pr = projectRes.ResInfo{
		ProjectID:           int(reply.ProjectID),
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
		CpuNodesArray:       reply.CpuNodesArray,
		CpuNodesStr:         reply.CpuNodesStr,
		GpuNodesArray:       reply.GpuNodesArray,
		GpuNodesStr:         reply.GpuNodesStr,
		StorageAllocInfo:    reply.StorageAllocInfo,
	}
	pr.CreatedAt, err = time.Parse(time.RFC3339Nano, reply.CreatedAt)
	if err != nil {
		return pr, fmt.Errorf("time parse CreatedAt error: %v", err)
	}
	pr.UpdatedAt, err = time.Parse(time.RFC3339Nano, reply.UpdatedAt)
	if err != nil {
		return pr, fmt.Errorf("time parse UpdatedAt error: %v", err)
	}

	return pr, nil
}

func (sc SchedulingClient) QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(sc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return projectRes.ResInfoLite{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//////////////////////////////////////////////////////////////////////////////
	reply, err := c.QueryProjectResLite(ctx, &pb.QueryProjectResRequest{
		ProjectID: int64(projectID),
	})

	if err != nil {
		return projectRes.ResInfoLite{},
			fmt.Errorf("grpc call QueryProjectResLite error: %v", err)
	}

	prl = projectRes.ResInfoLite{
		ProjectID:           int(reply.ProjectID),
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
	}

	return prl, nil
}
