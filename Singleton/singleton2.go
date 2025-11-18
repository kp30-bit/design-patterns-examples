package singleton

import "sync"

type configManager struct {
	config map[string]string
	mu     sync.RWMutex // for thread safety, read lock and write lock
}

var configInstance *configManager
var configOnce sync.Once // for thread safety, ensures that the configManager is created only once

func GetConfigManager() *configManager {
	configOnce.Do(func() { // once.Do is used to ensure that the configManager is created only once
		configInstance = &configManager{config: make(map[string]string)}
	})
	return configInstance
}

func (c *configManager) SetConfig(key, value string) { // write lock is used to ensure that the config is not modified while it is being read
	c.mu.Lock()
	defer c.mu.Unlock()
	c.config[key] = value
}

func (c *configManager) GetConfig(key string) string { // read lock is used to ensure that the config is not modified while it is being read
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.config[key]
}
