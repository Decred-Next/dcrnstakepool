package stakepool

import (
	"fmt"
	"github.com/decred/dcrd/chaincfg/v2"
	"testing"
)

func TestStakePoolTicketFee(t *testing.T) {
	//params := chaincfg.MainNetParams()
	//tests := []struct {
	//	StakeDiff dcrutil.Amount
	//	Fee       dcrutil.Amount
	//	Height    int32
	//	PoolFee   float64
	//	Expected  dcrutil.Amount
	//}{
	//	0: {10 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.01500463 * 1e8},
	//	1: {20 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.01621221 * 1e8},
	//	2: {5 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.03310616 * 1e8},
	//	3: {15 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.03956376 * 1e8},
	//}
	//for i, test := range tests {
	//
	//	if poolFeeAmt != test.Expected {
	//		t.Errorf("Test %d: Got %v: Want %v", i, poolFeeAmt, test.Expected)
	//	}
	//}
	params := chaincfg.TestNet3Params()
	poolFeeAmt := StakePoolTicketFee(200000000, 5420, 719,
		5, params)
	fmt.Println(poolFeeAmt)
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 1067,
		5, params)
	fmt.Println(poolFeeAmt) //0.04261479
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 9067,
		5, params)
	fmt.Println(poolFeeAmt)
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 719,
		0.2, params)
	fmt.Println(poolFeeAmt)
	//200000000
	//200000000
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 9067,
		0, params)
	fmt.Println(poolFeeAmt)
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 1193,
		0.2, params)
	fmt.Println(poolFeeAmt)
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 1196,
		7.5, params)
	fmt.Println(poolFeeAmt)
	poolFeeAmt = StakePoolTicketFee(200000000, 5420, 1196,
		7.5, params)
	fmt.Println(poolFeeAmt)
}
