package lb

import (
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

var (
	LB *LoadBalancer = NewLoadBalancer()
)

type Instance struct {
	URL string
}

type ServiceBalancer struct {
	instances []*Instance
	next      int
}

type LoadBalancer struct {
	services map[string]*ServiceBalancer
	hasProxy map[string]bool
}

func NewServiceBalancer(url string) *ServiceBalancer {
	instances := make([]*Instance, 0)
	instances = append(instances, NewInstance(url))
	return &ServiceBalancer{
		instances: instances,
		next:      0,
	}
}
func NewInstance(url string) *Instance {
	return &Instance{URL: url}
}
func (sb *ServiceBalancer) GetCurrentInstance() *Instance {
	if len(sb.instances) == 0 {
		return nil
	}
	if len(sb.instances) == 1 {
		return sb.instances[0]
	}
	if sb.next == 0 {
		return sb.instances[len(sb.instances)-1]
	}
	return sb.instances[sb.next-1]
}
func (sb *ServiceBalancer) GetNextInstance() *Instance {
	if len(sb.instances) == 0 {
		return nil
	} else if len(sb.instances) == 1 {
		return sb.instances[0]
	} else if len(sb.instances)-1 < sb.next {
		sb.next = 0
	}
	instance := sb.instances[sb.next]
	sb.next = (sb.next + 1) % len(sb.instances)
	return instance
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		services: make(map[string]*ServiceBalancer),
		hasProxy: make(map[string]bool),
	}
}

func (lb *LoadBalancer) AddService(name string, url string) {
	if _, ok := lb.services[name]; !ok {
		lb.services[name] = NewServiceBalancer(url)
	} else {
		insts := lb.services[name].instances
		logutil.Info("inst: %v", insts)
		for _, v := range insts {
			if url == v.URL {
				// no need to add service if the same url already exists in the instances
				return
			}
		}
		lb.services[name].instances = append(lb.services[name].instances, NewInstance(url))
	}
}

func (lb *LoadBalancer) GetServiceBalancer(name string) *ServiceBalancer {
	return lb.services[name]
}

func (lb *LoadBalancer) GetAllBalancer() map[string]*ServiceBalancer {
	return lb.services
}

func (lb *LoadBalancer) SetHasProxy(name string, hasProxy bool) {
	lb.hasProxy[name] = hasProxy
}

func (lb *LoadBalancer) GetHasProxy(name string) bool {
	if _, ok := lb.hasProxy[name]; ok {
		return lb.hasProxy[name]
	}
	return false
}

// remove offline service balancer
func (lb *LoadBalancer) RemoveOfflineService(serviceName, svcUrl string) {
	instanceArr := lb.GetServiceBalancer(serviceName).instances
	logutil.Info("instanceArr: %v", instanceArr)
	for idx, val := range instanceArr {
		logutil.Info("svcUrl: %v, val.URL: %v", svcUrl, val.URL)
		if val.URL == svcUrl {
			logutil.Info("remove offline service: %v", svcUrl)
			lb.GetServiceBalancer(serviceName).instances = append(instanceArr[:idx], instanceArr[idx+1:]...)
			lb.GetServiceBalancer(serviceName).next = 0
			break
		}
	}
}
