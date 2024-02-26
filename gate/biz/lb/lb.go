package lb

type Instance struct {
	URL string
}

type ServiceBalancer struct {
	instances []*Instance
	next      int
}

type LoadBalancer struct {
	services map[string]*ServiceBalancer
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

func (sb *ServiceBalancer) GetNextInstance() *Instance {
	instance := sb.instances[sb.next]
	sb.next = (sb.next + 1) % len(sb.instances)
	return instance
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		services: make(map[string]*ServiceBalancer),
	}
}

func (lb *LoadBalancer) AddService(name string, url string) {
	if _, ok := lb.services[name]; !ok {
		lb.services[name] = NewServiceBalancer(url)
	} else {
		lb.services[name].instances = append(lb.services[name].instances, NewInstance(url))
	}
}

func (lb *LoadBalancer) GetServiceBalancer(name string) *ServiceBalancer {
	return lb.services[name]
}

func (lb *LoadBalancer) GetAllBalancer() map[string]*ServiceBalancer {
	return lb.services
}
