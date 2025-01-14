package evm

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vizualni/whoops"
)

var (
	log1 = types.Log{
		BlockNumber: 1,
	}
	log2 = types.Log{
		BlockNumber: 2,
	}
	log3 = types.Log{
		BlockNumber: 3,
	}
	log4 = types.Log{
		BlockNumber: 4,
	}
	log5 = types.Log{
		BlockNumber: 5,
	}
)

func TestReadingStoredContracts(t *testing.T) {
	t.Run("it successfully reads the hello.json contract", func(t *testing.T) {
		c := StoredContracts()
		require.GreaterOrEqual(t, len(c), 1)
		require.Contains(t, c, "hello")
	})
}

func TestExecutingSmartContract(t *testing.T) {
	require.Contains(t, StoredContracts(), "simple")
	// ctx := context.Background()
	cryptokey, err := crypto.HexToECDSA(privateKeyBob)
	require.NoError(t, err)
	fakeErr := whoops.String("oh no")
	for _, tt := range []struct {
		name        string
		setup       func(t *testing.T, args *executeSmartContractIn)
		expectedErr error
	}{
		{
			name: "happy path",
			setup: func(t *testing.T, args *executeSmartContractIn) {
				ethMock := newMockEthClienter(t)

				ethMock.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(333), nil)

				ethMock.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(444), nil)

				ethMock.On("PendingCodeAt", mock.Anything, args.contract).Return([]byte("a"), nil)

				ethMock.On("EstimateGas", mock.Anything, mock.Anything).Return(uint64(222), nil)

				ethMock.On("SendTransaction", mock.Anything, mock.Anything).Return(nil)

				args.ethClient = ethMock
			},
		},
		{
			name:        "nonce returns an error and it returns error back",
			expectedErr: fakeErr,
			setup: func(t *testing.T, args *executeSmartContractIn) {
				ethMock := newMockEthClienter(t)

				ethMock.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), fakeErr)

				args.ethClient = ethMock
			},
		},
		{
			name:        "sending transaction returns and error and that error is sent back",
			expectedErr: fakeErr,
			setup: func(t *testing.T, args *executeSmartContractIn) {
				ethMock := newMockEthClienter(t)

				ethMock.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(333), nil)

				ethMock.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(444), nil)

				ethMock.On("PendingCodeAt", mock.Anything, args.contract).Return([]byte("a"), nil)

				ethMock.On("EstimateGas", mock.Anything, mock.Anything).Return(uint64(222), nil)

				ethMock.On("SendTransaction", mock.Anything, mock.Anything).Return(fakeErr)

				args.ethClient = ethMock
			},
		},
		{
			name:        "gas estimation returns an error and it returns error back",
			expectedErr: fakeErr,
			setup: func(t *testing.T, args *executeSmartContractIn) {
				ethMock := newMockEthClienter(t)

				ethMock.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(333), nil)

				ethMock.On("SuggestGasPrice", mock.Anything).Return(nil, fakeErr)

				args.ethClient = ethMock
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore.NewKeyStore(t.TempDir(), keystore.StandardScryptN, keystore.StandardScryptP)
			require.NoError(t, err)
			acc, err := ks.ImportECDSA(cryptokey, "bla")
			require.NoError(t, err)
			ks.Unlock(acc, "bla")
			contract := StoredContracts()["simple"]
			args := executeSmartContractIn{
				chainID:       big.NewInt(1337),
				gasAdjustment: 1.0,
				contract:      common.HexToAddress("0xBABA"),
				signingAddr:   acc.Address,
				abi:           contract.ABI,
				method:        "store",
				arguments:     []any{big.NewInt(123)},
				keystore:      ks,
			}

			tt.setup(t, &args)

			_, err = callSmartContract(
				context.Background(),
				args,
			)

			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestFilterLogs(t *testing.T) {
	fakeErr := whoops.String("fake error")

	hashPtr := func(h common.Hash) *common.Hash {
		return &h
	}

	for _, tt := range []struct {
		name string

		filterQuery ethereum.FilterQuery
		blockHeight *big.Int
		callback    func(*testing.T) func([]types.Log) bool

		setup func(t *testing.T) *mockEthClientToFilterLogs

		expErr error
		expRes []types.Log
	}{
		{
			name: "unable to get the current block height returns an error",
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				srv := newMockEthClientToFilterLogs(t)
				srv.On("HeaderByNumber", mock.Anything, (*big.Int)(nil)).Return(nil, fakeErr)
				return srv
			},
			expErr: fakeErr,
		},
		{
			name: "if it's looking at exact block, then it will not change any arguments",
			filterQuery: ethereum.FilterQuery{
				BlockHash: hashPtr(common.HexToHash("abc")),
			},
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				srv := newMockEthClientToFilterLogs(t)
				srv.On("HeaderByNumber", mock.Anything, mock.Anything).Return(&types.Header{
					Number: big.NewInt(134),
				}, nil)
				srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
					BlockHash: hashPtr(common.HexToHash("abc")),
				}).Return(nil, nil)
				return srv
			},
		},
		{
			name:        "if BlockHash is nil, and ToBlock and/or FromBlocks are also nil, then they are changed",
			filterQuery: ethereum.FilterQuery{},
			blockHeight: big.NewInt(134),
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				srv := newMockEthClientToFilterLogs(t)
				srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
					ToBlock:   big.NewInt(134),
					FromBlock: big.NewInt(0),
				}).Return(nil, nil)
				return srv
			},
		},
		{
			name:        "if there are results, then callback is called",
			filterQuery: ethereum.FilterQuery{},
			blockHeight: big.NewInt(134),
			callback: func(t *testing.T) func([]types.Log) bool {
				return func(logs []types.Log) bool {
					require.Empty(t, logs)
					return true
				}
			},
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				srv := newMockEthClientToFilterLogs(t)
				srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
					ToBlock:   big.NewInt(134),
					FromBlock: big.NewInt(0),
				}).Return(nil, nil)
				return srv
			},
		},
		{
			name:        "if there are more than 10000 results, then it calls it recursively",
			filterQuery: ethereum.FilterQuery{},
			blockHeight: big.NewInt(8),
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				results := [][]types.Log{
					{log1, log2, log3},
					{log4, log5},
				}
				srv := newMockEthClientToFilterLogs(t)

				srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
					ToBlock:   big.NewInt(8),
					FromBlock: big.NewInt(0),
				}).Times(1).Return(nil, whoops.String("query returned more than 10000 results"))
				callResults := func(from, to int64, index int) {
					srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
						FromBlock: big.NewInt(from),
						ToBlock:   big.NewInt(to),
					}).Times(1).Return(results[index], nil)
				}

				callResults(0, 4, 0)
				callResults(5, 8, 1)
				return srv
			},
			expRes: []types.Log{log1, log2, log3, log4, log5},
		},
		{
			name:        "any other error is returned",
			filterQuery: ethereum.FilterQuery{},
			blockHeight: big.NewInt(8),
			setup: func(t *testing.T) *mockEthClientToFilterLogs {
				srv := newMockEthClientToFilterLogs(t)
				srv.On("FilterLogs", mock.Anything, ethereum.FilterQuery{
					ToBlock:   big.NewInt(8),
					FromBlock: big.NewInt(0),
				}).Return(nil, fakeErr)
				return srv
			},
			expErr: fakeErr,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ethClienter := tt.setup(t)

			var res []types.Log
			defaultCallback := func(logs []types.Log) bool {
				res = append(res, logs...)
				return true
			}

			var fnc func([]types.Log) bool
			if tt.callback != nil {
				fn := tt.callback(t)
				fnc = func(logs []types.Log) bool {
					defaultCallback(logs)
					return fn(logs)
				}
			} else {
				fnc = defaultCallback
			}

			_, err := filterLogs(ctx, ethClienter, tt.filterQuery, tt.blockHeight, fnc)

			require.ErrorIs(t, err, tt.expErr)
			require.Equal(t, tt.expRes, res)
		})
	}
}
