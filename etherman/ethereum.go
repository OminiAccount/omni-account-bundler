package etherman

import (
	"context"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
	"strings"
	"time"
)

type EtherMan struct {
	ctx          context.Context
	cancel       func()
	cfg          Config
	chainsClient map[uint64]*EthereumClient
	db           *PostgresStorage
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(ctx context.Context, cfg Config, pg *pgxpool.Pool) (*EtherMan, error) {
	etherCtx, cancel := context.WithCancel(ctx)
	em := &EtherMan{
		ctx:          etherCtx,
		cancel:       cancel,
		cfg:          cfg,
		chainsClient: make(map[uint64]*EthereumClient, 200),
		db:           NewPostgresStorage(pg),
	}
	var nw = []Network{
		cfg.VizingNetwork,
		cfg.ArbitNetwork,
		cfg.BaseNetwork,
		cfg.OpNetwork,
	}
	for _, n := range nw {
		if n.ChainId < 1 {
			continue
		}
		cli, err := NewClient(n)
		if err != nil {
			return nil, err
		}
		opts, _, err := em.LoadAuthFromKeyStore(n.PrivateKeys.Path, n.PrivateKeys.Password, uint64(n.ChainId))
		if err != nil {
			return nil, err
		}
		cli.sender = opts.From
		cli.opAuth = *opts
		em.chainsClient[n.ChainId] = cli
	}
	return em, nil
}

func (ether *EtherMan) GetChains() map[uint64]*EthereumClient {
	return ether.chainsClient
}

func (ether *EtherMan) GetChainCli(c uint64) *EthereumClient {
	return ether.chainsClient[c]
}

func (ether *EtherMan) Start() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				tempUser.Range(func(key, value any) bool {
					lu := value.(lockUser)
					if time.Since(lu.timeAt) > time.Minute {
						tempUser.Delete(key)
					}
					return true
				})
				list := ether.db.GetFailedSalts(ether.ctx, nil)
				if list == nil || len(list) < 1 {
					continue
				}
				for _, ufs := range list {
					cids := strings.Split(ufs.FailedSalt, ",")
					if len(cids) < 1 {
						continue
					}
					for _, cid := range cids {
						nid, err := strconv.ParseUint(cid, 10, 64)
						if err != nil {
							continue
						}
						pd := PendingData{
							User:    common.HexToAddress(ufs.User),
							ChainID: nid,
							Salt:    ufs.Salt,
						}
						ok, err := ether.db.IsExistChainByOwner(ether.ctx, ufs.User, pd.ChainID, nil)
						if err != nil {
							log.Errorf("IsExistChainByOwner err: %+v", err)
							continue
						}
						if ok {
							ether.saveFailedCreateAA(pd, true)
							continue
						}
						_, _ = ether.pendingCreateAA(pd)
					}
				}
			case <-ether.ctx.Done():
				log.Info("EtherMan exit...")
				return
			}
		}
	}()
}

func (ether *EtherMan) Stop() {
	log.Info("EtherMan stop")
	ether.cancel()
}
