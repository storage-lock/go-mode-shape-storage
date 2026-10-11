package modeshape_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// ModeShapeStorage mode-shape存储实现
// 尚未接入真实服务端，仅为占位实现，不可用于分布式锁。
// 真正接入时需借助服务端原子语义实现 CAS。
type ModeShapeStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &ModeShapeStorage{}

func NewModeShapeStorage(ctx context.Context) (*ModeShapeStorage, error) {
	return &ModeShapeStorage{}, nil
}

const StorageName = "mode-shape-storage"

func (x *ModeShapeStorage) GetName() string { return StorageName }
func (x *ModeShapeStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{storage.CapabilityReliableTime, storage.CapabilityWatch}
}

func (x *ModeShapeStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}

func (x *ModeShapeStorage) Init(ctx context.Context) error { return nil }
func (x *ModeShapeStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *ModeShapeStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *ModeShapeStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *ModeShapeStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}
func (x *ModeShapeStorage) GetTime(ctx context.Context) (time.Time, error) { return time.Now(), nil }
func (x *ModeShapeStorage) Close(ctx context.Context) error { return nil }
func (x *ModeShapeStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}
