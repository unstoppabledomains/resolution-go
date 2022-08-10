package resolution

import (
	"encoding/json"
	"fmt"
	"strings"
)

// supportedKeys struct of supported keys
type supportedKeys map[string]struct {
	DeprecatedKeyName string
	Deprecated        bool
	ValidationRegex   string
}

const emailKey = "whois.email.value"

var ipfsKeys = []string{"dweb.ipfs.hash", "ipfs.html.value"}
var redirectUrlKeys = []string{"browser.redirect_url", "ipfs.redirect_domain.value"}

// buildCryptoKey returns raw key for crypto currency which is used to query blockchain
func buildCryptoKey(ticker string) (string, error) {
	var key strings.Builder
	_, err := fmt.Fprintf(&key, "crypto.%s.address", strings.ToUpper(ticker))
	if err != nil {
		return "", err
	}
	return key.String(), nil
}

// buildCryptoKeyVersion returns raw key for multi-chain currency which is used to query blockchain
func buildCryptoKeyVersion(ticker string, version string) (string, error) {
	var key strings.Builder
	_, err := fmt.Fprintf(&key, "crypto.%s.version.%s.address", strings.ToUpper(ticker), strings.ToUpper(version))
	if err != nil {
		return "", err
	}
	return key.String(), nil
}

// returnFirstNonEmpty returns first not empty elements from provided records and keys order
func returnFirstNonEmpty(records map[string]string, keysSequence []string) string {
	for _, key := range keysSequence {
		if records[key] != "" {
			return records[key]
		}
	}

	return ""
}

// newSupportedKeys returns supportedKeys
func newSupportedKeys() (supportedKeys, error) {
	var keysObject struct {
		Keys supportedKeys
	}
	err := json.Unmarshal(supportedKeysJSON, &keysObject)
	if err != nil {
		return nil, err
	}
	return keysObject.Keys, nil
}

