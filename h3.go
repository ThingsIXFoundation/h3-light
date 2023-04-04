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

func (c Cell) LatLon() (float64, float64) {
	geo := h3.ToGeo(h3.H3Index(c))
	return geo.Latitude, geo.Longitude
}

func (c Cell) DatabaseCell() DatabaseCell {
	return DatabaseCellFromCell(uint64(c))
}

func (c *Cell) DatabaseCellPtr() *DatabaseCell {
	if c == nil {
		return nil
	}

	ret := c.DatabaseCell()

	return &ret
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

func GetRes0Cells() []Cell {
	return []Cell{0x8001fffffffffff, 0x8003fffffffffff, 0x8005fffffffffff, 0x8007fffffffffff, 0x8009fffffffffff, 0x800bfffffffffff, 0x800dfffffffffff, 0x800ffffffffffff, 0x8011fffffffffff, 0x8013fffffffffff, 0x8015fffffffffff, 0x8017fffffffffff, 0x8019fffffffffff, 0x801bfffffffffff, 0x801dfffffffffff, 0x801ffffffffffff, 0x8021fffffffffff, 0x8023fffffffffff, 0x8025fffffffffff, 0x8027fffffffffff, 0x8029fffffffffff, 0x802bfffffffffff, 0x802dfffffffffff, 0x802ffffffffffff, 0x8031fffffffffff, 0x8033fffffffffff, 0x8035fffffffffff, 0x8037fffffffffff, 0x8039fffffffffff, 0x803bfffffffffff, 0x803dfffffffffff, 0x803ffffffffffff, 0x8041fffffffffff, 0x8043fffffffffff, 0x8045fffffffffff, 0x8047fffffffffff, 0x8049fffffffffff, 0x804bfffffffffff, 0x804dfffffffffff, 0x804ffffffffffff, 0x8051fffffffffff, 0x8053fffffffffff, 0x8055fffffffffff, 0x8057fffffffffff, 0x8059fffffffffff, 0x805bfffffffffff, 0x805dfffffffffff, 0x805ffffffffffff, 0x8061fffffffffff, 0x8063fffffffffff, 0x8065fffffffffff, 0x8067fffffffffff, 0x8069fffffffffff, 0x806bfffffffffff, 0x806dfffffffffff, 0x806ffffffffffff, 0x8071fffffffffff, 0x8073fffffffffff, 0x8075fffffffffff, 0x8077fffffffffff, 0x8079fffffffffff, 0x807bfffffffffff, 0x807dfffffffffff, 0x807ffffffffffff, 0x8081fffffffffff, 0x8083fffffffffff, 0x8085fffffffffff, 0x8087fffffffffff, 0x8089fffffffffff, 0x808bfffffffffff, 0x808dfffffffffff, 0x808ffffffffffff, 0x8091fffffffffff, 0x8093fffffffffff, 0x8095fffffffffff, 0x8097fffffffffff, 0x8099fffffffffff, 0x809bfffffffffff, 0x809dfffffffffff, 0x809ffffffffffff, 0x80a1fffffffffff, 0x80a3fffffffffff, 0x80a5fffffffffff, 0x80a7fffffffffff, 0x80a9fffffffffff, 0x80abfffffffffff, 0x80adfffffffffff, 0x80affffffffffff, 0x80b1fffffffffff, 0x80b3fffffffffff, 0x80b5fffffffffff, 0x80b7fffffffffff, 0x80b9fffffffffff, 0x80bbfffffffffff, 0x80bdfffffffffff, 0x80bffffffffffff, 0x80c1fffffffffff, 0x80c3fffffffffff, 0x80c5fffffffffff, 0x80c7fffffffffff, 0x80c9fffffffffff, 0x80cbfffffffffff, 0x80cdfffffffffff, 0x80cffffffffffff, 0x80d1fffffffffff, 0x80d3fffffffffff, 0x80d5fffffffffff, 0x80d7fffffffffff, 0x80d9fffffffffff, 0x80dbfffffffffff, 0x80ddfffffffffff, 0x80dffffffffffff, 0x80e1fffffffffff, 0x80e3fffffffffff, 0x80e5fffffffffff, 0x80e7fffffffffff, 0x80e9fffffffffff, 0x80ebfffffffffff, 0x80edfffffffffff, 0x80effffffffffff, 0x80f1fffffffffff, 0x80f3fffffffffff}
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
