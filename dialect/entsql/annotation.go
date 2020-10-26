// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package entsql

import "github.com/facebook/ent/schema"

// Annotation is a builtin schema annotation for attaching
// SQL metadata to schema objects for both codegen and runtime.
type Annotation struct {
	// The Table option allows overriding the default table
	// name that is generated by ent. For example:
	//
	//	entsql.Annotation{
	//		Table: "Users",
	//	}
	//
	Table string
}

// Name describes the annotation name.
func (Annotation) Name() string {
	return "EntSQL"
}

var _ schema.Annotation = (*Annotation)(nil)