package bandwidth

const (
	// Maximum bandwidth of network
	DesirableNetworkBandwidthForRecoveryPeriod = LinkMsgCost * 2000 * 1000

	// Bandwidth cost of specific messages and tx itself
	LinkMsgCost    int64 = 100
	TxCost               = LinkMsgCost * 3
	NonLinkMsgCost       = LinkMsgCost * 5

	// Number of blocks to recover full bandwidth
	RecoveryPeriod = 18000
	// Number of blocks before next adjust price
	AdjustPricePeriod                     = 10
	BaseCreditPrice               float64 = 1.0
	SlidingWindowSize                     = RecoveryPeriod
	ShouldBeSpentPerSlidingWindow         = float64(DesirableNetworkBandwidthForRecoveryPeriod)
)
