package zookeeper

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	networking "istio.io/api/networking/v1alpha3"
)

const (
	Dubbo        = "dubbo://"
	Rest         = "rest://"
	Protocol     = "dubbo"
	DUBBO_SUFFIX = ".dubbo"
)
const (
	dubbo_param_group_key          = "group"
	dubbo_param_version_key        = "version"
	dns1123LabelMaxLength   int    = 63
	dns1123LabelFmt         string = "[a-zA-Z0-9]([-a-z-A-Z0-9]*[a-zA-Z0-9])?"
	// a wild-card prefix is an '*', a normal DNS1123 label with a leading '*' or '*-', or a normal DNS1123 label
	wildcardPrefix = `(\*|(\*|\*-)?` + dns1123LabelFmt + `)`

	// Using kubernetes requirement, a valid key must be a non-empty string consist
	// of alphanumeric characters, '-', '_' or '.', and must start and end with an
	// alphanumeric character (e.g. 'MyValue',  or 'my_value',  or '12345'
	qualifiedNameFmt string = "([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]"
)

var (
	wildcardPrefixRegexp = regexp.MustCompile("^" + wildcardPrefix + "$")
)

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}

func convertServiceEntry(providers, consumers []string, service string, gatewayModel bool) (map[string]*networking.ServiceEntry, error) {

	if providers == nil || len(providers) == 0 {
		return nil, Error{
			msg: fmt.Sprintf("This no instance in SVC: %s", service),
		}
	}
	uniquePort := make(map[string]map[uint32]struct{})
	serviceEntryByServiceKey := make(map[string]*networking.ServiceEntry)
	// 获取provider 信息
	for _, child := range providers {
		ss := splitUrl(child)
		if ss == nil {
			continue
		}
		address := strings.Split(ss[2], ":")
		if verifyIP(address) {
			p, _ := strconv.Atoi(address[1])
			portNum := uint32(p)
			port := convertPort(portNum)
			if meta, ok := verifyMeta(ss[3]); ok {
				serviceKey := buildServiceKey(service, meta)
				var se *networking.ServiceEntry
				if _, ok := serviceEntryByServiceKey[serviceKey]; !ok {
					se = &networking.ServiceEntry{
						Ports: make([]*networking.Port, 0),
					}
					if gatewayModel {
						se.Hosts = []string{serviceKey + DUBBO_SUFFIX}
					} else {
						se.Hosts = []string{serviceKey}
					}
					se.Resolution = networking.ServiceEntry_STATIC
					serviceEntryByServiceKey[serviceKey] = se
					// todo build sidecar event
					/*for _, child := range consumers {
						ss := splitUrl(child)
						if ss == nil {
							continue
						}
						address := strings.Split(ss[2], ":")
						var port *networking.Port
						if len(address) == 2 {
							p, _ := strconv.Atoi(address[1])
							portNum := uint32(p)
							port = convertPort(portNum)
						} else {
							port = nil
						}
						if meta, ok := verifyMeta(ss[3]); ok {
							se.InboundEndPoints = append(se.InboundEndPoints, convertInboundEndpoint(address[0], meta, port))
						}
					}*/
				} else {
					se = serviceEntryByServiceKey[serviceKey]
				}
				se.Endpoints = append(se.Endpoints, convertEndpoint(address[0], meta, port))
				// 防止重复
				if _, ok := uniquePort[serviceKey]; !ok {
					uniquePort[serviceKey] = make(map[uint32]struct{})
				}

				if _, ok := uniquePort[serviceKey][port.Number]; !ok {
					if gatewayModel {
						//for gateway model use define 80 as ServiceEntry
						portNum := uint32(80)
						port := convertPort(portNum)
						if len(se.Ports) == 0 {
							se.Ports = append(se.Ports, port)
						}
					} else {
						se.Ports = append(se.Ports, port)
					}
					uniquePort[serviceKey][port.Number] = struct{}{}
				}
			}
		} else {
			//log.Errorf("Invaild IP address")
			continue
		}
	}

	return serviceEntryByServiceKey, nil
}

func verifyIP(address []string) bool {
	if len(address) != 2 {
		return false
	}
	if net.ParseIP(address[0]) == nil {
		return false
	}
	_, err := strconv.Atoi(address[1])
	return err == nil
}

func buildServiceKey(service string, meta map[string]string) string {
	group := meta[dubbo_param_group_key]
	version := meta[dubbo_param_version_key]
	serviceKey := service
	if len(group) > 0 {
		serviceKey = serviceKey + ":" + group
	}
	if len(version) > 0 {
		if group == "" {
			serviceKey = serviceKey + "::" + version
		} else {
			serviceKey = serviceKey + ":" + version
		}
	}
	return serviceKey
}

func verifyMeta(metaStr string) (map[string]string, bool) {
	if !strings.Contains(metaStr, "?") {
		//log.Errorf("Invaild dubbo url, missing '?'")
		return nil, false
	}
	metaStr = metaStr[strings.Index(metaStr, "?")+1:]
	entries := strings.Split(metaStr, "&")
	meta := make(map[string]string, len(entries))
	for _, entry := range entries {
		kv := strings.Split(entry, "=")
		if len(kv) != 2 {
			//log.Errorf("Invaild dubbo url, invaild meta info")
			return nil, false
		}
		kv[1] = strings.ReplaceAll(kv[1], ",", "_")
		kv[1] = strings.ReplaceAll(kv[1], ":", "_")
		meta[kv[0]] = kv[1]
		/*		if wildcardPrefixRegexp.MatchString(kv[1]) {

				} else {
					log.Warnf("invalid tag value: %s", kv[1])
				}*/
	}
	return meta, true
}

func convertPort(port uint32) *networking.Port {
	return &networking.Port{
		// Todo: change to dubbo, when pilot/envoy support dubbo
		Protocol: Protocol,
		Number:   port,
		Name:     Protocol,
	}
}

func convertEndpoint(ip string, meta map[string]string, port *networking.Port) *networking.WorkloadEntry {
	return &networking.WorkloadEntry{
		Address: ip,
		Ports: map[string]uint32{
			port.Name: port.Number,
		},
		Labels: meta,
	}
}

/*func convertInboundEndpoint(ip string, meta map[string]string, port *networking.Port) *networking.ServiceEntry_InboundEndPoint {
	inboundEndpoint := &networking.ServiceEntry_InboundEndPoint{
		Address: ip,
		Labels:  meta,
	}
	if port != nil {
		inboundEndpoint.Ports = map[string]uint32{
			port.Name: port.Number,
		}
	}
	return inboundEndpoint
}*/

func splitUrl(zkChild string) []string {
	path, err := url.PathUnescape(zkChild)
	if err != nil {
		// log.Errorf(err.Error())
		return nil
	}
	if strings.HasPrefix(path, Rest) {
		return nil
	}
	ss := strings.Split(path, "/")
	if len(ss) < 4 {
		//log.Errorf("Invaild dubbo Url")
		return nil
	}
	return ss

}
