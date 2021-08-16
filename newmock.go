package ovirtclient

import (
	"sync"

	"github.com/google/uuid"
)

// NewMock creates a new in-memory mock client. This client can be used as a testing facility for
// higher level code.
func NewMock() MockClient {
	testCluster := generateTestCluster()
	testHost := generateTestHost(testCluster)
	testStorageDomain := generateTestStorageDomain()
	testDatacenter := generateTestDatacenter(testCluster)
	blankTemplate := &template{
		id:          BlankTemplateID,
		name:        "Blank",
		description: "Blank template",
	}

	client := &mockClient{
		url:  "https://localhost/ovirt-engine/api",
		lock: &sync.Mutex{},
		storageDomains: map[string]*storageDomain{
			testStorageDomain.ID(): testStorageDomain,
		},
		disks: map[string]*diskWithData{},
		clusters: map[string]*cluster{
			testCluster.ID(): testCluster,
		},
		hosts: map[string]*host{
			testHost.ID(): testHost,
		},
		templates: map[string]*template{
			blankTemplate.ID(): blankTemplate,
		},
		dataCenters: map[string]*datacenterWithClusters{
			testDatacenter.ID(): testDatacenter,
		},
	}

	testCluster.client = client
	testHost.client = client
	blankTemplate.client = client
	testStorageDomain.client = client
	testDatacenter.client = client

	return client
}

func generateTestDatacenter(testCluster *cluster) *datacenterWithClusters {
	return &datacenterWithClusters{
		datacenter: datacenter{
			id:   uuid.NewString(),
			name: "test",
		},
		clusters: []string{
			testCluster.ID(),
		},
	}
}

func generateTestStorageDomain() *storageDomain {
	return &storageDomain{
		id:             uuid.NewString(),
		name:           "Test storage domain",
		available:      10 * 1024 * 1024 * 1024,
		status:         StorageDomainStatusActive,
		externalStatus: StorageDomainExternalStatusNA,
	}
}

func generateTestCluster() *cluster {
	return &cluster{
		id:   uuid.NewString(),
		name: "Test cluster",
	}
}

func generateTestHost(c *cluster) *host {
	return &host{
		id:        uuid.NewString(),
		clusterID: c.ID(),
		status:    HostStatusUp,
	}
}
