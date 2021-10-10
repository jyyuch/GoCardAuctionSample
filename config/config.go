package config

/**
 * config.go could use viper to get var from env for k8s deployment.
 */
const (
	NUM_ROUTINE_TO_SCAN  = 10
	NUM_BLOCKS_SCAN_ONCE = 20
	PG_DSN               = "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432"
)
