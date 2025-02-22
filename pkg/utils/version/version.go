/*
 * @Author: LinkLeong link@icewhale.com
 * @Date: 2022-05-13 18:15:46
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-06-29 14:29:34
 * @FilePath: /CasaOS/pkg/utils/version/version.go
 * @Description:
 * @Website: https://www.casaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package version

import (
	"strconv"
	"strings"

	"github.com/IceWhaleTech/CasaOS/model"
	"github.com/IceWhaleTech/CasaOS/types"
)

func IsNeedUpdate(version model.Version) (bool, model.Version) {

	v1 := strings.Split(version.Version, ".")

	v2 := strings.Split(types.CURRENTVERSION, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}
	for i := 0; i < len(v1); i++ {
		a, _ := strconv.Atoi(v1[i])
		b, _ := strconv.Atoi(v2[i])
		if i == 0 && a > b {
			continue
		}
		if a > b {
			return true, version
		}
		if a < b {
			return false, version
		}
	}
	return false, version
}
