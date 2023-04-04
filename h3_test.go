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

package h3light_test

import (
	"math"
	"reflect"
	"testing"

	//"github.com/Kl1mn/h3-go"
	. "github.com/ThingsIXFoundation/h3-light"
	h3light "github.com/ThingsIXFoundation/h3-light"
	"github.com/uber/h3-go/v4"
)

func TestLatLonToRes0ToCell(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "801ffffffffffff",
			args: args{
				lat: float64(51.443655034915295),
				lon: float64(5.44695810089299),
			},
			want: h3.LatLngToCell(h3.LatLng{Lat: float64(51.443655034915295), Lng: float64(5.44695810089299)}, 0).String(),
		},
		{
			name: "801ffffffffffff",
			args: args{
				lat: float64(11.443655034915295),
				lon: float64(2.44695810089299),
			},
			want: h3.LatLngToCell(h3.LatLng{Lat: float64(11.443655034915295), Lng: float64(2.44695810089299)}, 0).String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LatLonToRes0ToCell(tt.args.lat, tt.args.lon).String(); got != tt.want {
				t.Errorf("LatLonToRes0ToCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatLonToRes0ToCellAll(t *testing.T) {
	for _, res0 := range h3.Res0Cells() {
		computedRes0 := LatLonToRes0ToCell(res0.LatLng().Lat, res0.LatLng().Lng)
		if res0.String() != computedRes0.String() {
			t.Errorf("LatLonToRes0ToCell() = %v, want %v", computedRes0, res0)
		}

	}
}

func TestCell_Resolution(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want int
	}{
		{
			name: "res5",
			c:    MustCellFromString("85283473fffffff"),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Resolution(); got != tt.want {
				t.Errorf("Cell.Resolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Parent(t *testing.T) {
	type args struct {
		res int
	}
	tests := []struct {
		name string
		c    Cell
		args args
		want Cell
	}{
		{
			name: "85283473fffffff to res4",
			c:    MustCellFromString("85283473fffffff"),
			args: args{
				res: 4,
			},
			want: MustCellFromString("8428347ffffffff"),
		},
		{
			name: "85283473fffffff to res0",
			c:    MustCellFromString("85283473fffffff"),
			args: args{
				res: 0,
			},
			want: MustCellFromString("8029fffffffffff"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Parent(tt.args.res); got != tt.want {
				t.Errorf("Cell.Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_LatLon(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		lat  float64
		lon  float64
	}{
		{
			name: "801ffffffffffff",
			c:    MustCellFromString("801ffffffffffff"),
			lat:  h3.Cell(h3.IndexFromString("801ffffffffffff")).LatLng().Lat,
			lon:  h3.Cell(h3.IndexFromString("801ffffffffffff")).LatLng().Lng,
		},
		{
			name: "85283473fffffff",
			c:    MustCellFromString("85283473fffffff"),
			lat:  h3.Cell(h3.IndexFromString("85283473fffffff")).LatLng().Lat,
			lon:  h3.Cell(h3.IndexFromString("85283473fffffff")).LatLng().Lng,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lat, lon := tt.c.LatLon()
			if math.Abs(lat-tt.lat) > 0.000000001 {
				t.Errorf("Cell.LatLon() got = %v, want %v (diff %v)", lat, tt.lat, math.Abs(lat-tt.lat))
			}
			if math.Abs(lon-tt.lon) > 0.000000001 {
				t.Errorf("Cell.LatLon() got = %v, want %v (diff %v)", lon, tt.lon, math.Abs(lon-tt.lon))
			}
		})
	}
}

func h3CellsToH3LightCells(cells []h3.Cell) []h3light.Cell {
	ret := make([]h3light.Cell, len(cells))
	for i, c := range cells {
		ret[i] = h3light.Cell(c)
	}

	return ret
}

func TestGetRes0Cells(t *testing.T) {
	tests := []struct {
		name string
		want []Cell
	}{
		{
			"test",
			h3CellsToH3LightCells(h3.Res0Cells()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRes0Cells(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRes0Cells() = %v, want %v", got, tt.want)
			}
		})
	}
}
