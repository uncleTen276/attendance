package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	attendance_record "github.com/uncleTen276/attendance.git/server/abi"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
)

func GetContractInstance() (*attendance_record.AttendanceRecord, *ethclient.Client, error) {
	cfg := configs.AppConfig()
	client, err := ethclient.Dial(cfg.Contract.Infura_Url)
	if err != nil {
		return nil, nil, err
	}

	contractAddress := common.HexToAddress(cfg.Contract.Contract_Address)
	contract, err := attendance_record.NewAttendanceRecord(contractAddress, client)
	if err != nil {
		return nil, nil, err
	}

	return contract, client, nil
}

func GetAuth(chainId *big.Int) (*bind.TransactOpts, error) {
	cfg := configs.AppConfig()
	privateKey, err := crypto.HexToECDSA(cfg.Contract.Private_Key)
	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactorWithChainID(privateKey, chainId)
}
