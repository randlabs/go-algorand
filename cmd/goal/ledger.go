// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	ledgerCmd.AddCommand(supplyCmd)
	ledgerCmd.AddCommand(accountsCmd)
}

var ledgerCmd = &cobra.Command{
	Use:   "ledger",
	Short: "Access ledger-related details",
	Long:  "Access ledger-related details",
	Args:  validateNoPosArgsFn,
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments passed, we should fallback to help
		cmd.HelpFunc()(cmd, args)
	},
}

var supplyCmd = &cobra.Command{
	Use:   "supply",
	Short: "Show ledger token supply",
	Long:  "Show ledger token supply. All units are in microAlgos.",
	Args:  validateNoPosArgsFn,
	Run: func(cmd *cobra.Command, _ []string) {
		dataDir := ensureSingleDataDir()
		response, err := ensureAlgodClient(dataDir).LedgerSupply()
		if err != nil {
			reportErrorf(errorRequestFail, err)
		}

		fmt.Printf("Round: %v microAlgos\nTotal Money: %v microAlgos\nOnline Money: %v microAlgos\nAll Money: %v microAlgos\n", response.Round, response.TotalMoney, response.OnlineMoney, response.AllMoney)
	},
}

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Show ledger accounts (account, status, balance)",
	Long:  "Show status of all accounts in ledger, and total supply. All units are in microAlgos.",
	Args:  validateNoPosArgsFn,
	Run: func(cmd *cobra.Command, _ []string) {
		dataDir := ensureSingleDataDir()
		response, err := ensureAlgodClient(dataDir).LedgerAccounts()
		if err != nil {
			reportErrorf(errorRequestFail, err)
		}

		if len(response) == 0 {
			reportErrorln(errorNoAccounts)
		}

		round := response[0].Round
		fmt.Printf("Round: %v\n", round)

		total := uint64(0)
		for _, acct := range response {
			fmt.Printf("Acct: %s\t%-20s\tMoney: %20d\tPending Rewards: %15d\tw/o Rewards: %20d\tRewards: %15d\n", acct.Address, acct.Status, acct.Amount, acct.PendingRewards, acct.AmountWithoutPendingRewards, acct.Rewards)
			total += acct.Amount
		}
		fmt.Printf("Round: %d  |  Total Money: %d\n", round, total)
	},
}
