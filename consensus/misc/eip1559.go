// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package misc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// VerifyEip1559Header verifies some header attributes which were changed in EIP-1559,
// - gas limit check
// - basefee check
func VerifyEip1559Header(config *params.ChainConfig, parent, header *types.Header) error {
	// Verify that the gas limit remains within allowed bounds
	parentGasLimit := parent.GasLimit
	if wemixminer.IsPoW() {
		if !config.IsLondon(parent.Number) {
			parentGasLimit = parent.GasLimit * params.ElasticityMultiplier
		}
	} else {
		_, _, _, _, gasTargetPercentage, err := wemixminer.GetBlockBuildParameters(parent.Number)
		if err == wemixminer.ErrNotInitialized {
			return nil
		}
		if !config.IsLondon(parent.Number) {
			parentGasLimit = parent.GasLimit * uint64(gasTargetPercentage) / 100
		}
	}
	if err := VerifyGaslimit(parentGasLimit, header.GasLimit); err != nil {
		return err
	}
	// Verify the header is not malformed
	if header.BaseFee == nil {
		return fmt.Errorf("header is missing baseFee")
	}
	// Verify the baseFee is correct based on the parent header.
	expectedBaseFee := CalcBaseFee(config, parent)
	if header.BaseFee.Cmp(expectedBaseFee) != 0 {
		return fmt.Errorf("invalid baseFee: have %s, want %s, parentBaseFee %s, parentGasUsed %d",
			header.BaseFee, expectedBaseFee, parent.BaseFee, parent.GasUsed)
	}
	return nil
}

// CalcBaseFee calculates the basefee of the header.
func CalcBaseFee(config *params.ChainConfig, parent *types.Header) *big.Int {
	// If the current block is the first EIP-1559 block, return the InitialBaseFee.
	if !config.IsLondon(parent.Number) {
		return new(big.Int).SetInt64(params.InitialBaseFee)
	}

	var (
		parentGasTarget          = parent.GasLimit / params.ElasticityMultiplier
		parentGasTargetBig       = new(big.Int).SetUint64(parentGasTarget)
		baseFeeChangeDenominator = new(big.Int).SetUint64(params.BaseFeeChangeDenominator)
		baseFeeChangeRate        *big.Int
		maxBaseFee               *big.Int
	)
	if !wemixminer.IsPoW() {
		// NB: in Wemix both elasticityMultiplier & baseFeeChangeDenominator are percentage numbers
		_, maxBaseFeeGov, _, baseFeeMaxChangeRate, gasTargetPercentage, err := wemixminer.GetBlockBuildParameters(parent.Number)
		if err == wemixminer.ErrNotInitialized {
			return new(big.Int).Set(parent.BaseFee)
		}
		parentGasTarget = parent.GasLimit * uint64(gasTargetPercentage) / 100
		parentGasTargetBig = new(big.Int).SetUint64(parentGasTarget)
		if parentGasTargetBig.Cmp(common.Big0) == 0 {
			parentGasTargetBig = new(big.Int).SetUint64(parentGasTarget)
		}
		baseFeeChangeRate = new(big.Int).SetInt64(baseFeeMaxChangeRate)
		maxBaseFee = maxBaseFeeGov
	}
	// If the parent gasUsed is the same as the target, the baseFee remains unchanged.
	if parent.GasUsed == parentGasTarget {
		return new(big.Int).Set(parent.BaseFee)
	}
	if parent.GasUsed > parentGasTarget {
		// If the parent block used more gas than its target, the baseFee should increase.
		gasUsedDelta := new(big.Int).SetUint64(parent.GasUsed - parentGasTarget)
		x := new(big.Int).Mul(parent.BaseFee, gasUsedDelta)
		y := new(big.Int)
		if parentGasTargetBig.Cmp(common.Big0) == 0 {
			y = x
		} else {
			y = x.Div(x, parentGasTargetBig)
		}
		var baseFeeDelta *big.Int
		if wemixminer.IsPoW() {
			baseFeeDelta = math.BigMax(
				x.Div(y, baseFeeChangeDenominator),
				common.Big1,
			)

			return x.Add(parent.BaseFee, baseFeeDelta)
		} else {
			baseFeeDelta = math.BigMax(
				x.Div(y.Mul(y, baseFeeChangeRate), big.NewInt(100)),
				common.Big1,
			)

			return math.BigMin(x.Add(parent.BaseFee, baseFeeDelta), maxBaseFee)
		}
	} else {
		// Otherwise if the parent block used less gas than its target, the baseFee should decrease.
		gasUsedDelta := new(big.Int).SetUint64(parentGasTarget - parent.GasUsed)
		x := new(big.Int).Mul(parent.BaseFee, gasUsedDelta)
		y := x.Div(x, parentGasTargetBig)
		var baseFeeDelta *big.Int
		if wemixminer.IsPoW() {
			baseFeeDelta = x.Div(y, baseFeeChangeDenominator)
		} else {
			baseFeeDelta = x.Div(y.Mul(y, baseFeeChangeRate), big.NewInt(100))
			if baseFeeDelta.Cmp(common.Big0) == 0 && parent.BaseFee.Cmp(common.Big1) > 0 {
				baseFeeDelta.SetUint64(1)
			}
		}
		return math.BigMax(
			x.Sub(parent.BaseFee, baseFeeDelta),
			common.Big1,
		)
	}
}
