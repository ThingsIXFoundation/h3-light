// Copyright 2022 Stichting ThingsIX Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package h3light

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tidwall/geojson/geometry"
)

type Cell uint64

func (c Cell) String() string {
	return fmt.Sprintf("%x", uint64(c))
}

func (c Cell) Uint64() uint64 {
	return uint64(c)
}

func (c Cell) Parent(res uint8) Cell {
	if c.Resolution() < res {
		return 0
	}
	if c.Resolution() == res {
		return c
	}

	u := uint64(c)

	resfill := (^uint64(0)) >> (19 + res*3)
	u |= resfill

	mask := uint64(0b1111 << 52)
	u &^= mask
	u |= (uint64(res) << 52)

	return Cell(u)
}

func (c Cell) Resolution() uint8 {
	return uint8((uint64(c) >> 52) & 0b1111)
}

func CellFromString(str string) (Cell, error) {
	i, err := strconv.ParseUint(str, 16, 64)
	if err != nil {
		return 0, err
	}

	return Cell(i), nil
}

func MustCellFromString(str string) Cell {
	c, err := CellFromString(str)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func LatLonToRes0ToCell(lat, lon float64) Cell {
	for res0, boundary := range res0map {
		poly := geometry.NewPoly(boundary, nil, nil)
		point := geometry.Point{X: lon, Y: lat}

		if poly.ContainsPoint(point) {
			return Cell(res0)
		}
	}

	return 0
}
