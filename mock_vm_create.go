package ovirtclient

import (
	"github.com/google/uuid"
)

func (m *mockClient) CreateVM(clusterID string, name string, templateID string, _ ...RetryStrategy) (VM, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if err := validateVMCreationParameters(clusterID, name, templateID); err != nil {
		return nil, err
	}
	if _, ok := m.clusters[clusterID]; !ok {
		return nil, newError(ENotFound, "cluster with ID %s not found", clusterID)
	}
	if _, ok := m.templates[templateID]; !ok {
		return nil, newError(ENotFound, "template with ID %s not found", templateID)
	}

	id := uuid.Must(uuid.NewUUID()).String()
	vm := &vm{
		client: m,

		id:         id,
		clusterID:  clusterID,
		templateID: templateID,
	}
	m.vms[id] = vm
	return vm, nil
}
