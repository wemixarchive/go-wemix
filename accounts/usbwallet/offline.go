// Copyright 2019 The go-ethereum / go-metadium / go-wemix Authors
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

package usbwallet

import (
	"errors"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/karalabe/usb"
)

// a simple wrapper for offline use, mainly for Close()
type offlineWallet struct {
	driver driver
	info   usb.DeviceInfo
	device usb.Device
	log    log.Logger
}

// dummy Hub structures for offline wallet use
var (
	offlineHubs []*Hub
)

func init() {
	// nano ledger
	offlineHubs = append(offlineHubs, &Hub{
		scheme:     LedgerScheme,
		vendorID:   0x2c97,
		productIDs: []uint16{0x0000, 0x0001},
		usageID:    0xffa0,
		endpointID: 0,
	})
	// trezor 1
	offlineHubs = append(offlineHubs, &Hub{
		scheme:     TrezorScheme,
		vendorID:   0x534c,
		productIDs: []uint16{0x0001},
		usageID:    0xff00,
		endpointID: 0,
	})
	// trezor 2, this doesn't seem to work yet.
	offlineHubs = append(offlineHubs, &Hub{
		scheme:     TrezorScheme,
		vendorID:   0x1209,
		productIDs: []uint16{0x53c0, 0x53c1},
		usageID:    0xf1d0,
		endpointID: -1,
	})
}

func (w *offlineWallet) Status() (string, error) {
	return w.driver.Status()
}

func (w *offlineWallet) Open(device io.ReadWriter, passphrase string) error {
	return w.driver.Open(device, passphrase)
}

func (w *offlineWallet) Close() error {
	if w.device == nil {
		return nil
	}
	w.driver.Close()
	w.device.Close()
	w.device = nil
	return nil
}

func (w *offlineWallet) Heartbeat() error {
	return w.driver.Heartbeat()
}

func (w *offlineWallet) Derive(path accounts.DerivationPath) (common.Address, error) {
	return w.driver.Derive(path)
}

func (w *offlineWallet) SignTx(path accounts.DerivationPath, tx *types.Transaction, chainID *big.Int) (common.Address, *types.Transaction, error) {
	return w.driver.SignTx(path, tx, chainID)
}

// ListDevices returns the array of ledger or trezor device paths
// for mostly informational use
func ListDevices(scheme string) []string {
	var devices []string

	// triple loop!
	for _, hub := range offlineHubs {
		if len(scheme) > 0 && hub.scheme != scheme {
			continue
		}
		infos, err := usb.Enumerate(hub.vendorID, 0)
		if err != nil {
			break
		}
		for _, info := range infos {
			for _, id := range hub.productIDs {
				if info.ProductID == id && (info.UsagePage == hub.usageID || info.Interface == hub.endpointID) {
					devices = append(devices, info.Path)
					break
				}
			}
		}
	}
	return devices
}

// OpenOffline opens usb wallet for offline use.
// If path is not specified, the first device is used.
// Note that Trezor 1 after firmware upgrade doesn't work any more.
func OpenOffline(scheme, path string) (interface{}, error) {
	if scheme != LedgerScheme && scheme != TrezorScheme {
		return nil, errors.New("Not Supported")
	}
	if path == "" {
		paths := ListDevices(scheme)
		if len(paths) == 0 {
			return nil, errors.New("Not Found")
		}
		path = paths[0]
	}

	for _, hub := range offlineHubs {
		if hub.scheme != scheme {
			continue
		}
		infos, err := usb.Enumerate(hub.vendorID, 0)
		if err != nil {
			break
		}
		for _, info := range infos {
			if info.Path == path {
				if scheme == TrezorScheme && info.Interface == -1 {
					// Trezor 2. It's going to hang. Stop here.
					return nil, errors.New("Trezor 2 doesn't work yet. You shouldn't have upgraded the firmware.")
				}

				logger := log.New("url", accounts.URL{Scheme: scheme, Path: path})
				var drv driver
				switch scheme {
				case LedgerScheme:
					drv = newLedgerDriver(logger)
				case TrezorScheme:
					drv = newTrezorDriver(logger)
				}
				device, err := info.Open()
				if err != nil {
					return nil, err
				}
				err = drv.Open(device, "")
				if err != nil {
					return nil, err
				}
				w := &offlineWallet{
					driver: drv,
					info:   info,
					device: device,
					log:    logger,
				}
				return w, nil
			}
		}
	}

	return nil, errors.New("Not Found")
}

// EOF
