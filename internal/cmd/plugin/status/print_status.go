package status

type PrintClusterSummary struct {
	TargetSummary           interface{} `table:"Target Summary"`
	Name                    string      `table:"Name"`
	Namespace               string      `table:"Namespace"`
	SystemID                *string     `table:"First Name"`
	PostgreSQLImage         string      `table:"PostgreSQL Image"`
	DesignatedPrimary       *string     `table:"Designated primary"`
	SourceCluster           *string     `table:"Source cluster"`
	PrimaryInstance         *string     `table:"Primary instance"`
	PrimaryStartTime        *string     `table:"Primary start time"`
	PrimaryStartTimeError   interface{} `table:"Primary start time error"`
	Status                  string      `table:"Status"`
	LiveInstances           *int        `table:"First Name"`
	MissingInstances        *int        `table:"First Name"`
	ReadyInstances          *int        `table:"First Name"`
	NotReadyInstances       *int        `table:"First Name"`
	PrimaryInstanceFenced   *string     `table:"First Name"`
	FencedInstances         *string     `table:"First Name"`
	PrimaryTransitionStatus *string     `table:"First Name"`
	CurrentWriteLSN         *string     `table:"First Name"`
}

func strPointer(v string) *string { return &v }

// func ifacePointer(v interface{}) *interface{} { return &v }
