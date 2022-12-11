package repo

import (
	"bytes"
	"fmt"
	"strconv"
)

type AssetID int

func (aid AssetID) MarshalJSON() ([]byte, error) {
	aidStr := fmt.Sprintf("%06d", aid)
	aidStr = fmt.Sprintf("%s-%s", aidStr[:3], aidStr[3:])
	return []byte(fmt.Sprintf(`"%s"`, aidStr)), nil

}

func (aid *AssetID) UnmarshalJSON(d []byte) error {
	d = bytes.Replace(d, []byte(`"`), []byte(``), -1)
	d = bytes.Replace(d, []byte(`-`), []byte(``), -1)

	aidInt, err := strconv.Atoi(string(d))
	if err != nil {
		return err
	}

	*aid = AssetID(aidInt)
	return nil

}
