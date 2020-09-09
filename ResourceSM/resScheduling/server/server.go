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
			return &pb.SchedulingReply{
				ErrorMessage: fmt.Sprintf("SchedulingCPU error: %v", err),
			}, nil
		}
		return &pb.SchedulingReply{}, nil
	case 2:
		//GPU
		err := s.TheResScheduling.SchedulingGPU(int(in.ProjectID), in.NodesAfter, int(in.CtrlID), in.CtrlCN)
		if err != nil {
			return &pb.SchedulingReply{
				ErrorMessage: fmt.Sprintf("SchedulingGPU error: %v", err),
			}, nil
		}
		return &pb.SchedulingReply{}, nil
	default:
		// Unsupported type
		return &pb.SchedulingReply{},
			status.Errorf(codes.InvalidArgument, " Unsupported type: %d", in.CgpuType)
	}
}

func (s *server) SchedulingStorage(ctx context.Context, in *pb.SchedulingStorageRequest) (*pb.SchedulingReply, error) {
	err := s.TheResScheduling.SchedulingStorage(int(in.ProjectID), int(in.StorageSizeAfter), in.StorageAllocInfoAfter, int(in.CtrlID), in.CtrlCN)
	if err != nil {
		return &pb.SchedulingReply{
			ErrorMessage: fmt.Sprintf("SchedulingStorage error: %v", err),
		}, nil
	}
	return &pb.SchedulingReply{ErrorMessage: ""}, nil
}
func (s *server) QueryCGpuTree(ctx context.Context, in *pb.QueryTreeRequest) (*pb.QueryTreeReply, error) {
	switch in.CgpuType {
	case 1:
		// CPU
		switch in.QueryType {
		case 1:
			// Allocated
			jsonTree, err := s.TheResScheduling.QueryCpuTreeAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryCpuTreeAllocated error: %v", err),
				}, nil
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 2:
			// IdleAndAllocated
			jsonTree, err := s.TheResScheduling.QueryCpuTreeIdleAndAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryCpuTreeIdleAndAllocated error: %v", err),
				}, nil
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 3:
			// All
			jsonTree, err := s.TheResScheduling.QueryCpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryCpuTreeAll error: %v", err),
				}, nil
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
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryGpuTreeAllocated error: %v", err),
				}, nil
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 2:
			// IdleAndAllocated
			jsonTree, err := s.TheResScheduling.QueryGpuTreeIdleAndAllocated(int(in.ProjectID))
			if err != nil {
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryGpuTreeIdleAndAllocated error: %v", err),
				}, nil
			}
			return &pb.QueryTreeReply{JsonTree: jsonTree}, nil
		case 3:
			// All
			jsonTree, err := s.TheResScheduling.QueryGpuTreeAll()
			if err != nil {
				return &pb.QueryTreeReply{
					JsonTree:     "",
					ErrorMessage: fmt.Sprintf("QueryGpuTreeAll error: %v", err),
				}, nil
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
func (s *server) QueryStorage(ctx context.Context, in *pb.QueryStorageRequest) (*pb.QueryStorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStorage not implemented")
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