var supportedKeysJSON = []byte(`
{
  "version": "1.1.1",
  "keys": {
    "crypto.BTC.address": {
      "deprecatedKeyName": "BTC",
      "deprecated": false,
      "validationRegex": "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$"
    },
    "crypto.ETH.address": {
      "deprecatedKeyName": "ETH",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.ZIL.address": {
      "deprecatedKeyName": "ZIL",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$|^zil1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{38}$"
    },
    "crypto.LTC.address": {
      "deprecatedKeyName": "LTC",
      "deprecated": false,
      "validationRegex": "^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$|^ltc1[a-zA-HJ-NP-Z0-9]{25,39}$"
    },
    "crypto.ETC.address": {
      "deprecatedKeyName": "ETC",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.EQL.address": {
      "deprecatedKeyName": "EQL",
      "deprecated": false,
      "validationRegex": "^bnb[0-9a-z]{39}$"
    },
    "crypto.LINK.address": {
      "deprecatedKeyName": "LINK",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.USDC.address": {
      "deprecatedKeyName": "USDC",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.BAT.address": {
      "deprecatedKeyName": "BAT",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.REP.address": {
      "deprecatedKeyName": "REP",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.ZRX.address": {
      "deprecatedKeyName": "ZRX",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.DAI.address": {
      "deprecatedKeyName": "DAI",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.BCH.address": {
      "deprecatedKeyName": "BCH",
      "deprecated": false,
      "validationRegex": "^[13][a-km-zA-HJ-NP-Z1-9]{33}$|^((bitcoincash|bchreg|bchtest):)?(q|p)[a-z0-9]{41}$|^((BITCOINCASH:)?(Q|P)[A-Z0-9]{41})$"
    },
    "crypto.XMR.address": {
      "deprecatedKeyName": "XMR",
      "deprecated": false,
      "validationRegex": "^[48]{1}[0-9AB][1-9A-HJ-NP-Za-km-z]{93}$"
    },
    "crypto.DASH.address": {
      "deprecatedKeyName": "DASH",
      "deprecated": false,
      "validationRegex": "^X[1-9A-HJ-NP-Za-km-z]{33}$"
    },
    "crypto.NEO.address": {
      "deprecatedKeyName": "NEO",
      "deprecated": false,
      "validationRegex": "^A[0-9a-zA-Z]{33}$"
    },
    "crypto.SWTH.address": {
      "deprecatedKeyName": "SWTH",
      "deprecated": false,
      "validationRegex": "^A[0-9a-zA-Z]{33}$"
    },
    "crypto.DOGE.address": {
      "deprecatedKeyName": "DOGE",
      "deprecated": false,
      "validationRegex": "^D[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}$"
    },
    "crypto.XRP.address": {
      "deprecatedKeyName": "XRP",
      "deprecated": false,
      "validationRegex": "^r[1-9a-km-zA-HJ-NP-Z]{24,34}$"
    },
    "crypto.ZEC.address": {
      "deprecatedKeyName": "ZEC",
      "deprecated": false,
      "validationRegex": "^z([a-zA-Z0-9]){94}$|^zs1([a-zA-Z0-9]){75}$|^t([a-zA-Z0-9]){34}$"
    },
    "crypto.ADA.address": {
      "deprecatedKeyName": "ADA",
      "deprecated": false,
      "validationRegex": "^[1-9a-km-zA-HJ-NP-Z]{104}$|^A[1-9A-HJ-NP-Za-km-z]{58}$|^addr[0-9a-zA-Z]{99}$"
    },
    "crypto.EOS.address": {
      "deprecatedKeyName": "EOS",
      "deprecated": false,
      "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$"
    },
    "crypto.XLM.address": {
      "deprecatedKeyName": "XLM",
      "deprecated": false,
      "validationRegex": "^G[A-Z2-7]{55}$"
    },
    "crypto.BNB.address": {
      "deprecatedKeyName": "BNB",
      "deprecated": false,
      "validationRegex": "^bnb[0-9a-z]{39}$"
    },
    "crypto.BTG.address": {
      "deprecatedKeyName": "BTG",
      "deprecated": false,
      "validationRegex": "^[GA][a-km-zA-HJ-NP-Z1-9]{33}$"
    },
    "crypto.NANO.address": {
      "deprecatedKeyName": "NANO",
      "deprecated": false,
      "validationRegex": "^nano_[1-9a-z]{60}$"
    },
    "crypto.WAVES.address": {
      "deprecatedKeyName": "WAVES",
      "deprecated": false,
      "validationRegex": "^3[a-km-zA-HJ-NP-Z1-9]{34}$"
    },
    "crypto.KMD.address": {
      "deprecatedKeyName": "KMD",
      "deprecated": false,
      "validationRegex": "^R[a-km-zA-Z1-9]{33}$"
    },
    "crypto.AE.address": {
      "deprecatedKeyName": "AE",
      "deprecated": false,
      "validationRegex": "^ak_[a-km-zA-Z1-9]{48,52}$"
    },
    "crypto.RSK.address": {
      "deprecatedKeyName": "RSK",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.WAN.address": {
      "deprecatedKeyName": "WAN",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.STRAT.address": {
      "deprecatedKeyName": "STRAT",
      "deprecated": false,
      "validationRegex": "^S[a-km-zA-HJ-NP-Z1-9]{33}$"
    },
    "crypto.UBQ.address": {
      "deprecatedKeyName": "UBQ",
      "deprecated": false,
      "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
    },
    "crypto.XTZ.address": {
      "deprecatedKeyName": "XTZ",
      "deprecated": false,
      "validationRegex": "^(tz|KT)[a-km-zA-HJ-NP-Z1-9]{34}$"
    },
    "crypto.IOTA.address": {
      "deprecatedKeyName": "IOTA",
      "deprecated": false,
      "validationRegex": "^[A-Z0-9]{90}$|^iota1[a-z0-9]{59}$"
    },
    "crypto.VET.address": {
      "deprecatedKeyName": "VET",
      "deprecated": false,
      "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
    },
    "crypto.QTUM.address": {
      "deprecatedKeyName": "QTUM",
      "deprecated": false,
      "validationRegex": "^Q[a-km-zA-HJ-NP-Z1-9]{33}$"
    },
    "crypto.ICX.address": {
      "deprecatedKeyName": "ICX",
      "deprecated": false,
      "validationRegex": "^[a-km-zA-HJ-NP-Z0-9]{42}$"
    },
    "crypto.DGB.address": {
      "deprecatedKeyName": "DGB",
      "deprecated": false,
      "validationRegex": "(^[a-km-zA-HJ-NP-Z1-9]{34}$)|(^[a-zA-Z1-9]{42}$)|(^dgb1[a-zA-Z0-9]{39}$)"
    },
    "crypto.XZC.address": {
      "deprecatedKeyName": "XZC",
      "deprecated": false,
      "validationRegex": "^[a-km-zA-HJ-NP-Z1-9]{34}$"
    },
    "crypto.BURST.address": {
      "deprecatedKeyName": "BURST",
      "deprecated": false,
      "validationRegex": "^BURST-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{5}"
    },
    "crypto.DCR.address": {
      "deprecatedKeyName": "DCR",
      "deprecated": false,
      "validationRegex": null
    },
    "crypto.XEM.address": {
      "deprecatedKeyName": "XEM",
      "deprecated": false,
      "validationRegex": "^N[ABCDEFGHIJKLMNOPQRSTUVWXYZ234567]{39}$"
    },
    "crypto.LSK.address": {
      "deprecatedKeyName": "LSK",
      "deprecated": false,
      "validationRegex": "^\\d{1,21}[L]$"
    },
    "crypto.ATOM.address": {
      "deprecatedKeyName": "ATOM",
      "deprecated": false,
      "validationRegex": "^(cosmos)1([qpzry9x8gf2tvdw0s3jn54khce6mua7l]+)$"
    },
    "crypto.ONG.address": {
      "deprecatedKeyName": "ONG",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.ONT.address": {
      "deprecatedKeyName": "ONT",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.SMART.address": {
      "deprecatedKeyName": "SMART",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.TPAY.address": {
      "deprecatedKeyName": "TPAY",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.GRS.address": {
      "deprecatedKeyName": "GRS",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.BSV.address": {
      "deprecatedKeyName": "BSV",
      "deprecated": false,
      "validationRegex": "^bitcoincash:[a-zA-Z0-9]{42}$"
    },
    "crypto.GAS.address": {
      "deprecatedKeyName": "GAS",
      "deprecated": false,
      "validationRegex": null
    },
    "crypto.TRX.address": {
      "deprecatedKeyName": "TRX",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.VTHO.address": {
      "deprecatedKeyName": "VTHO",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{42}$"
    },
    "crypto.BCD.address": {
      "deprecatedKeyName": "BCD",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.BTT.address": {
      "deprecatedKeyName": "BTT",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.KIN.address": {
      "deprecatedKeyName": "KIN",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{56}$"
    },
    "crypto.RVN.address": {
      "deprecatedKeyName": "RVN",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.ARK.address": {
      "deprecatedKeyName": "ARK",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.XVG.address": {
      "deprecatedKeyName": "XVG",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.ALGO.address": {
      "deprecatedKeyName": "ALGO",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{58}$"
    },
    "crypto.NEBL.address": {
      "deprecatedKeyName": "NEBL",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.XPM.address": {
      "deprecatedKeyName": "XPM",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.ONE.address": {
      "deprecatedKeyName": "ONE",
      "deprecated": false,
      "validationRegex": "^one[a-zA-Z0-9]{39}$"
    },
    "crypto.BNTY.address": {
      "deprecatedKeyName": "BNTY",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.CRO.address": {
      "deprecatedKeyName": "CRO",
      "deprecated": false,
      "validationRegex": "^0x[a-fA-F0-9]{40}$"
    },
    "crypto.TWT.address": {
      "deprecatedKeyName": "TWT",
      "deprecated": false,
      "validationRegex": "^bnb[0-9a-z]{39}$"
    },
    "crypto.SIERRA.address": {
      "deprecatedKeyName": "SIERRA",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{34}$"
    },
    "crypto.VSYS.address": {
      "deprecatedKeyName": "VSYS",
      "deprecated": false,
      "validationRegex": "^[a-zA-Z0-9]{35}$"
    },
    "crypto.HIVE.address": {
      "deprecatedKeyName": "HIVE",
      "validationRegex": "^(?!s*$).+",
      "deprecated": false
    },
    "crypto.HT.address": {
      "deprecatedKeyName": "HT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ENJ.address": {
      "deprecatedKeyName": "ENJ",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.YFI.address": {
      "deprecatedKeyName": "YFI",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MTA.address": {
      "deprecatedKeyName": "MTA",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.COMP.address": {
      "deprecatedKeyName": "COMP",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BAL.address": {
      "deprecatedKeyName": "BAL",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AMPL.address": {
      "deprecatedKeyName": "AMPL",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LEND.address": {
      "deprecatedKeyName": "LEND",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TLOS.address": {
      "deprecatedKeyName": "TLOS",
      "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$",
      "deprecated": false
    },
    "crypto.XDC.address": {
      "deprecatedKeyName": "XDC",
      "validationRegex": "^xdc[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XST.address": {
      "deprecatedKeyName": "XST",
      "validationRegex": "(?:RwxQ3jUs2BjKhseNX1em4msn2GyV5XAec[PQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]|RwxQ3jUs2BjKhseNX1em4msn2GyV5XAe[defghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]|RwxQ3jUs2BjKhseNX1em4msn2GyV5XA[fghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{2}|RwxQ3jUs2BjKhseNX1em4msn2GyV5X[BCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{3}|RwxQ3jUs2BjKhseNX1em4msn2GyV5[YZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{4}|RwxQ3jUs2BjKhseNX1em4msn2GyV[6789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{5}|RwxQ3jUs2BjKhseNX1em4msn2Gy[WXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{6}|RwxQ3jUs2BjKhseNX1em4msn2G[z][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{7}|RwxQ3jUs2BjKhseNX1em4msn2[HJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{8}|RwxQ3jUs2BjKhseNX1em4msn[3456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{9}|RwxQ3jUs2BjKhseNX1em4ms[opqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{10}|RwxQ3jUs2BjKhseNX1em4m[tuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{11}|RwxQ3jUs2BjKhseNX1em4[nopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{12}|RwxQ3jUs2BjKhseNX1em[56789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{13}|RwxQ3jUs2BjKhseNX1e[nopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{14}|RwxQ3jUs2BjKhseNX1[fghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{15}|RwxQ3jUs2BjKhseNX[23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{16}|RwxQ3jUs2BjKhseN[YZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{17}|RwxQ3jUs2BjKhse[PQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{18}|RwxQ3jUs2BjKhs[fghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{19}|RwxQ3jUs2BjKh[tuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{20}|RwxQ3jUs2BjK[ijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{21}|RwxQ3jUs2Bj[LMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{22}|RwxQ3jUs2B[kmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{23}|RwxQ3jUs2[CDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{24}|RwxQ3jUs[3456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{25}|RwxQ3jU[tuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{26}|RwxQ3j[VWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{27}|RwxQ3[kmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{28}|RwxQ[456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{29}|Rwx[RSTUVWXYZabcdefghijkmnopqrstuvwxyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{30}|Rw[yz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{31}|R[xyz][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{32}|S[123456789ABCDEFGHJKL][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{32}|SM[123456789ABCDEFGH][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{31}|SMJ11[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{29}|SMJ11[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{29}|SMJ12[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnop][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{28}|SMJ12q[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkm][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{27}|SMJ12qn[12345678][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{26}|SMJ12qn9[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghi][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{25}|SMJ12qn9j[123456789ABCDEFGHJKLM][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{24}|SMJ12qn9jN[123456789AB][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{23}|SMJ12qn9jNC[123456789AB][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{22}|SMJ12qn9jNCC[123456789ABCDEFGHJKLMNPQRSTUVW][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{21}|SMJ12qn9jNCCX[123456789ABCDEFGH][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{20}|SMJ12qn9jNCCXJ[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkm][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{19}|SMJ12qn9jNCCXJn[123456789ABCDEFGHJKLMNPQRS][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{18}|SMJ12qn9jNCCXJnT[123456789ABCDEFGHJKLMNPQRSTUVWX][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{17}|SMJ12qn9jNCCXJnTY[123456789ABCDEFGHJKLMNPQ][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{16}|SMJ12qn9jNCCXJnTYR[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxy][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{15}|SMJ12qn9jNCCXJnTYRz[1234][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{14}|SMJ12qn9jNCCXJnTYRz5[123456789ABCDEFGHJKLMNPQRSTUVWX][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{13}|SMJ12qn9jNCCXJnTYRz5Y[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrst][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{12}|SMJ12qn9jNCCXJnTYRz5Yu[12345678][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{11}|SMJ12qn9jNCCXJnTYRz5Yu9[123456789ABCDEFGHJKLMNPQRSTUVWXY][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{10}|SMJ12qn9jNCCXJnTYRz5Yu9Z[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcd][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{9}|SMJ12qn9jNCCXJnTYRz5Yu9Ze[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkm][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{8}|SMJ12qn9jNCCXJnTYRz5Yu9Zen[123456789ABCD][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{7}|SMJ12qn9jNCCXJnTYRz5Yu9ZenE[123456789ABCDEFGHJKLMNPQ][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{6}|SMJ12qn9jNCCXJnTYRz5Yu9ZenER[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkm][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{5}|SMJ12qn9jNCCXJnTYRz5Yu9ZenERn[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghij][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{4}|SMJ12qn9jNCCXJnTYRz5Yu9ZenERnk[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghij][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{3}|SMJ12qn9jNCCXJnTYRz5Yu9ZenERnkk[123456789ABCDEFGHJKLMNPQRST][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{2}|SMJ12qn9jNCCXJnTYRz5Yu9ZenERnkkU[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstu][123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]|SMJ12qn9jNCCXJnTYRz5Yu9ZenERnkkUv[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghi])",
      "deprecated": false
    },
    "crypto.STRAX.address": {
      "deprecatedKeyName": "STRAX",
      "validationRegex": "^X[a-km-zA-HJ-NP-Z1-9]{33}$",
      "deprecated": false
    },
    "crypto.SIGNA.address": {
      "deprecatedKeyName": "SIGNA",
      "validationRegex": "^S-((?=[A-Z2-9]{4})(?:[^IO]{4})-){3}(?=[A-Z2-9]{5})(?:[^IO]{5})$",
      "deprecated": false
    },
    "crypto.NIM.address": {
      "deprecatedKeyName": "NIM",
      "validationRegex": "^NQ[0-9]{2} ([A-Z0-9]{4} ){7}[A-Z0-9]{4}$",
      "deprecated": false
    },
    "crypto.ELA.version.ELA.address": {
      "deprecatedKeyName": "ELA_ELA",
      "validationRegex": "E[a-zA-HJ-NP-Z0-9]{33}",
      "deprecated": false
    },
    "crypto.ELA.version.ESC.address": {
      "deprecatedKeyName": "ELA_ESC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ELA.version.HRC20.address": {
      "deprecatedKeyName": "ELA_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ELA.version.ERC20.address": {
      "deprecatedKeyName": "ELA_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.USDT.version.ERC20.address": {
      "deprecatedKeyName": "USDT_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.USDT.version.TRON.address": {
      "deprecatedKeyName": "USDT_TRON",
      "validationRegex": "^[T][a-zA-HJ-NP-Z0-9]{33}$",
      "deprecated": false
    },
    "crypto.USDT.version.EOS.address": {
      "deprecatedKeyName": "USDT_EOS",
      "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$",
      "deprecated": false
    },
    "crypto.USDT.version.OMNI.address": {
      "deprecatedKeyName": "USDT_OMNI",
      "validationRegex": "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
      "deprecated": false
    },
    "crypto.FTM.version.ERC20.address": {
      "deprecatedKeyName": "FTM_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FTM.version.BEP2.address": {
      "deprecatedKeyName": "FTM_BEP2",
      "validationRegex": "^(bnb|tbnb)[a-zA-HJ-NP-Z0-9]{39}$",
      "deprecated": false
    },
    "crypto.FTM.version.OPERA.address": {
      "deprecatedKeyName": "FTM_OPERA",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FUSE.version.ERC20.address": {
      "deprecatedKeyName": "FUSE_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FUSE.version.FUSE.address": {
      "deprecatedKeyName": "FUSE_FUSE",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "social.payid.name": {
      "deprecatedKeyName": "payid",
      "validationRegex": "^[0-9a-zA-Z]+\\$[0-9a-zA-Z]+\\.[0-9a-zA-Z]+$",
      "deprecated": false
    },
    "social.picture.value": {
      "deprecatedKeyName": "picture",
      "validationRegex": null,
      "deprecated": false
    },
    "whois.email.value": {
      "deprecatedKeyName": "email",
      "validationRegex": "^[^@]+@[^\\.]+\\..+$",
      "deprecated": false
    },
    "whois.for_sale.value": {
      "deprecatedKeyName": "for_sale",
      "validationRegex": "(true)|(false)",
      "deprecated": false
    },
    "ipfs.html.value": {
      "deprecatedKeyName": "html",
      "validationRegex": ".{0,100}",
      "deprecated": false
    },
    "ipfs.redirect_domain.value": {
      "deprecatedKeyName": "redirect_domain",
      "validationRegex": ".{0,253}",
      "deprecated": false
    },
    "dweb.ipfs.hash": {
      "deprecatedKeyName": "dweb_hash",
      "validationRegex": ".{0,100}",
      "deprecated": false
    },
    "browser.redirect_url": {
      "deprecatedKeyName": "browser_redirect",
      "validationRegex": ".{0,253}",
      "deprecated": false
    },
    "browser.preferred_protocols": {
      "deprecatedKeyName": "browser_preferred_protocols",
      "validationRegex": null,
      "deprecated": false
    },
    "gundb.username.value": {
      "deprecatedKeyName": "gundb_username",
      "validationRegex": null,
      "deprecated": false
    },
    "gundb.public_key.value": {
      "deprecatedKeyName": "gundb_public_key",
      "validationRegex": null,
      "deprecated": false
    },
    "social.image.value": {
      "deprecatedKeyName": "image",
      "validationRegex": null,
      "deprecated": false
    },
    "social.twitter.username": {
      "deprecatedKeyName": "twitter_username",
      "validationRegex": null,
      "deprecated": false
    },
    "validation.social.twitter.username": {
      "deprecatedKeyName": "validation_twitter_username",
      "validationRegex": null,
      "deprecated": false
    },
    "forwarding.url": {
      "deprecatedKeyName": "forwarding_url",
      "validationRegex": "^(https?)://[^\\s/$.?#].[^\\s]*$",
      "deprecated": false
    },
    "dns.ttl": {
      "deprecatedKeyName": "dns_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.A": {
      "deprecatedKeyName": "dns_A",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.A.ttl": {
      "deprecatedKeyName": "dns_A_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.AAAA": {
      "deprecatedKeyName": "dns_AAAA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.AAAA.ttl": {
      "deprecatedKeyName": "dns_AAAA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.AFSDB": {
      "deprecatedKeyName": "dns_AFSDB",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.AFSDB.ttl": {
      "deprecatedKeyName": "dns_AFSDB_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.APL": {
      "deprecatedKeyName": "dns_APL",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.APL.ttl": {
      "deprecatedKeyName": "dns_APL_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CAA": {
      "deprecatedKeyName": "dns_CAA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CAA.ttl": {
      "deprecatedKeyName": "dns_CAA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CDNSKEY": {
      "deprecatedKeyName": "dns_CDNSKEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CDNSKEY.ttl": {
      "deprecatedKeyName": "dns_CDNSKEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CDS": {
      "deprecatedKeyName": "dns_CDS",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CDS.ttl": {
      "deprecatedKeyName": "dns_CDS_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CERT": {
      "deprecatedKeyName": "dns_CERT",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CERT.ttl": {
      "deprecatedKeyName": "dns_CERT_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CNAME": {
      "deprecatedKeyName": "dns_CNAME",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CNAME.ttl": {
      "deprecatedKeyName": "dns_CNAME_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CSYNC": {
      "deprecatedKeyName": "dns_CSYNC",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.CSYNC.ttl": {
      "deprecatedKeyName": "dns_CSYNC_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DHCID": {
      "deprecatedKeyName": "dns_DHCID",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DHCID.ttl": {
      "deprecatedKeyName": "dns_DHCID_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DLV": {
      "deprecatedKeyName": "dns_DLV",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DLV.ttl": {
      "deprecatedKeyName": "dns_DLV_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DNAME": {
      "deprecatedKeyName": "dns_DNAME",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DNAME.ttl": {
      "deprecatedKeyName": "dns_DNAME_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DNSKEY": {
      "deprecatedKeyName": "dns_DNSKEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DNSKEY.ttl": {
      "deprecatedKeyName": "dns_DNSKEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DS": {
      "deprecatedKeyName": "dns_DS",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.DS.ttl": {
      "deprecatedKeyName": "dns_DS_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.EUI48": {
      "deprecatedKeyName": "dns_EUI48",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.EUI48.ttl": {
      "deprecatedKeyName": "dns_EUI48_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.EUI64": {
      "deprecatedKeyName": "dns_EUI64",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.EUI64.ttl": {
      "deprecatedKeyName": "dns_EUI64_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HINFO": {
      "deprecatedKeyName": "dns_HINFO",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HINFO.ttl": {
      "deprecatedKeyName": "dns_HINFO_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HIP": {
      "deprecatedKeyName": "dns_HIP",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HIP.ttl": {
      "deprecatedKeyName": "dns_HIP_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HTTPS": {
      "deprecatedKeyName": "dns_HTTPS",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.HTTPS.ttl": {
      "deprecatedKeyName": "dns_HTTPS_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.IPSECKEY": {
      "deprecatedKeyName": "dns_IPSECKEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.IPSECKEY.ttl": {
      "deprecatedKeyName": "dns_IPSECKEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.KEY": {
      "deprecatedKeyName": "dns_KEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.KEY.ttl": {
      "deprecatedKeyName": "dns_KEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.KX": {
      "deprecatedKeyName": "dns_KX",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.KX.ttl": {
      "deprecatedKeyName": "dns_KX_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.LOC": {
      "deprecatedKeyName": "dns_LOC",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.LOC.ttl": {
      "deprecatedKeyName": "dns_LOC_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.MX": {
      "deprecatedKeyName": "dns_MX",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.MX.ttl": {
      "deprecatedKeyName": "dns_MX_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NAPTR": {
      "deprecatedKeyName": "dns_NAPTR",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NAPTR.ttl": {
      "deprecatedKeyName": "dns_NAPTR_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NS": {
      "deprecatedKeyName": "dns_NS",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NS.ttl": {
      "deprecatedKeyName": "dns_NS_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC": {
      "deprecatedKeyName": "dns_NSEC",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC.ttl": {
      "deprecatedKeyName": "dns_NSEC_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC3": {
      "deprecatedKeyName": "dns_NSEC3",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC3.ttl": {
      "deprecatedKeyName": "dns_NSEC3_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC3PARAM": {
      "deprecatedKeyName": "dns_NSEC3PARAM",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.NSEC3PARAM.ttl": {
      "deprecatedKeyName": "dns_NSEC3PARAM_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.OPENPGPKEY": {
      "deprecatedKeyName": "dns_OPENPGPKEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.OPENPGPKEY.ttl": {
      "deprecatedKeyName": "dns_OPENPGPKEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.PTR": {
      "deprecatedKeyName": "dns_PTR",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.PTR.ttl": {
      "deprecatedKeyName": "dns_PTR_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.RP": {
      "deprecatedKeyName": "dns_RP",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.RP.ttl": {
      "deprecatedKeyName": "dns_RP_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.RRSIG": {
      "deprecatedKeyName": "dns_RRSIG",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.RRSIG.ttl": {
      "deprecatedKeyName": "dns_RRSIG_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SIG": {
      "deprecatedKeyName": "dns_SIG",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SIG.ttl": {
      "deprecatedKeyName": "dns_SIG_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SMIMEA": {
      "deprecatedKeyName": "dns_SMIMEA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SMIMEA.ttl": {
      "deprecatedKeyName": "dns_SMIMEA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SOA": {
      "deprecatedKeyName": "dns_SOA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SOA.ttl": {
      "deprecatedKeyName": "dns_SOA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SRV": {
      "deprecatedKeyName": "dns_SRV",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SRV.ttl": {
      "deprecatedKeyName": "dns_SRV_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SSHFP": {
      "deprecatedKeyName": "dns_SSHFP",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SSHFP.ttl": {
      "deprecatedKeyName": "dns_SSHFP_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SVCB": {
      "deprecatedKeyName": "dns_SVCB",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.SVCB.ttl": {
      "deprecatedKeyName": "dns_SVCB_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TA": {
      "deprecatedKeyName": "dns_TA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TA.ttl": {
      "deprecatedKeyName": "dns_TA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TKEY": {
      "deprecatedKeyName": "dns_TKEY",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TKEY.ttl": {
      "deprecatedKeyName": "dns_TKEY_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TLSA": {
      "deprecatedKeyName": "dns_TLSA",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TLSA.ttl": {
      "deprecatedKeyName": "dns_TLSA_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TSIG": {
      "deprecatedKeyName": "dns_TSIG",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TSIG.ttl": {
      "deprecatedKeyName": "dns_TSIG_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TXT": {
      "deprecatedKeyName": "dns_TXT",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.TXT.ttl": {
      "deprecatedKeyName": "dns_TXT_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.URI": {
      "deprecatedKeyName": "dns_URI",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.URI.ttl": {
      "deprecatedKeyName": "dns_URI_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.ZONEMD": {
      "deprecatedKeyName": "dns_ZONEMD",
      "validationRegex": null,
      "deprecated": false
    },
    "dns.ZONEMD.ttl": {
      "deprecatedKeyName": "dns_ZONEMD_ttl",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.DOT.address": {
      "deprecatedKeyName": "DOT",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UNI.version.ERC20.address": {
      "deprecatedKeyName": "UNI_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.UNI.version.BEP20.address": {
      "deprecatedKeyName": "UNI_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.UNI.version.MATIC.address": {
      "deprecatedKeyName": "UNI_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UNI.version.HRC20.address": {
      "deprecatedKeyName": "UNI_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.UNI.version.XDAI.address": {
      "deprecatedKeyName": "UNI_XDAI",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SOL.address": {
      "deprecatedKeyName": "SOL",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.BUSD.version.ERC20.address": {
      "deprecatedKeyName": "BUSD_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BUSD.version.BEP20.address": {
      "deprecatedKeyName": "BUSD_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BUSD.version.HRC20.address": {
      "deprecatedKeyName": "BUSD_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ICP.address": {
      "deprecatedKeyName": "ICP",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.THETA.address": {
      "deprecatedKeyName": "THETA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.WBTC.version.ERC20.address": {
      "deprecatedKeyName": "WBTC_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.WBTC.version.MATIC.address": {
      "deprecatedKeyName": "WBTC_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.WBTC.version.FANTOM.address": {
      "deprecatedKeyName": "WBTC_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.WBTC.version.HRC20.address": {
      "deprecatedKeyName": "WBTC_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.WBTC.version.XDAI.address": {
      "deprecatedKeyName": "WBTC_XDAI",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.FIL.address": {
      "deprecatedKeyName": "FIL",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CDAI.address": {
      "deprecatedKeyName": "CDAI",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KSM.address": {
      "deprecatedKeyName": "KSM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.OKB.address": {
      "deprecatedKeyName": "OKB",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AAVE.version.ERC20.address": {
      "deprecatedKeyName": "AAVE_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AAVE.version.MATIC.address": {
      "deprecatedKeyName": "AAVE_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.AAVE.version.FANTOM.address": {
      "deprecatedKeyName": "AAVE_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.AAVE.version.HRC20.address": {
      "deprecatedKeyName": "AAVE_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SHIB.version.ERC20.address": {
      "deprecatedKeyName": "SHIB_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SHIB.version.MATIC.address": {
      "deprecatedKeyName": "SHIB_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SHIB.version.FANTOM.address": {
      "deprecatedKeyName": "SHIB_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CEL.version.ERC20.address": {
      "deprecatedKeyName": "CEL_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CEL.version.MATIC.address": {
      "deprecatedKeyName": "CEL_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CEL.version.FANTOM.address": {
      "deprecatedKeyName": "CEL_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CEL.version.HRC20.address": {
      "deprecatedKeyName": "CEL_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CUSDC.address": {
      "deprecatedKeyName": "CUSDC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CETH.address": {
      "deprecatedKeyName": "CETH",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AMP.address": {
      "deprecatedKeyName": "AMP",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CAKE.version.BEP20.address": {
      "deprecatedKeyName": "CAKE_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CAKE.version.HRC20.address": {
      "deprecatedKeyName": "CAKE_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MIOTA.address": {
      "deprecatedKeyName": "MIOTA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.FTT.address": {
      "deprecatedKeyName": "FTT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MKR.address": {
      "deprecatedKeyName": "MKR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TFUEL.address": {
      "deprecatedKeyName": "TFUEL",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.KLAY.address": {
      "deprecatedKeyName": "KLAY",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.LUNA.address": {
      "deprecatedKeyName": "LUNA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.AVAX.address": {
      "deprecatedKeyName": "AVAX",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.LEO.address": {
      "deprecatedKeyName": "LEO",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SAFEMOON.version.BEP20.address": {
      "deprecatedKeyName": "SAFEMOON_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SAFEMOON.version.HRC20.address": {
      "deprecatedKeyName": "SAFEMOON_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.UST.address": {
      "deprecatedKeyName": "UST",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RUNE.address": {
      "deprecatedKeyName": "RUNE",
      "validationRegex": "^(bnb|tbnb)[a-zA-HJ-NP-Z0-9]{39}$",
      "deprecated": false
    },
    "crypto.HBAR.address": {
      "deprecatedKeyName": "HBAR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.TEL.version.ERC20.address": {
      "deprecatedKeyName": "TEL_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TEL.version.MATIC.address": {
      "deprecatedKeyName": "TEL_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CHZ.address": {
      "deprecatedKeyName": "CHZ",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SUSHI.version.ERC20.address": {
      "deprecatedKeyName": "SUSHI_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SUSHI.version.BEP20.address": {
      "deprecatedKeyName": "SUSHI_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SUSHI.version.MATIC.address": {
      "deprecatedKeyName": "SUSHI_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SUSHI.version.FANTOM.address": {
      "deprecatedKeyName": "SUSHI_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SUSHI.version.HRC20.address": {
      "deprecatedKeyName": "SUSHI_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.EGLD.address": {
      "deprecatedKeyName": "EGLD",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.TUSD.version.ERC20.address": {
      "deprecatedKeyName": "TUSD_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TUSD.version.BEP20.address": {
      "deprecatedKeyName": "TUSD_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TUSD.version.AVAX.address": {
      "deprecatedKeyName": "TUSD_AVAX",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.TUSD.version.HRC20.address": {
      "deprecatedKeyName": "TUSD_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TUSD.version.BEP2.address": {
      "deprecatedKeyName": "TUSD_BEP2",
      "validationRegex": "^(bnb|tbnb)[a-zA-HJ-NP-Z0-9]{39}$",
      "deprecated": false
    },
    "crypto.TUSD.version.TRON.address": {
      "deprecatedKeyName": "TUSD_TRON",
      "validationRegex": "^[T][a-zA-HJ-NP-Z0-9]{33}$",
      "deprecated": false
    },
    "crypto.HBTC.version.ERC20.address": {
      "deprecatedKeyName": "HBTC_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HBTC.version.HRC20.address": {
      "deprecatedKeyName": "HBTC_HRC20",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SNX.version.ERC20.address": {
      "deprecatedKeyName": "SNX_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SNX.version.MATIC.address": {
      "deprecatedKeyName": "SNX_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SNX.version.FANTOM.address": {
      "deprecatedKeyName": "SNX_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SNX.version.HRC20.address": {
      "deprecatedKeyName": "SNX_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HOT.version.ERC20.address": {
      "deprecatedKeyName": "HOT_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HOT.version.HRC20.address": {
      "deprecatedKeyName": "HOT_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NEAR.address": {
      "deprecatedKeyName": "NEAR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.HNT.address": {
      "deprecatedKeyName": "HNT",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.STETH.address": {
      "deprecatedKeyName": "STETH",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NEXO.version.ERC20.address": {
      "deprecatedKeyName": "NEXO_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NEXO.version.FANTOM.address": {
      "deprecatedKeyName": "NEXO_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.PAX.address": {
      "deprecatedKeyName": "PAX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.STX.address": {
      "deprecatedKeyName": "STX",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MANA.version.ERC20.address": {
      "deprecatedKeyName": "MANA_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MANA.version.MATIC.address": {
      "deprecatedKeyName": "MANA_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MDX.version.HRC20.address": {
      "deprecatedKeyName": "MDX_HRC20",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MDX.version.BEP20.address": {
      "deprecatedKeyName": "MDX_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ZEN.address": {
      "deprecatedKeyName": "ZEN",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ARRR.address": {
      "deprecatedKeyName": "ARRR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.BNT.address": {
      "deprecatedKeyName": "BNT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LUSD.version.ERC20.address": {
      "deprecatedKeyName": "LUSD_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LUSD.version.MATIC.address": {
      "deprecatedKeyName": "LUSD_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.GRT.version.ERC20.address": {
      "deprecatedKeyName": "GRT_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GRT.version.MATIC.address": {
      "deprecatedKeyName": "GRT_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.GRT.version.HRC20.address": {
      "deprecatedKeyName": "GRT_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SC.address": {
      "deprecatedKeyName": "SC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.HUSD.version.ERC20.address": {
      "deprecatedKeyName": "HUSD_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HUSD.version.HRC20.address": {
      "deprecatedKeyName": "HUSD_HRC20",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CRV.version.ERC20.address": {
      "deprecatedKeyName": "CRV_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CRV.version.MATIC.address": {
      "deprecatedKeyName": "CRV_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.CRV.version.FANTOM.address": {
      "deprecatedKeyName": "CRV_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UMA.address": {
      "deprecatedKeyName": "UMA",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.WRX.version.BEP2.address": {
      "deprecatedKeyName": "WRX_BEP2",
      "validationRegex": "^(bnb|tbnb)[a-zA-HJ-NP-Z0-9]{39}$",
      "deprecated": false
    },
    "crypto.WRX.version.MATIC.address": {
      "deprecatedKeyName": "WRX_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.AR.address": {
      "deprecatedKeyName": "AR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.OMG.address": {
      "deprecatedKeyName": "OMG",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GT.address": {
      "deprecatedKeyName": "GT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.QNT.address": {
      "deprecatedKeyName": "QNT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CHSB.address": {
      "deprecatedKeyName": "CHSB",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.IOST.address": {
      "deprecatedKeyName": "IOST",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.NXM.address": {
      "deprecatedKeyName": "NXM",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KCS.address": {
      "deprecatedKeyName": "KCS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LPT.version.ERC20.address": {
      "deprecatedKeyName": "LPT_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LPT.version.HRC20.address": {
      "deprecatedKeyName": "LPT_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XSUSHI.address": {
      "deprecatedKeyName": "XSUSHI",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CUSDT.address": {
      "deprecatedKeyName": "CUSDT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FLOW.address": {
      "deprecatedKeyName": "FLOW",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ANKR.address": {
      "deprecatedKeyName": "ANKR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HBC.address": {
      "deprecatedKeyName": "HBC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.VGX.address": {
      "deprecatedKeyName": "VGX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FEI.address": {
      "deprecatedKeyName": "FEI",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BAKE.version.BEP20.address": {
      "deprecatedKeyName": "BAKE_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BAKE.version.HRC20.address": {
      "deprecatedKeyName": "BAKE_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.1INCH.version.ERC20.address": {
      "deprecatedKeyName": "1INCH_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.1INCH.version.BEP20.address": {
      "deprecatedKeyName": "1INCH_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.1INCH.version.MATIC.address": {
      "deprecatedKeyName": "1INCH_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.1INCH.version.HRC20.address": {
      "deprecatedKeyName": "1INCH_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CKB.address": {
      "deprecatedKeyName": "CKB",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.WOO.version.ERC20.address": {
      "deprecatedKeyName": "WOO_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.WOO.version.HRC20.address": {
      "deprecatedKeyName": "WOO_HRC20",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.TITAN.address": {
      "deprecatedKeyName": "TITAN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.OMI.address": {
      "deprecatedKeyName": "OMI",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MINA.address": {
      "deprecatedKeyName": "MINA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SETH.address": {
      "deprecatedKeyName": "SETH",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RSR.address": {
      "deprecatedKeyName": "RSR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.OXY.version.SOLANA.address": {
      "deprecatedKeyName": "OXY_SOLANA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.OXY.version.ERC20.address": {
      "deprecatedKeyName": "OXY_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.REN.version.ERC20.address": {
      "deprecatedKeyName": "REN_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.REN.version.HRC20.address": {
      "deprecatedKeyName": "REN_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XCH.address": {
      "deprecatedKeyName": "XCH",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.RENBTC.version.ERC20.address": {
      "deprecatedKeyName": "RENBTC_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RENBTC.version.BEP20.address": {
      "deprecatedKeyName": "RENBTC_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RENBTC.version.HRC20.address": {
      "deprecatedKeyName": "RENBTC_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.USDN.address": {
      "deprecatedKeyName": "USDN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BCHA.address": {
      "deprecatedKeyName": "BCHA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.LRC.address": {
      "deprecatedKeyName": "LRC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.PUNDIX.address": {
      "deprecatedKeyName": "PUNDIX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ERG.address": {
      "deprecatedKeyName": "ERG",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.WIN.address": {
      "deprecatedKeyName": "WIN",
      "validationRegex": "^[T][a-zA-HJ-NP-Z0-9]{33}$",
      "deprecated": false
    },
    "crypto.NPXS.address": {
      "deprecatedKeyName": "NPXS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TRIBE.address": {
      "deprecatedKeyName": "TRIBE",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MAID.address": {
      "deprecatedKeyName": "MAID",
      "validationRegex": "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
      "deprecated": false
    },
    "crypto.ASD.address": {
      "deprecatedKeyName": "ASD",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CUNI.address": {
      "deprecatedKeyName": "CUNI",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CELO.address": {
      "deprecatedKeyName": "CELO",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.DENT.address": {
      "deprecatedKeyName": "DENT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SNT.address": {
      "deprecatedKeyName": "SNT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FEG.version.ERC20.address": {
      "deprecatedKeyName": "FEG_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FEG.version.HRC20.address": {
      "deprecatedKeyName": "FEG_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SKL.address": {
      "deprecatedKeyName": "SKL",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ALUSD.address": {
      "deprecatedKeyName": "ALUSD",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MIR.version.ERC20.address": {
      "deprecatedKeyName": "MIR_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MIR.version.BEP20.address": {
      "deprecatedKeyName": "MIR_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GLM.address": {
      "deprecatedKeyName": "GLM",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.PAXG.version.ERC20.address": {
      "deprecatedKeyName": "PAXG_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.PAXG.version.HRC20.address": {
      "deprecatedKeyName": "PAXG_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CFX.address": {
      "deprecatedKeyName": "CFX",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UOS.address": {
      "deprecatedKeyName": "UOS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SVCS.address": {
      "deprecatedKeyName": "SVCS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.REEF.version.ERC20.address": {
      "deprecatedKeyName": "REEF_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.REEF.version.BEP20.address": {
      "deprecatedKeyName": "REEF_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.REEF.version.HRC20.address": {
      "deprecatedKeyName": "REEF_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GNO.address": {
      "deprecatedKeyName": "GNO",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.USDP.address": {
      "deprecatedKeyName": "USDP",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KAVA.address": {
      "deprecatedKeyName": "KAVA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ALCX.address": {
      "deprecatedKeyName": "ALCX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.EWT.address": {
      "deprecatedKeyName": "EWT",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.TON.address": {
      "deprecatedKeyName": "TON",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RLC.address": {
      "deprecatedKeyName": "RLC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AXS.address": {
      "deprecatedKeyName": "AXS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AUDIO.address": {
      "deprecatedKeyName": "AUDIO",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XVS.address": {
      "deprecatedKeyName": "XVS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BAND.version.ERC20.address": {
      "deprecatedKeyName": "BAND_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.BAND.version.FANTOM.address": {
      "deprecatedKeyName": "BAND_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.NMR.address": {
      "deprecatedKeyName": "NMR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.INJ.version.ERC20.address": {
      "deprecatedKeyName": "INJ_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.INJ.version.BEP20.address": {
      "deprecatedKeyName": "INJ_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.WAXP.address": {
      "deprecatedKeyName": "WAXP",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UQC.address": {
      "deprecatedKeyName": "UQC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.IOTX.address": {
      "deprecatedKeyName": "IOTX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FUN.address": {
      "deprecatedKeyName": "FUN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.OCEAN.address": {
      "deprecatedKeyName": "OCEAN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SAND.version.ERC20.address": {
      "deprecatedKeyName": "SAND_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SAND.version.HRC20.address": {
      "deprecatedKeyName": "SAND_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CTSI.version.ERC20.address": {
      "deprecatedKeyName": "CTSI_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CTSI.version.BEP20.address": {
      "deprecatedKeyName": "CTSI_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CTSI.version.MATIC.address": {
      "deprecatedKeyName": "CTSI_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.RAY.address": {
      "deprecatedKeyName": "RAY",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ANC.version.TERRA.address": {
      "deprecatedKeyName": "ANC_TERRA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ANC.version.ERC20.address": {
      "deprecatedKeyName": "ANC_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.IQ.version.ERC20.address": {
      "deprecatedKeyName": "IQ_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.IQ.version.BEP20.address": {
      "deprecatedKeyName": "IQ_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.IQ.version.MATIC.address": {
      "deprecatedKeyName": "IQ_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SUSD.version.ERC20.address": {
      "deprecatedKeyName": "SUSD_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SUSD.version.FANTOM.address": {
      "deprecatedKeyName": "SUSD_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.KLV.address": {
      "deprecatedKeyName": "KLV",
      "validationRegex": "^[T][a-zA-HJ-NP-Z0-9]{33}$",
      "deprecated": false
    },
    "crypto.BTCST.address": {
      "deprecatedKeyName": "BTCST",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TLM.address": {
      "deprecatedKeyName": "TLM",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AKT.address": {
      "deprecatedKeyName": "AKT",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.STMX.address": {
      "deprecatedKeyName": "STMX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.PROM.address": {
      "deprecatedKeyName": "PROM",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XPRT.address": {
      "deprecatedKeyName": "XPRT",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.SRM.version.ERC20.address": {
      "deprecatedKeyName": "SRM_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SRM.version.SOLANA.address": {
      "deprecatedKeyName": "SRM_SOLANA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.RPL.address": {
      "deprecatedKeyName": "RPL",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AGIX.address": {
      "deprecatedKeyName": "AGIX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CELR.address": {
      "deprecatedKeyName": "CELR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.FET.address": {
      "deprecatedKeyName": "FET",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.OXT.address": {
      "deprecatedKeyName": "OXT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ARDR.address": {
      "deprecatedKeyName": "ARDR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MATH.address": {
      "deprecatedKeyName": "MATH",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.10SET.address": {
      "deprecatedKeyName": "10SET",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.POLY.address": {
      "deprecatedKeyName": "POLY",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GUSD.address": {
      "deprecatedKeyName": "GUSD",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NKN.address": {
      "deprecatedKeyName": "NKN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CVC.address": {
      "deprecatedKeyName": "CVC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.GTC.address": {
      "deprecatedKeyName": "GTC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.STEEM.address": {
      "deprecatedKeyName": "STEEM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ORN.address": {
      "deprecatedKeyName": "ORN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KEEP.version.ERC20.address": {
      "deprecatedKeyName": "KEEP_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KEEP.version.HRC20.address": {
      "deprecatedKeyName": "KEEP_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HXRO.address": {
      "deprecatedKeyName": "HXRO",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ORBS.address": {
      "deprecatedKeyName": "ORBS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ALPHA.version.ERC20.address": {
      "deprecatedKeyName": "ALPHA_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ALPHA.version.BEP20.address": {
      "deprecatedKeyName": "ALPHA_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.DODO.version.ERC20.address": {
      "deprecatedKeyName": "DODO_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.DODO.version.BEP20.address": {
      "deprecatedKeyName": "DODO_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.OGN.address": {
      "deprecatedKeyName": "OGN",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KNCL.version.ERC20.address": {
      "deprecatedKeyName": "KNCL_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KNCL.version.FANTOM.address": {
      "deprecatedKeyName": "KNCL_FANTOM",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.KNCL.version.HRC20.address": {
      "deprecatedKeyName": "KNCL_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MED.address": {
      "deprecatedKeyName": "MED",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.XAUT.address": {
      "deprecatedKeyName": "XAUT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.VLX.address": {
      "deprecatedKeyName": "VLX",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.PHA.address": {
      "deprecatedKeyName": "PHA",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.KOBE.address": {
      "deprecatedKeyName": "KOBE",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.PERP.address": {
      "deprecatedKeyName": "PERP",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.XHV.address": {
      "deprecatedKeyName": "XHV",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.META.address": {
      "deprecatedKeyName": "META",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SEUR.address": {
      "deprecatedKeyName": "SEUR",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.MONA.address": {
      "deprecatedKeyName": "MONA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ANT.address": {
      "deprecatedKeyName": "ANT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.HYDRA.address": {
      "deprecatedKeyName": "HYDRA",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.ZKS.address": {
      "deprecatedKeyName": "ZKS",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SXP.version.ERC20.address": {
      "deprecatedKeyName": "SXP_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SXP.version.BEP20.address": {
      "deprecatedKeyName": "SXP_BEP20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.SXP.version.HRC20.address": {
      "deprecatedKeyName": "SXP_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.CSPR.address": {
      "deprecatedKeyName": "CSPR",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.MTL.address": {
      "deprecatedKeyName": "MTL",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NU.address": {
      "deprecatedKeyName": "NU",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ZMT.address": {
      "deprecatedKeyName": "ZMT",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.LOC.address": {
      "deprecatedKeyName": "LOC",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.TKO.address": {
      "deprecatedKeyName": "TKO",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.ETN.address": {
      "deprecatedKeyName": "ETN",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.UBT.version.ERC20.address": {
      "deprecatedKeyName": "UBT_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.UBT.version.MATIC.address": {
      "deprecatedKeyName": "UBT_MATIC",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.EXRD.address": {
      "deprecatedKeyName": "EXRD",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.NMX.address": {
      "deprecatedKeyName": "NMX",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.RIF.address": {
      "deprecatedKeyName": "RIF",
      "validationRegex": null,
      "deprecated": false
    },
    "crypto.STORJ.version.ERC20.address": {
      "deprecatedKeyName": "STORJ_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.STORJ.version.HRC20.address": {
      "deprecatedKeyName": "STORJ_HRC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.AVA.address": {
      "deprecatedKeyName": "AVA",
      "validationRegex": "^(bnb|tbnb)[a-zA-HJ-NP-Z0-9]{39}$",
      "deprecated": false
    },
    "crypto.DPI.version.ERC20.address": {
      "deprecatedKeyName": "DPI_ERC20",
      "validationRegex": "^0x[a-fA-F0-9]{40}$",
      "deprecated": false
    },
    "crypto.DPI.version.MATIC.address": {
      "deprecatedKeyName": "DPI_MATIC",
      "validationRegex": null,
      "deprecated": false
    }
  }
}
`)
