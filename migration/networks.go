// Copyright © 2017 Swarm Market <info@swarm.market>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migration

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
)

var networks map[string]string

func InitNetworks() {
	networks = viper.GetStringMapString("networks")
	fmt.Println(networks)
}

type Network struct {
	rpc_client *rpc.Client
	client     *ethclient.Client
	accounts   []common.Address
}

func Dial(name string) (*Network, error) {
	if url, ok := networks[name]; ok {
		rpc_client, err := rpc.Dial(url)

		if err != nil {
			return nil, err
		}

		client := ethclient.NewClient(rpc_client)

		var accounts []common.Address
		rpc_client.Call(&accounts, "eth_accounts")
		fmt.Println(accounts, len(accounts))

		ret := &Network{rpc_client, client, accounts}

		return ret, nil
	}

	return nil, errors.New("No such network " + name)
}

func (n *Network) Backend() bind.ContractBackend {
	return n.client
}

func (n *Network) NewTransactor(account uint) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   n.accounts[account],
		Signer: nil,
	}
}
