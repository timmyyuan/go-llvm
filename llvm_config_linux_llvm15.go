//go:build !byollvm && linux && !llvm14 && !llvm16
// +build !byollvm,linux,!llvm14,!llvm16

package llvm

// #cgo CPPFLAGS: -I/usr/lib/llvm-15/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++14
import "C"

type run_build_sh int
