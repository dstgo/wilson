package dstparser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseClusterIni(t *testing.T) {
	filename := "testdata/cluster/cluster.ini"
	bytes, err := os.ReadFile(filename)
	assert.Nil(t, err)
	clusterIni, err := ParseClusterInI(bytes)
	assert.Nil(t, err)
	t.Logf("%+v", clusterIni)
}

func TestToClusterIni(t *testing.T) {
	filename := "testdata/cluster/cluster.ini"
	bytes, err := os.ReadFile(filename)
	assert.Nil(t, err)
	clusterIni, err := ParseClusterInI(bytes)
	assert.Nil(t, err)
	t.Logf("%+v", clusterIni)

	clusterInIData, err := ToClusterInI(clusterIni)
	assert.Nil(t, err)

	fmt.Println(string(clusterInIData))
}

func TestParseServerIni(t *testing.T) {
	filename := "testdata/cluster/server.master.ini"
	bytes, err := os.ReadFile(filename)
	assert.Nil(t, err)
	serverIni, err := ParseServerInI(bytes)
	assert.Nil(t, err)
	t.Logf("%+v", serverIni)
}

func TestToServerInI(t *testing.T) {
	filename := "testdata/cluster/server.master.ini"
	bytes, err := os.ReadFile(filename)
	assert.Nil(t, err)
	serverIni, err := ParseServerInI(bytes)
	assert.Nil(t, err)
	t.Logf("%+v", serverIni)

	serverInIData, err := ToServerInI(serverIni)
	assert.Nil(t, err)

	fmt.Println(string(serverInIData))
}
