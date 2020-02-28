// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-wtc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-wtc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package wtcclient

import "github.com/wtc/go-wtc"

// Verify that Client implements the wtc interfaces.
var (
	_ = wtc.ChainReader(&Client{})
	_ = wtc.TransactionReader(&Client{})
	_ = wtc.ChainStateReader(&Client{})
	_ = wtc.ChainSyncReader(&Client{})
	_ = wtc.ContractCaller(&Client{})
	_ = wtc.GasEstimator(&Client{})
	_ = wtc.GasPricer(&Client{})
	_ = wtc.LogFilterer(&Client{})
	_ = wtc.PendingStateReader(&Client{})
	// _ = wtc.PendingStateEventer(&Client{})
	_ = wtc.PendingContractCaller(&Client{})
)
