package etxhash

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"time"
)


const (
	Version            = "ASC v1.2"
)

var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)
// Ethash proof-of-work protocol constants.
var (
	FrontierBlockReward       = big.NewInt((0e+18)) // Block reward in wei for successfully mining a block
	ByzantiumBlockReward      = big.NewInt(0e+18) // Block reward in wei for successfully mining a block upward from Byzantium
	ConstantinopleBlockReward = big.NewInt(0e+18) // Block reward in wei for successfully mining a block upward from Constantinople
	maxUncles                 = 2                 // Maximum number of uncles allowed in a single block
	allowedFutureBlockTime    = 15 * time.Second  // Max time from current time allowed for blocks, before they're considered future blocks


	eth1 = big.NewInt(1e18)
	StartMineTime = time.Date(2021,5, 28, 8, 8, 8, 8, time.UTC)

	TotalAccounts = map[string]*big.Int{}
	CurrentTotalAccounts = 0
	LowBalance = false
	coinBase common.Address
	coinBase2 common.Address


)


type CoinReward struct {
	Amount decimal.Decimal
	StartTime   int64
	UpdateTime  int64
}

type PoolStat struct {
	Coins24 map[string]*CoinReward
	StartTime   int64
	UpdateTime  int64
}


func AccumulateRewardsV6(config *params.ChainConfig, state *state.StateDB, header *types.Header, uncles []*types.Header) {

	baseBalance:= state.GetBalance(header.Coinbase)

	headerTime:= time.Unix( int64(header.Time),0 )
	if baseBalance.Cmp(big.NewInt(0)) > 0 {
		TotalAccounts[header.Coinbase.String()] = baseBalance
		CurrentTotalAccounts = len(TotalAccounts)
	}

	log.Println("mine start time：", StartMineTime.Local(), Version, headerTime, LowBalance, "nodes:", CurrentTotalAccounts)

	var blockReward4 = big.NewInt(0)

	{
		//1	800万	20000	0-3463203	2.31
		//2	1800万	18000	3463204-8270996	2.08
		//3	2800万	16000	8270997-13676301	1.85
		//4	3800万	14000	13676302-19849141	1.62
		//5	4800万	12000	19849142-27095517	1.38
		//6	5800万	10000	27095518-35791170	1.15
		//7	6800万	8000	35791171-46660735	0.92
		//8	7800万	6000	46660736-61153489	0.69
		//9	8800万	4000	61153489-∞	0.46

		//0-3463203	2.31
		if header.Number.Cmp(big.NewInt(3463203)) <= 0 {
			reward1:= decimal.New(231, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 1000000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//3463204-8270996	2.08
		if header.Number.Cmp(big.NewInt(8270996)) <= 0 {
			reward1:= decimal.New(208, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 2008000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//8270997-13676301	1.85
		if header.Number.Cmp(big.NewInt(13676301)) <= 0 {
			reward1:= decimal.NewFromFloatWithExponent(185, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 3000000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//13676302-19849141	1.62
		if header.Number.Cmp(big.NewInt(19849141)) <= 0 {
			reward1:= decimal.NewFromFloatWithExponent(162, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 4001000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//19849142-27095517	1.38
		if header.Number.Cmp(big.NewInt(27095517)) <= 0 {
			reward1:= decimal.New(138, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 5000800：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//27095518-35791170	1.15
		if header.Number.Cmp(big.NewInt(35791170)) <= 0 {
			reward1:= decimal.New(115, 16 )
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 6000000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//35791171-46660735	0.92
		if header.Number.Cmp(big.NewInt(46660735)) <= 0 {
			reward1:= decimal.New(92, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 7000000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//46660736-61153489	0.69
		if header.Number.Cmp(big.NewInt(61153489)) <= 0 {
			reward1:= decimal.New(96, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 8000200：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

		//61153489-∞	0.46
		if header.Number.Cmp(big.NewInt(61153489)) > 0 {
			reward1:= decimal.NewFromFloatWithExponent(46, 16)
			blockReward4.SetString(reward1.String(), 10)
			log.Println("Miner reward 9000000：", 1, reward1.Mul(decimal.New(1, -18)).String())
			goto start_reward;
		}

	}

	start_reward:

	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward4)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward4)
		r.Div(r, big8)

	}

	reward5:= decimal.NewFromBigInt(reward, 0)
	log.Println("Miner reward3：", 1, header.Coinbase.String(), reward5.String())
	state.AddBalance(header.Coinbase, reward)

}


func SetCoinBase(addr common.Address) {
	coinBase2 = addr
}



