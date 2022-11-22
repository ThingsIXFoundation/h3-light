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

package main

import (
	"log"
	"os"
	"text/template"

	"github.com/tidwall/geojson/geometry"
	"github.com/uber/h3-go/v4"
)

func main() {

	res0 := h3.Res0Cells()
	res0Map := make(map[uint64][]geometry.Point)
	for _, res0 := range res0 {
		boundary := make([]geometry.Point, 0)
		for _, point := range res0.Boundary() {
			boundary = append(boundary, geometry.Point{Y: point.Lat, X: point.Lng})
		}

		res0Map[uint64(res0)] = boundary
	}

	f, err := os.Create("res0_generated.go")
	if err != nil {
		log.Fatalf("%v", err)
	}

	templ, err := template.New("res0_generated.go").Parse(
		`
	package h3light

	import (
		"github.com/tidwall/geojson/geometry"
	)

	var res0map = map[uint64][]geometry.Point {
	{{ range $hex, $boundary := . }}
	  {{ printf "0x%x" $hex }}: []geometry.Point{
		{{ range $boundary }}
		geometry.Point {
		  X: float64({{printf "%v" .X}}),
		  Y: float64({{printf "%v" .Y}}),
		},
		{{ end }}
	  },
	{{ end }}
	}
	`)

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = templ.Execute(f, res0Map)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
