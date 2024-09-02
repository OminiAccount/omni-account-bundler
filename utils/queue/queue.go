package queue

//
//import (
//	"fmt"
//	"github.com/OAAC/database"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethdb"
//	"github.com/ethereum/go-ethereum/log"
//)
//
//type ChallengeQueue struct {
//	ChallengeInput *utils.ChallengeInput
//	Provider       *providers.BaseProvider
//}
//
//type DispatchChallengesQueue struct {
//	work                  chan *ChallengeQueue
//	submitChallengesPools map[common.Hash]*ChallengeQueue
//	commitChallengesPools map[common.Hash]*ChallengeQueue
//	status                chan int
//	Read                  bool // collect challenge -> queue
//	Write                 bool // queue -> submit challenge
//
//	db ethdb.Database
//
//	waitTime *utils.InstantWaitTime
//
//	logger log.Logger
//}
//
//const (
//	ChanMax = 200
//
//	CollectChallengesChannelStop = 0
//	SubmitChallengesChanStop     = 1
//	CommitChallengesChannelStop  = 2
//)
//
//func NewQueue(leveldb ethdb.Database) *DispatchChallengesQueue {
//	q := &DispatchChallengesQueue{
//		work:                  make(chan *ChallengeQueue, ChanMax),
//		submitChallengesPools: make(map[common.Hash]*ChallengeQueue, ChanMax),
//		commitChallengesPools: make(map[common.Hash]*ChallengeQueue, ChanMax),
//		status:                make(chan int),
//		Read:                  true,
//		Write:                 true,
//		db:                    leveldb,
//		logger:                log.New("service", "channel"),
//	}
//	go func() {
//		for {
//			select {
//			case status := <-q.status:
//				switch status {
//				case CollectChallengesChannelStop:
//					q.Read = false
//				case SubmitChallengesChanStop:
//					q.logger.Info("Submit-Channel is closed, waiting for Commit-Channel to close")
//				case CommitChallengesChannelStop:
//					q.Write = false
//					q.logger.Info("Commit-Channel is closed")
//				}
//			}
//		}
//	}()
//
//	q.logger.Info("Start channel for work with proof")
//	return q
//}
//
//// Enqueue Add a new challenge to the last position in the queue
//func (q *DispatchChallengesQueue) Enqueue(pq *ChallengeQueue) error {
//	return q.collectChallenge(pq)
//}
//
//func (q *DispatchChallengesQueue) Submit() {
//	q.submitChallenges()
//}
//
//// Dequeue commit challenge
//func (q *DispatchChallengesQueue) Dequeue() {
//	q.commitChallenges()
//}
//
//// Stop reading channel information, but does not affect sending information to the channel
//func (q *DispatchChallengesQueue) Stop() {
//	q.logger.Warn("Stop the Collect-Channel, but does not affect the Submit-Channel or Commit-Channel")
//	q.logger.Info("Processing remaining commit challenges in submitChallengesPools and commitChallengesPools", "submitChallengesPools-number", q.SubmitChallengesPoolsNumber(), "commitChallengesPools-number", q.CommitChallengesPoolsNumber())
//	q.status <- CollectChallengesChannelStop
//	for {
//		if q.SubmitChallengesPoolsIsEmpty() {
//			q.status <- SubmitChallengesChanStop
//			break
//		}
//	}
//	for {
//		if q.CommitChallengesPoolsIsEmpty() {
//			q.status <- CommitChallengesChannelStop
//			return
//		}
//	}
//}
//
//// Restart the channel service
//func (q *DispatchChallengesQueue) Restart() {
//	if q.Write || q.Read {
//		return
//	}
//
//	go func() {
//
//		q.Write = true
//		q.Read = true
//
//		for {
//			select {
//			case pq := <-q.work:
//				fmt.Println("queue len", len(q.work))
//				fmt.Println(pq)
//				//q.executeProof(pq)
//			case status := <-q.status:
//				switch status {
//				case CollectChallengesChannelStop:
//					q.Read = false
//				case CommitChallengesChannelStop:
//					q.Write = false
//				}
//				return
//			default:
//				if !q.Write {
//					q.Read = false
//					fmt.Println("none channel")
//				}
//			}
//
//		}
//	}()
//
//	q.logger.Info("Restart channel for work with proof")
//}
//
//func (q *DispatchChallengesQueue) CollectChannelStatus() bool {
//	return q.Read
//}
//
//func (q *DispatchChallengesQueue) SubmitChallengesPoolsNumber() int {
//	return len(q.submitChallengesPools)
//}
//
//func (q *DispatchChallengesQueue) SubmitChallengesPoolsIsEmpty() bool {
//	return q.SubmitChallengesPoolsNumber() == 0
//}
//
//func (q *DispatchChallengesQueue) DropChallengeInSubmitChallengesPools(challengeHash common.Hash) error {
//	delete(q.submitChallengesPools, challengeHash)
//	return database.DelChallenge(q.db, challengeHash.Bytes())
//}
//
//func (q *DispatchChallengesQueue) CommitChallengesPoolsNumber() int {
//	return len(q.commitChallengesPools)
//}
//
//func (q *DispatchChallengesQueue) CommitChallengesPoolsIsEmpty() bool {
//	return q.CommitChallengesPoolsNumber() == 0
//}
//
//func (q *DispatchChallengesQueue) DropChallengeInCommitChallengesPools(challengeHash common.Hash) error {
//	delete(q.commitChallengesPools, challengeHash)
//	return database.DelChallenge(q.db, challengeHash.Bytes())
//}
//
//func (q *DispatchChallengesQueue) ReloadWaitTime(instantWaitTime *utils.InstantWaitTime) {
//	q.waitTime = instantWaitTime
//}
//
//func (q *DispatchChallengesQueue) calSubmitWait(isSource bool, chainId uint64) int64 {
//	switch chainId {
//	case core.MainnetChainId, core.SepoliaChainId:
//		if isSource {
//			return q.waitTime.EthereumSubmitWait
//		} else {
//			return q.waitTime.ZkSyncSubmitWait
//		}
//	case core.ZkSyncEraMainnetChainId, core.ZkSyncEraSepoliaChainId:
//		if isSource {
//			return q.waitTime.ZkSyncSubmitWait
//		} else {
//			return q.waitTime.EthereumSubmitWait
//		}
//	default:
//		return 0
//	}
//}
//
//func (q *DispatchChallengesQueue) calCommitWait(isSource bool, chainId uint64) int64 {
//	switch chainId {
//	case core.MainnetChainId, core.SepoliaChainId:
//		if isSource {
//			return q.waitTime.EthereumCommitWait
//		} else {
//			return q.waitTime.ZkSyncCommitWait
//		}
//	case core.ZkSyncEraMainnetChainId, core.ZkSyncEraSepoliaChainId:
//		if isSource {
//			return q.waitTime.ZkSyncCommitWait
//		} else {
//			return q.waitTime.EthereumCommitWait
//		}
//	default:
//		return 0
//	}
//}
