package database

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type ChallengesModel struct {
	gorm.Model
	Transaction   common.Hash   `gorm:"index;size:66"`
	ChallengeId   common.Hash   `gorm:"index;size:66"`
	ChallengeHash common.Hash   `gorm:"unique;size:66"`
	Source        bool          `gorm:"index"`
	Proof         hexutil.Bytes `gorm:"serializer:json"`
	Message       string
	Status        uint                  `gorm:"index"`                                    // 1 is `No proof generated`; 2 is `Proof generated`; 3 is `Proof generation failed`;
	IsDelete      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // use `1` `0`
}

type ChallengesDB interface {
	InsertChallenge(txProof *ChallengesModel) error
	QueryChallengeByTransactionHash(taskId common.Hash) (*ChallengesModel, error)
	UpdateChallengeByTransactionHash(txProof *ChallengesModel) error
	DeleteByIDUnscoped(transaction common.Hash) error
}

type challengeDB struct {
	gorm *gorm.DB
}

func NewChallengesProofDB(db *gorm.DB) ChallengesDB {
	db.AutoMigrate(&ChallengesModel{})
	return &challengeDB{gorm: db}
}

func (cd *challengeDB) InsertChallenge(model *ChallengesModel) error {
	result := cd.gorm.Create(model)
	return result.Error
}

func (cd *challengeDB) QueryChallengeByTransactionHash(hash common.Hash) (*ChallengesModel, error) {
	var challengesModel *ChallengesModel
	result := cd.gorm.Where("transaction = ?", hash).Take(&challengesModel)
	return challengesModel, result.Error
}

func (cd *challengeDB) UpdateChallengeByTransactionHash(model *ChallengesModel) error {
	result := cd.gorm.Model(&model).Updates(model)
	return result.Error
}

// DeleteByIDUnscoped will permanently delete the data from the database
func (cd *challengeDB) DeleteByIDUnscoped(transaction common.Hash) error {
	txProofQuery, err := cd.QueryChallengeByTransactionHash(transaction)
	if err != nil {
		return err
	}
	result := cd.gorm.Unscoped().Delete(txProofQuery)
	return result.Error
}
