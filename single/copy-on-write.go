package main

type Config struct {
    Routes   map[string]net.Addr
    Backends []net.Addr
}config unsafe.Pointer  // actual type is *Config

// Worker goroutines use this function to obtain the current config.
func CurrentConfig() *Config {
    return (*Config)(atomic.LoadPointer(&config))
}

// Background goroutine periodically creates a new Config object
// as sets it as current using this function.
func UpdateConfig(cfg *Config) {
    atomic.StorePointer(&config, unsafe.Pointer(cfg))
}

func main(){
	
}