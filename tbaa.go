//===- tbaa.go - Bindings for TBAA --------------------------------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file defines bindings for the TBAA component.
//
//===----------------------------------------------------------------------===//

package llvm

/*
#include "llvm-c/Core.h"
#include "llvm-c/Comdat.h"
#include "IRBindings.h"
#include "TBAABindings.h"
#include <stdlib.h>
*/
import "C"

func IsTypeBasedAlias(lhs Value, rhs Value) bool {
	return C.LLVMTBAAIsAlias(lhs.C, rhs.C) != 0
}