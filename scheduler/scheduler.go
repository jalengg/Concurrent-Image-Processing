package scheduler

type Config struct {
	DataDirs string //Represents the data directories to use to load the images.
	Mode     string // Represents which scheduler scheme to use
	// If Mode == "s" run the sequential version
	// If Mode == "balance" run the balancing executorservice version
	// If Mode == "steal" run the stealing executorservice version
	// These are the only values for Version
	ThreadCount int // Runs the parallel version of the program with the
	// specified number of threads (i.e., goroutines)
}

//Run the correct version based on the Mode field of the configuration value
func Schedule(config Config) {
	if config.Mode == "s" {
		RunSequential(config)
	} else if config.Mode == "balance" || config.Mode == "steal"{
		RunParallel(config)
	} else {
		panic("Invalid scheduling scheme given.")
	}
}
