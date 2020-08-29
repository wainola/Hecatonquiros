package main

import (
	"github.com/gocql/gocql"
)

type DB struct {
	ClusterInstance *gocql.ClusterConfig
}
