// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package zpages implements a collection of HTML pages that display RPC stats
// and trace data, and also functions to write that same data in plain text to
// an io.Writer.
//
// Users can also embed the HTML for stats and traces in custom status pages.
//
// zpages are currrently work-in-process and cannot display minutely and
// hourly stats correctly.
//
// Performance
//
// Installing the zpages has a performance overhead because additional traces
// and stats will be collected in-process. In most cases, we expect this
// overhead will not be significant but it depends on many factors, including
// how many spans your process creates and how richly annotated they are.
package zpages // import "github.com/davidwalter0/go-opencensus/zpages"

import (
	"net/http"
)

// Handler is an http.Handler that serves the zpages.
var Handler http.Handler

func init() {
	zpagesMux := http.NewServeMux()
	zpagesMux.HandleFunc("/rpcz", rpczHandler)
	zpagesMux.HandleFunc("/tracez", tracezHandler)
	zpagesMux.Handle("/public/", http.FileServer(fs))
	Handler = zpagesMux
}
