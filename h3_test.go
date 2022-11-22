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
	"testing"

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

func TestCell_Resolution(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want uint8
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
		res uint8
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
