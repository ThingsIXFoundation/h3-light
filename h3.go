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

	"github.com/Kl1mn/h3-go"
)

type Cell uint64

func (c Cell) String() string {
	return fmt.Sprintf("%x", uint64(c))
}

func (c Cell) Int64() int64 {
	return int64(c)
}

func (c Cell) Parent(res int) Cell {
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
	u |= uint64(res) << 52

	return Cell(u)
}

func (c Cell) Resolution() int {
	return int((uint64(c) >> 52) & 0b1111)
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
	return Cell(h3.FromGeo(h3.GeoCoord{Latitude: lat, Longitude: lon}, 0))
}

func LatLonToCell(lat, lon float64, res int) Cell {
	return Cell(h3.FromGeo(h3.GeoCoord{Latitude: lat, Longitude: lon}, res))
}

// MarshalText implements the encoding.TextMarshaler interface.
func (c Cell) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (c *Cell) UnmarshalText(text []byte) error {
	var err error
	*c, err = CellFromString(string(text))
	if err != nil {
		return err
	}

	return nil
}