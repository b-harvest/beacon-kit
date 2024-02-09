// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package flags

const (
	// Beacon Kit Root Flag.
	beaconKitRoot = "beacon-kit."

	// Execution Client.
	engineRoot              = beaconKitRoot + "engine."
	RPCDialURL              = engineRoot + "rpc-dial-url"
	RPCRetries              = engineRoot + "rpc-retries"
	RPCTimeout              = engineRoot + "rpc-timeout"
	RPCStartupCheckInterval = engineRoot + "rpc-startup-check-interval"
	RPCHealthCheckInteval   = engineRoot + "rpc-health-check-interval"
	RPCJWTRefreshInterval   = engineRoot + "rpc-jwt-refresh-interval"
	JWTSecretPath           = engineRoot + "jwt-secret-path"
	RequiredChainID         = engineRoot + "required-chain-id"

	// Beacon Config.
	beaconConfigRoot = beaconKitRoot + "beacon-config."

	// Fork Config.
	forkRoot       = beaconConfigRoot + "forks."
	DenebForkEpoch = forkRoot + "deneb-fork-epoch"

	// Validator Config.
	validator             = beaconKitRoot + "validator"
	SuggestedFeeRecipient = validator + "suggested-fee-recipient"
	Graffiti              = validator + "graffiti"
	PrepareAllPayloads    = validator + "prepare-all-payloads"

	// ABCI Config.
	abciRoot            = beaconKitRoot + "abci."
	BeaconBlockPosition = abciRoot + "beacon-block-proposal-position"
)