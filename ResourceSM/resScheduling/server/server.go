package main

import (
	"context"
	"fmt"
	pb "github.com/dkzhang/RmsGo/ResourceSM/resScheduling/grpc"
	"github.com/dkzhang/RmsGo/ResourceSM/resScheduling/resourceScheduling"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

const (
	port = ":50071"
)

type server struct {
	pb.UnimplementedSchedulingServer

	TheResScheduling resourceScheduling.ResScheduling
}

func (s *server) SchedulingCGpu(ctx context.Context, in *pb.SchedulingCGpuRequest) (*pb.SchedulingReply, error) {
	switch in.CgpuType {
	case 1:
		//CPU
		err := s.TheResScheduling.SchedulingCPU(int(in.ProjectID), in.NodesAfter, int(in.CtrlID), in.CtrlCN)
		if err != nil {
			return &pb.SchedulingReply{},
				status.Errorf(codes.NotFound, "SchedulingCPU error: %v", err)
		}
	case 2:
		//GPU
		err := s.TheResScheduling.SchedulingGPU(int(in.ProjectID), in.NodesAfter, int(in.CtrlID), in.CtrlCN)
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

func (s *server) SchedulingStorage(ctx context.Context, in *pb.SchedulingStorageRequest) (*pb.SchedulingReply, error) {
	err := s.TheResScheduling.SchedulingStorage(int(in.ProjectID), int(in.StorageSizeAfter), in.StorageAllocInfoAfter, int(in.CtrlID), in.CtrlCN)
	if err != nil {
		return &pb.SchedulingReply{},
			status.Errorf(codes.NotFound, "SchedulingStorage error: %v", err)
	}
	return &pb.SchedulingReply{}, nil
}

///////////////////////////////////////////////////////////////////////////////

func (s *server) QueryCGpuTree(ctx context.Context, in *pb.QueryTreeRequest) (*pb.QueryTreeReply, error) {
	switch in.CgpuType {
	case 1:
		// CPU
		switch in.QueryType {
		case 1:
			// Allocated
			jsonTree, err := s.TheResScheduling.QueryCpuTreeAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeAllocated error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 2:
			// IdleAndAllocated
			jsonTree, err := s.TheResScheduling.QueryCpuTreeIdleAndAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeIdleAndAllocated error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 3:
			// All
			jsonTree, err := s.TheResScheduling.QueryCpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryCpuTreeAll error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		default:
			return &pb.QueryTreeReply{},
				status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.QueryType)
		}
	case 2:
		// GPU
		switch in.QueryType {
		case 1:
			// Allocated
			jsonTree, err := s.TheResScheduling.QueryGpuTreeAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeAllocated error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 2:
			// IdleAndAllocated
			jsonTree, err := s.TheResScheduling.QueryGpuTreeIdleAndAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeIdleAndAllocated error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 3:
			// All
			jsonTree, err := s.TheResScheduling.QueryGpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{},
					status.Errorf(codes.NotFound, "QueryGpuTreeAll error: %v", err)
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		default:
			return &pb.QueryTreeReply{},
				status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.QueryType)
		}
	default:
		// Unsupported type
		return &pb.QueryTreeReply{},
			status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.CgpuType)
	}
}

func (s *server) QueryProjectRes(ctx context.Context, in *pb.QueryProjectResRequest) (*pb.QueryProjectResReply, error) {
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

func (s *server) QueryProjectResLite(ctx context.Context, in *pb.QueryProjectResRequest) (*pb.QueryProjectResReplyLite, error) {
	pr, err := s.TheResScheduling.QueryProjectResLiteByID(int(in.ProjectID))
	if err != nil {
		return &pb.QueryProjectResReplyLite{},
			status.Errorf(codes.NotFound, "QueryProjectResByID error: %v", err)
	}
	return &pb.QueryProjectResReplyLite{
		ProjectID:           int64(pr.ProjectID),
		CpuNodesAcquired:    int64(pr.CpuNodesAcquired),
		GpuNodesAcquired:    int64(pr.GpuNodesAcquired),
		StorageSizeAcquired: int64(pr.StorageSizeAcquired),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulingServer(s, &server{})
	fmt.Printf("Begin to serve %v \n", time.Now())
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
