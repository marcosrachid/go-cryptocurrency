package global

const (
	MINIMUM_TRANSACTION         float64 = 0.00000001
	SUPPLY_LIMIT                float64 = 1000000000.0
	MINING_TIME_RATE            uint64  = 600000000000 //nanoseconds
	MINING_TIME_RATE_ERROR      float64 = 0.5
	MAX_DIFFICULTY              uint8   = 32
	REWARD                      float64 = 200.0
	FEES                        float64 = 0.001
	BLOCK_SIZE                  uint64  = 1000000
	COINBASE                            = "Lorem ipsum dolor sit amet"
	DIFFICULTY_ADJUSTMENT_BLOCK uint64  = 2016
	END                                 = "\r\n\r\n"
)
