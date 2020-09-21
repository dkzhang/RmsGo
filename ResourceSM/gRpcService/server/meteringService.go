package main

import (
	"context"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDM"
	pb "github.com/dkzhang/RmsGo/ResourceSM/gRpcService/grpc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type MeteringServer struct {
	pb.UnimplementedSchedulingServiceServer

	TheMeteringDM meteringDM.MeteringDM
}

func (s *MeteringServer) QueryMetering(ctx context.Context, in *pb.QmRequest) (*pb.QmReply, error) {
	ms, err := s.TheMeteringDM.QueryWithCI(int(in.ProjectID), int(in.MeteringType), in.TypeInfo)
	if err != nil {
		return &pb.QmReply{},
			status.Errorf(codes.NotFound, "MeteringDM.QueryWithCI error: %v", err)
	}

	return statementToQmReply(ms), nil
}

func (s *MeteringServer) ComputeMetering(ctx context.Context, in *pb.QmRequest) (*pb.QmReply, error) {
	switch in.MeteringType {
	case metering.TypeMonthly:
		return &pb.QmReply{},
			status.Errorf(codes.Unimplemented, "metering Monthly feature is not implemented yet")
	case metering.TypeQuarterly:
		return &pb.QmReply{},
			status.Errorf(codes.Unimplemented, "metering Quarterly feature is not implemented yet")
	case metering.TypeAnnual:
		return &pb.QmReply{},
			status.Errorf(codes.Unimplemented, "metering Quarterly feature is not implemented yet")
	case metering.TypeAnyPeriod:
		return &pb.QmReply{},
			status.Errorf(codes.Unimplemented, "metering Quarterly feature is not implemented yet")
	case metering.TypeSettlement:
		ms, err := s.TheMeteringDM.ComputeSettlement(int(in.ProjectID))
		if err != nil {
			return &pb.QmReply{},
				status.Errorf(codes.NotFound, "MeteringDM.QueryWithCI error: %v", err)
		}
		return statementToQmReply(ms), nil
	default:
		return &pb.QmReply{},
			status.Errorf(codes.InvalidArgument, "metering Quarterly feature is not implemented yet")
	}
}

// Register Cron Metering Task
func (s *MeteringServer) RegisterCronMeteringTask(ctx context.Context, in *pb.RegCmtRequest) (*pb.RegCmtReply, error) {
	return &pb.RegCmtReply{
		Msg: "This feature is not yet implemented",
	}, nil
}

func statementToQmReply(ms metering.Statement) (reply *pb.QmReply) {
	return &pb.QmReply{
		ProjectID:            int64(ms.ProjectID),
		MeteringType:         int64(ms.MeteringType),
		MeteringTypeInfo:     ms.MeteringTypeInfo,
		CpuAmountInDays:      int64(ms.CpuAmountInDays),
		GpuAmountInDays:      int64(ms.GpuAmountInDays),
		StorageAmountInDays:  int64(ms.StorageAmountInDays),
		CpuAmountInHours:     int64(ms.CpuAmountInHours),
		GpuAmountInHours:     int64(ms.GpuAmountInHours),
		StorageAmountInHours: int64(ms.StorageAmountInHours),
		CpuNodeMeteringJson:  ms.CpuNodeMeteringJson,
		GpuNodeMeteringJson:  ms.GpuNodeMeteringJson,
		StorageMeteringJson:  ms.StorageMeteringJson,
		CreatedAt:            ms.CreatedAt.Format(time.RFC3339Nano),
	}
}
