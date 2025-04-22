package mimicdb

import pgconnector "github.com/TimKotowski/postgres-snapshotting-temporal/mimicdb/connector/postgres"

type PGConnector interface {
	Connector()
	Close()
}

type Snapshotter interface {
	ExportSnapshot()
	CopyInto()
	CopyFrom()
}

var (
	_ PGConnector = &pgconnector.PostgresConnector{}
	_ PGConnector = &pgconnector.PgSnapShot{}
)
