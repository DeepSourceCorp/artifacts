package lock

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisLockManager provides distributed locking using Redis
type RedisLockManager struct {
	client *redis.Client
}

// Config holds Redis configuration
type Config struct {
	Addr     string
	Password string
	DB       int
}

// Lock represents an acquired lock
type Lock struct {
	key    string
	client *redis.Client
}

// NewRedisLockManager creates a new Redis lock manager
func NewRedisLockManager(config Config) (*RedisLockManager, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisLockManager{
		client: client,
	}, nil
}

// NewRedisLockManagerWithClient creates a Redis lock manager with existing client
func NewRedisLockManagerWithClient(client *redis.Client) *RedisLockManager {
	return &RedisLockManager{
		client: client,
	}
}

// AcquireLock attempts to acquire a distributed lock
// Returns the lock if successful, or an error if the lock is already held
func (r *RedisLockManager) AcquireLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {
	// Use SET with NX (only set if not exists) and EX (set expiration)
	result := r.client.SetNX(ctx, key, "locked", expiration)
	if result.Err() != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", result.Err())
	}

	if !result.Val() {
		return nil, fmt.Errorf("lock already held for key: %s", key)
	}

	return &Lock{
		key:    key,
		client: r.client,
	}, nil
}

// TryAcquireLock attempts to acquire a lock and returns whether it was successful
// Returns (lock, true, nil) if acquired, (nil, false, nil) if already held, or (nil, false, error) on error
func (r *RedisLockManager) TryAcquireLock(ctx context.Context, key string, expiration time.Duration) (*Lock, bool, error) {
	result := r.client.SetNX(ctx, key, "locked", expiration)
	if result.Err() != nil {
		return nil, false, fmt.Errorf("failed to acquire lock: %w", result.Err())
	}

	if !result.Val() {
		return nil, false, nil // Lock already held
	}

	return &Lock{
		key:    key,
		client: r.client,
	}, true, nil
}

// IsLocked checks if a lock exists for the given key
func (r *RedisLockManager) IsLocked(ctx context.Context, key string) (bool, error) {
	result := r.client.Exists(ctx, key)
	if result.Err() != nil {
		return false, fmt.Errorf("failed to check lock existence: %w", result.Err())
	}
	return result.Val() > 0, nil
}

// GetLockTTL returns the remaining time-to-live for a lock
func (r *RedisLockManager) GetLockTTL(ctx context.Context, key string) (time.Duration, error) {
	result := r.client.TTL(ctx, key)
	if result.Err() != nil {
		return 0, fmt.Errorf("failed to get lock TTL: %w", result.Err())
	}

	ttl := result.Val()
	if ttl == -2 {
		return 0, fmt.Errorf("lock does not exist for key: %s", key)
	}

	return ttl, nil
}

// ListLocks returns all keys matching the given pattern
func (r *RedisLockManager) ListLocks(ctx context.Context, pattern string) ([]string, error) {
	result := r.client.Keys(ctx, pattern)
	if result.Err() != nil {
		return nil, fmt.Errorf("failed to list locks: %w", result.Err())
	}
	return result.Val(), nil
}

// Release releases the lock
func (l *Lock) Release(ctx context.Context) error {
	result := l.client.Del(ctx, l.key)
	if result.Err() != nil {
		return fmt.Errorf("failed to release lock %s: %w", l.key, result.Err())
	}
	return nil
}

// Key returns the lock key
func (l *Lock) Key() string {
	return l.key
}

// Close closes the Redis client connection
func (r *RedisLockManager) Close() error {
	return r.client.Close()
}
