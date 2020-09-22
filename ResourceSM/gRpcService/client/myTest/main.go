package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/gRpcService/client"
	"github.com/sirupsen/logrus"
)

func main() {
	var resSchedulingClient client.ResSchedulingClient
	const (
		host = "localhost"
		port = 50071
	)
	resSchedulingClient = client.NewResSchedulingClient(host, port)

	///////////////////////////////////////////////////////
	projectID := 1
	cgpuType := 1
	nodesAfter := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ctrlID := 1
	ctrlCN := "zhang"
	allocInfo, err := resSchedulingClient.SchedulingCGpu(projectID, cgpuType, nodesAfter, ctrlID, ctrlCN)
	if err != nil {
		logrus.Errorf("resSchedulingClient.SchedulingCGpu CPU error: %v", err)
	}
	fmt.Printf("allocInfo=%v \n", allocInfo)

	///////////
	nodesAfter = []int64{6, 7, 8, 9, 10}
	allocInfo, err = resSchedulingClient.SchedulingCGpu(projectID, cgpuType, nodesAfter, ctrlID, ctrlCN)
	if err != nil {
		logrus.Errorf("resSchedulingClient.SchedulingCGpu CPU error: %v", err)
	}
	fmt.Printf("allocInfo=%v \n", allocInfo)

	///////////
	nodesAfter = []int64{}
	allocInfo, err = resSchedulingClient.SchedulingCGpu(projectID, cgpuType, nodesAfter, ctrlID, ctrlCN)
	if err != nil {
		logrus.Errorf("resSchedulingClient.SchedulingCGpu CPU error: %v", err)
	}
	fmt.Printf("allocInfo=%v \n", allocInfo)

	///////////////////////////////////////////////////////
	projectID = 1
	cgpuType = 2
	nodesAfter = []int64{31, 32, 33, 39, 40}

	allocInfo, err = resSchedulingClient.SchedulingCGpu(projectID, cgpuType, nodesAfter, ctrlID, ctrlCN)
	if err != nil {
		logrus.Errorf("resSchedulingClient.SchedulingCGpu GPU error: %v", err)
	}
	fmt.Printf("allocInfo=%v \n", allocInfo)

	///////////////////////////////////////////////////////
	projectID = 1
	nodesAfter = []int64{31, 32, 33, 39, 40}
	storageSizeAfter := 100
	storageAllocInfoAfter := "/home/zhang"

	allocInfo, err = resSchedulingClient.SchedulingStorage(projectID, storageSizeAfter, storageAllocInfoAfter, ctrlID, ctrlCN)
	if err != nil {
		logrus.Errorf("resSchedulingClient.SchedulingCGpu Storage error: %v", err)
	}
	fmt.Printf("allocInfo=%v \n", allocInfo)
}
