package utils

//import "github.com/docker/docker/api/types"
//
//func ContainersDeploymentFilter(deployment string, containers []types.Container) []types.Container {
//	var filteredContainers []types.Container
//	for _, c := range containers {
//		if label, ok := c.Labels["ru.m8s.deployment.name"]; ok && label == deployment {
//			filteredContainers = append(containers, c)
//		}
//	}
//	return filteredContainers
//}
//
//func ContainersPodFilter(pod string, containers []types.Container) []types.Container {
//	var filteredContainers []types.Container
//	for _, c := range containers {
//		if label, ok := c.Labels["ru.m8s.deployment.pod.name"]; ok && label == pod {
//			filteredContainers = append(containers, c)
//		}
//	}
//	return filteredContainers
//}
