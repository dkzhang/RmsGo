package client

import (
	"context"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/project"
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

	reply, err := c.SchedulingCGpu(ctx, &pb.SchedulingCGpuRequest{
		ProjectID:  int64(projectID),
		CgpuType:   int64(cgpuType),
		NodesAfter: nodesAfter,
		CtrlID:     int64(ctrlID),
		CtrlCN:     ctrlCN,
	})
	if err != nil {
		return fmt.Errorf("grpc call SchedulingCGpu error: %v", err)
	}

	// Update project info by reply
	if reply.IsFirstAlloc {
		// update project status info
		rsc.pdm.UpdateStatusInfo(project.StatusInfo{
			ProjectID:   projectID,
			BasicStatus: project.BasicStatusRunning,
		})
	}

	rsc.pdm.UpdateAllocInfo(project.AllocInfo{
		ProjectID:           projectID,
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
	})

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
	reply, err := c.SchedulingStorage(ctx, &pb.SchedulingStorageRequest{
		ProjectID:             int64(projectID),
		StorageSizeAfter:      int64(storageSizeAfter),
		StorageAllocInfoAfter: storageAllocInfoAfter,
		CtrlID:                int64(ctrlID),
		CtrlCN:                ctrlCN,
	})
	if err != nil {
		return fmt.Errorf("grpc call SchedulingStorage error: %v", err)
	}

	// Update project info by reply
	if reply.IsFirstAlloc {
		// update project status info
		rsc.pdm.UpdateStatusInfo(project.StatusInfo{
			ProjectID:   projectID,
			BasicStatus: project.BasicStatusRunning,
		})
	}

	rsc.pdm.UpdateAllocInfo(project.AllocInfo{
		ProjectID:           projectID,
		CpuNodesAcquired:    int(reply.CpuNodesAcquired),
		GpuNodesAcquired:    int(reply.GpuNodesAcquired),
		StorageSizeAcquired: int(reply.StorageSizeAcquired),
	})
	return nil
}

func (rsc ResSchedulingClient) QueryCGpuTree(projectID int, cgpuType int, QueryType int) (jsonTree string, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rsc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
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

func (rsc ResSchedulingClient) QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rsc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return projectRes.ResInfo{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
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

func (rsc ResSchedulingClient) QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error) {
	//////////////////////////////////////////////////////////////////////////////
	// Common Operation
	conn, err := grpc.Dial(rsc.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return projectRes.ResInfoLite{},
			fmt.Errorf("fatal error! grpc.Dial cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSchedulingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
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
