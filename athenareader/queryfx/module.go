// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package queryfx

import (
	"database/sql"
	"github.com/uber/athenadriver/athenareader/configfx"
	drv "github.com/uber/athenadriver/go"
	"go.uber.org/fx"
)

// Module is to provide dependency of query to main app
var Module = fx.Provide(new)

// Params defines the dependencies or inputs
type Params struct {
	fx.In

	MC configfx.AthenaDriverConfig
}

// Result defines output
type Result struct {
	fx.Out

	QAD QueryAndDBConnection
}

// QueryAndDBConnection is the result of queryfx module
type QueryAndDBConnection struct {
	DB    *sql.DB
	Query string
}

func new(p Params) (Result, error) {
	// Open Connection.
	dsn := p.MC.DrvConfig.Stringify()
	db, _ := sql.Open(drv.DriverName, dsn)
	qad := QueryAndDBConnection{
		DB:    db,
		Query: p.MC.QueryString,
	}
	return Result{
		QAD: qad,
	}, nil
}
