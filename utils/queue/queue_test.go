package queue

//
//import (
//	"fmt"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/OAB/database"
//	"github.com/orbiter/spv-go/processor/utils"
//	"github.com/orbiter/spv-go/sdk/core"
//	"testing"
//)
//
//func TestEnqueue(t *testing.T) {
//	db, _ := database.NewDB("root:123456789@tcp(127.0.0.1:3306)/spv?charset=utf8mb4&parseTime=True&loc=Local")
//	queue := NewQueue()
//
//	str := "0x73bafed72039eeacfba7b6c8225c4b04fb49d125a8c3a44195097669f9489381"
//	work1 := &ChallengeQueue{
//		ChallengeInput: &utils.ChallengeInput{
//			ChallengeHash: common.HexToHash(str),
//			Body: &core.ChallengeBodyInput{
//				Hash: common.HexToHash(str),
//			},
//			ZkPool: "testUrl",
//		},
//		Provider: nil,
//		DB:       db,
//	}
//	err := queue.Enqueue(work1)
//	if err != nil {
//		fmt.Println("Enqueue error", err)
//	}
//
//	fmt.Println("len", len(queue.work))
//	fmt.Println("commitChallengesPools len", len(queue.commitChallengesPools))
//
//	str1 := "0x73bafed72039eeacfba7b6c8225c4b04fb49d125a8c3a44195097669f9489381"
//
//	r := &utils.CommitChallengeInput{
//		ChallengeHash: common.HexToHash(str1),
//		Proof:         nil,
//		Status:        0,
//		Message:       "",
//	}
//
//	err = queue.Dequeue(r)
//	if err != nil {
//		fmt.Println("Dequeue error", err)
//	}
//	fmt.Println("commitChallengesPools len", len(queue.commitChallengesPools))
//}
