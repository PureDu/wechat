// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package message

var _ Eventer = &EventSubscribe{}
var _ Eventer = &EventScan{}
var _ Eventer = &EventLocation{}
var _ Eventer = &EventClickView{}