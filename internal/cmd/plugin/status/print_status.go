package status

type PrintPostgresqlStatus struct {
	ClusterSummary        *ClusterSummary
	ReplicaClusterSummary *ClusterSummary
	SwitchingPrimary      *string
}

type ClusterSummary struct {
	Name                  string
	Namespace             string
	SystemID              *string
	PostgreSQLImage       string
	DesignatedPrimary     *string
	SourceCluster         *string
	PrimaryInstance       *string
	Status                string
	LiveInstances         *int
	MissingInstances      *int
	ReadyInstances        *int
	NotReadyInstances     *int
	PrimaryInstanceFenced *string
	FencedInstances       *string
}
