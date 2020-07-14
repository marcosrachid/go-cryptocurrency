package global

const (
	// DECIMAL max number of decimal characters
	DECIMAL        uint8 = 8
	DECIMAL_STRING       = "%.8f"

	// SUPPLY_LIMIT Max number of coins in network, after the supply reaches the value, no more reward will be given anymore
	SUPPLY_LIMIT float64 = 1000000000.0

	// MINING_TIME_RATE Network will lookup for the value below in nanoseconds between blocks mining with a marin of error of the value below
	MINING_TIME_RATE       uint64  = 600000000000 // Nanoseconds
	MINING_TIME_RATE_ERROR float64 = 0.5

	// MAX_DIFFICULTY Max difficulty on network due the maximum hash characters of 64
	MAX_DIFFICULTY uint8 = 32

	// REWARD Mocked reward value to be given to miners, no halven implemented
	REWARD float64 = 200.0

	// P1_FEES Using static fees based on block transactions instead of bitcoin algorithm
	P1_FEES float64 = 0.005
	P2_FEES float64 = 0.0025
	P3_FEES float64 = 0.001

	// BLOCK_SIZE Max block transactions size in bytes on every block
	BLOCK_SIZE uint64 = 1000000

	// COINBASE Coinbase name to be used on rewards transactions
	COINBASE = "Socialism will always fail"

	// DIFFICULTY_ADJUSTMENT_BLOCK Network difficulty will be adjusted after every below number of blocks
	DIFFICULTY_ADJUSTMENT_BLOCK uint64 = 100

	// END Network message end string
	END = "\r\n\r\n"
)
