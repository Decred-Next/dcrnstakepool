// Copyright (c) 2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package helpers

import (
	"github.com/Decred-Next/dcrnd/chaincfg/v8"
	"github.com/Decred-Next/dcrnd/dcrec/v8"
	"github.com/Decred-Next/dcrnd/dcrutil/version2/v8"
	"github.com/Decred-Next/dcrnd/hdkeychain/version2/v8"
)

const (
	// ExternalBranch is a helper value that needs to
	// match dcrwallet's udb.ExternalBranch
	ExternalBranch uint32 = 0
)

// DCRUtilAddressFromExtendedKey parses the public address of a hd extended key
// using a secp256k1 elliptic curve into a ECDSA public key, compresses it using
// ripemd160, and wraps it in a dcrutil AddressPubKeyHash in order to easily
// obtain its human readable formats. Returns an error upon a parsing error or
// if key is for the wrong network.
func DCRUtilAddressFromExtendedKey(key *hdkeychain.ExtendedKey, params *chaincfg.Params) (*dcrutil.AddressPubKeyHash, error) {
	ecPubKey, err := key.ECPubKey()
	if err != nil {
		return nil, err
	}
	return dcrutil.NewAddressPubKeyHash(dcrutil.Hash160(ecPubKey.Serialize()), params, dcrec.STEcdsaSecp256k1)
}
