//===- TBAABINDINGS.h - Additional bindings for IR ----------------*- C++ -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file defines additional C bindings for the TBAA component.
//
//===----------------------------------------------------------------------===//

#ifndef LLVM_BINDINGS_GO_LLVM_TBAABINDINGS_H
#define LLVM_BINDINGS_GO_LLVM_TBAABINDINGS_H

#include "llvm-c/Core.h"
#include "llvm-c/DebugInfo.h"
#ifdef __cplusplus
#include "llvm/IR/Metadata.h"
#include "llvm/Support/CBindingWrapping.h"
#endif

#include <stdint.h>

LLVM_C_EXTERN_C_BEGIN

LLVMBool LLVMTBAAIsAlias(LLVMValueRef LhsInst, LLVMValueRef RhsInst);

LLVM_C_EXTERN_C_END

#endif
