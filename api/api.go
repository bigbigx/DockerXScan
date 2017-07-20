package api

import (
	"time"
	"github.com/MXi4oyu/DockerXScan/database"
	"github.com/MXi4oyu/DockerXScan/common/stopper"
	"github.com/MXi4oyu/DockerXScan/versionfmt/dpkg"
)


type Config struct {
	Port                      int
	HealthPort                int
	Timeout                   time.Duration
	PaginationKey             string
	CertFile, KeyFile, CAFile string
}

func Run(cfg *Config, store database.Datastore, st *stopper.Stopper)  {
	//插入一条namespace
	store.InsertNamespace(database.Namespace{
		Name:"debian:8",
		VersionFormat:dpkg.ParserName,
	})





	defer st.End()
}