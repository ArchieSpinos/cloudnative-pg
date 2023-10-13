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
	LiveInstances           interface{} `table:"Live Instances"`
	MissingInstances        interface{} `table:"Missing Instances"`
	ReadyInstances          interface{} `table:"Ready Instances"`
	UnreadyInstances        interface{} `table:"Unready Instances"`
	PrimaryInstanceFenced   interface{} `table:"Fenced instances including primary"`
	FencedInstances         interface{} `table:"Fenced instances not including primary"`
	PrimaryTransitionStatus interface{} `table:"Primary Transition Status"`
	CurrentWriteLSN         *string     `table:"Current Write LSN"`
}

func strPointer(v string) *string { return &v }

// func ifacePointer(v interface{}) *interface{} { return &v }
