package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"path"
	"regexp"
	"strconv"
)

func FormatPath(paths ...string) (r string) {

	for i := range paths {
		r = path.Join(r, paths[i])
	}

	return
}

func EmptyEthAddress(addr string) bool {

	return addr == new(common.Address).String()
}

func ValidEthAddress(addr string) bool {

	re := regexp.MustCompile(`(?i)(0x)?[0-9a-f]{40}`) // (?i) case-insensitive, (0x)? optional hex prefix
	return re.MatchString(addr)
}

func MarshalJSON(v interface{}) string {

	// Encode any struct to JSON
	bytes, err := json.Marshal(v)
	CheckError(err, WarningMode)

	return string(bytes)
}

func UnmarshalJSON(data string, v interface{}) {

	// String to bytes slice
	bytes := []byte(data)

	// Decode JSON to any struct
	err := json.Unmarshal(bytes, v)
	CheckError(err, WarningMode)
}

func ConvertFloatToBigInt(v float64) *big.Int {

	strv := strconv.FormatFloat(v*1e18, 'f', -1, 64)
	bi := new(big.Int)
	bi.SetString(strv, 10)

	return bi
}

func ConvertBigIntToFloat(v *big.Int) (r float64) {

	r, _ = new(big.Float).Quo(new(big.Float).SetInt(v), new(big.Float).SetUint64(1e18)).Float64()

	return
}
