package server

import (
	"context"
	pb "github.com/dkzhang/RmsGo/ResourceSM/gRpcService/grpc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	"github.com/dkzhang/RmsGo/ResourceSM/resScheduling"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type SchedulingServer struct {
	pb.UnimplementedSchedulingServiceServer

	TheResScheduling resScheduling.ResScheduling
}

func (s *SchedulingServer) SchedulingCGpu(ctx context.Context, in *pb.SchedulingCGpuRequest) (*pb.SchedulingReply, error) {
	var err error
	switch in.CgpuType {
	case 1:
		//CPU
		_, err = s.TheResScheduling.SchedulingCPU(int(in.ProjectID), in.NodesAfter, int(in.CtrlID), in.CtrlCN)
		if err != nil {
			return &pb.SchedulingReply{},
				status.Errorf(codes.NotFound, "SchedulingCPU error: %v", err)
		}
	case 2:
		//GPU
		_, err = s.TheResScheduling.SchedulingGPU(int(in.ProjectID), in.NodesAfter, int(in.CtrlID), in.CtrlCN)
		if err != nil {
			return &pb.SchedulingReply{},
				status.Errorf(codes.NotFound, "SchedulingGPU error: %v", err)
		}
	default:
		// Unsupported type
		return &pb.SchedulingReply{},
			status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.CgpuType)
	}

	prl, err := s.TheResScheduling.QueryProjectResLiteByID(int(in.ProjectID))
	if err != nil {
		return &pb.SchedulingReply{},
			status.Errorf(codes.NotFound, "QueryProjectResByID error: %v", err)
	}
	return &pb.SchedulingReply{
		ProjectID:           int64(prl.ProjectID),
		CpuNodesAcquired:    int64(prl.CpuNodesAcquired),
		GpuNodesAcquired:    int64(prl.GpuNodesAcquired),
		StorageSizeAcquired: int64(prl.StorageSizeAcquired),
	}, nil
}

func (s *SchedulingServer) SchedulingStorage(ctx context.Context, in *pb.SchedulingStorageRequest) (*pb.SchedulingReply, error) {
	var err error
	_, err = s.TheResScheduling.SchedulingStorage(int(in.ProjectID), int(in.StorageSizeAfter), in.StorageAllocInfoAfter, int(in.CtrlID), in.CtrlCN)
	if err != nil {
		return &pb.SchedulingReply{},
			status.Errorf(codes.NotFound, "SchedulingStorage error: %v", err)
	}

	prl, err := s.TheResScheduling.QueryProjectResLiteByID(int(in.ProjectID))
	if err != nil {
		return &pb.SchedulingReply{},
			status.Errorf(codes.NotFound, "QueryProjectResByID error: %v", err)
	}
	return &pb.SchedulingReply{
		ProjectID:           int64(prl.ProjectID),
		CpuNodesAcquired:    int64(prl.CpuNodesAcquired),
		GpuNodesAcquired:    int64(prl.GpuNodesAcquired),
		StorageSizeAcquired: int64(prl.StorageSizeAcquired),
	}, nil
}

///////////////////////////////////////////////////////////////////////////////

func (s *SchedulingServer) QueryCGpuTree(ctx context.Context, in *pb.QueryTreeRequest) (*pb.QueryTreeReply, error) {
	var qTree *resGNodeTree.Tree
	var err error
	var selected []int64

	switch in.CgpuType {
	case typeCPU:
		// CPU
		switch in.QueryType {
		case typeOccupied:
			// Allocated
			qTree, selected, err = s.TheResScheduling.QueryCpuTreeAllocated(int(in.ProjectID), int(in.TreeFormat))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeAllocated error: %v", err)
			}
		case typeAvailable:
			// IdleAndAllocated
			qTree, selected, err = s.TheResScheduling.QueryCpuTreeIdleAndAllocated(int(in.ProjectID), int(in.TreeFormat))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeIdleAndAllocated error: %v", err)
			}
		case typeAll:
			// All
			qTree, selected, err = s.TheResScheduling.QueryCpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeAll error: %v", err)
			}
		default:
			return &pb.QueryTreeReply{},
				status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.QueryType)
		}
	case typeGPU:
		// GPU
		switch in.QueryType {
		case typeOccupied:
			// Allocated
			qTree, selected, err = s.TheResScheduling.QueryGpuTreeAllocated(int(in.ProjectID), int(in.TreeFormat))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeAllocated error: %v", err)
			}
		case typeAvailable:
			// IdleAndAllocated
			qTree, selected, err = s.TheResScheduling.QueryGpuTreeIdleAndAllocated(int(in.ProjectID), int(in.TreeFormat))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeIdleAndAllocated error: %v", err)
			}
		case typeAll:
			// All
			qTree, selected, err = s.TheResScheduling.QueryGpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeAll error: %v", err)
			}
		default:
			return &pb.QueryTreeReply{},
				status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.QueryType)
		}
	default:
		// Unsupported type
		return &pb.QueryTreeReply{},
			status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.CgpuType)
	}

	jsonTree, err := resGNodeTree.ToJsonForVue(*qTree)
	if err != nil {
		return &pb.QueryTreeReply{},
			status.Errorf(codes.NotFound, "resGNodeTree.ToJsonForVue error: %v", err)
	}
	return &pb.QueryTreeReply{
		JsonTree: jsonTree,
		NodesNum: int64(qTree.NodesNum),
		Selected: selected,
	}, nil
}

func (s *SchedulingServer) QueryProjectRes(ctx context.Context, in *pb.QueryProjectResRequest) (*pb.QueryProjectResReply, error) {
	pr, err := s.TheResScheduling.QueryProjectResByID(int(in.ProjectID))
	if err != nil {
		return &pb.QueryProjectResReply{},
			status.Errorf(codes.NotFound, "QueryProjectResByID error: %v", err)
	}
	return &pb.QueryProjectResReply{
		ProjectID:           int64(pr.ProjectID),
		CpuNodesAcquired:    int64(pr.CpuNodesAcquired),
		GpuNodesAcquired:    int64(pr.GpuNodesAcquired),
		StorageSizeAcquired: int64(pr.StorageSizeAcquired),
		CpuNodesArray:       pr.CpuNodesArray,
		CpuNodesStr:         pr.CpuNodesStr,
		GpuNodesArray:       pr.GpuNodesArray,
		GpuNodesStr:         pr.GpuNodesStr,
		StorageAllocInfo:    pr.StorageAllocInfo,
		CreatedAt:           pr.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:           pr.UpdatedAt.Format(time.RFC3339Nano),
	}, nil
}

func (s *SchedulingServer) QueryProjectResLite(ctx context.Context, in *pb.QueryProjectResRequest) (*pb.QueryProjectResLiteReply, error) {
	pr, err := s.TheResScheduling.QueryProjectResLiteByID(int(in.ProjectID))
	if err != nil {
		return &pb.QueryProjectResLiteReply{},
			status.Errorf(codes.NotFound, "QueryProjectResByID error: %v", err)
	}
	return &pb.QueryProjectResLiteReply{
		ProjectID:           int64(pr.ProjectID),
		CpuNodesAcquired:    int64(pr.CpuNodesAcquired),
		GpuNodesAcquired:    int64(pr.GpuNodesAcquired),
		StorageSizeAcquired: int64(pr.StorageSizeAcquired),
	}, nil
}

const (
	typeCPU = 1
	typeGPU = 2
)

const (
	typeOccupied  = 1
	typeAvailable = 2
	typeAll       = 3
)
