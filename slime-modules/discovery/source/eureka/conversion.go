package eureka

import (
	"sort"
	"strings"

	networking "istio.io/api/networking/v1alpha3"
)

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}

func ConvertServiceEntryMap(apps []*application, gatewayModel bool) (map[string]*networking.ServiceEntry, error) {
	seMap := make(map[string]*networking.ServiceEntry, 0)
	if apps == nil || len(apps) == 0 {
		return seMap, nil
	}
	for _, app := range apps {
		if gatewayModel {
			seMap[app.Name] = convertServiceEntry(app)
		} else {
			for k, v := range convertServiceEntryWithNs(app) {
				seMap[k] = v
			}
		}
	}
	return seMap, nil
}

// -------- for gateway mode --------
func convertServiceEntry(app *application) *networking.ServiceEntry {
	endpoints, ports := convertEndpoints(app.Instances)
	return &networking.ServiceEntry{
		Hosts:      []string{strings.ReplaceAll(strings.ToLower(app.Name), "_", "-") + ".eureka"},
		Resolution: networking.ServiceEntry_DNS,
		Endpoints:  endpoints,
		Ports:      ports,
	}
}

func convertEndpoints(instances []*instance) ([]*networking.WorkloadEntry, []*networking.Port) {
	endpoints := make([]*networking.WorkloadEntry, 0)
	ports := make([]*networking.Port, 0)
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].App < instances[j].App
	})

	port := &networking.Port{
		Protocol: "HTTP",
		Number:   80,
		Name:     "http",
	}
	ports = append(ports, port)

	for _, ins := range instances {
		instancePorts := make(map[string]uint32, 1)
		for _, v := range ports {
			instancePorts[v.Name] = uint32(ins.Port.Port)
		}
		filterLabels(ins.Metadata)
		endpoints = append(endpoints,
			&networking.WorkloadEntry{
				Address: ins.IPAddress,
				Ports:   instancePorts,
				Labels:  ins.Metadata,
			})
	}
	sort.Slice(ports, func(i, j int) bool {
		return ports[i].Number < ports[j].Number
	})
	return endpoints, ports
}

// -------- for sidecar mode --------
func convertServiceEntryWithNs(app *application) map[string]*networking.ServiceEntry {
	endpointMap, portMap := convertEndpointsWithNs(app.Instances)
	if len(endpointMap) > 0 {
		ses := make(map[string]*networking.ServiceEntry, len(endpointMap))
		for ns, endpoints := range endpointMap {
			ses[strings.ToLower(app.Name)+"."+ns] = &networking.ServiceEntry{
				Hosts:      []string{strings.ToLower(app.Name) + "." + ns + ".svc.cluster.local"},
				Resolution: networking.ServiceEntry_STATIC,
				Endpoints:  endpoints,
				Ports:      portMap[ns],
			}
		}
		return ses
	}
	return nil
}

func convertEndpointsWithNs(instances []*instance) (map[string][]*networking.WorkloadEntry, map[string][]*networking.Port) {
	endpointsMap := make(map[string][]*networking.WorkloadEntry, 0)
	portsMap := make(map[string][]*networking.Port, 0)

	for _, ins := range instances {
		filterLabels(ins.Metadata)
		metadata := ins.Metadata
		ns, exist := metadata["k8sNs"]
		if !exist {
			ns = "eureka"
		}
		endpoints, exist := endpointsMap[ns]
		if !exist {
			endpoints = make([]*networking.WorkloadEntry, 0)
		}
		ports, exist := portsMap[ns]
		if !exist {
			ports = make([]*networking.Port, 0)
			port := &networking.Port{
				Protocol: "HTTP",
				Number:   uint32(ins.Port.Port),
				Name:     "http",
			}
			ports = append(ports, port)
			portsMap[ns] = ports
		}

		instancePorts := make(map[string]uint32, 1)
		instancePorts["http"] = uint32(ins.Port.Port)

		endpoints = append(endpoints,
			&networking.WorkloadEntry{
				Address: ins.IPAddress,
				Ports:   instancePorts,
				Labels:  ins.Metadata,
			})
		endpointsMap[ns] = endpoints
	}
	return endpointsMap, portsMap
}

func filterLabels(labels map[string]string) {
	for k, _ := range labels {
		if k == "@class" {
			delete(labels, k)
		}
	}
}
