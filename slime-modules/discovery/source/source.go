package meshsource

type Source interface {
	Start()
	Stop()
	SetMappingNamespace(namespace string)
	GetMappingNamespace()string
}
