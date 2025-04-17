package core

type PGConnector interface {
	Connector()
	Close()
}

type Snapshotter interface {
	ExportSnapshot()
}

type Partitioner interface {
	GetPartitions()
}

type PartitionSyncPoller interface {
	PollPartitions()
}
