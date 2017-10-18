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

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/swarmdotmarket/perigord/migration"
	perigord "github.com/swarmdotmarket/perigord/perigord/cmd"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migration.InitNetworks()
		network, err := migration.Dial(viper.GetString("default_network"))
		if err != nil {
			perigord.Fatal(err)
		}

		if err := migration.RunMigrations(context.Background(), network); err != nil {
			perigord.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	migrateCmd.PersistentFlags().StringP("network", "n", "NAME", "network to run migrations on")

	viper.BindPFlag("default_network", migrateCmd.PersistentFlags().Lookup("network"))
	viper.SetDefault("default_network", "dev")
}
