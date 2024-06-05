package merkle_tree

/*
#cgo linux LDFLAGS: ${SRCDIR}/lib/libmerkle_tree.a -ldl -lrt -lm -lssl -lcrypto -Wl,--allow-multiple-definition
#cgo darwin LDFLAGS: ${SRCDIR}/lib/libmerkle_tree.dylib

#include "lib/merkle_tree.h"
*/
import "C"
import "unsafe"

const (
	MaxBatchSize = 8301147
)

func VerifyMerkleTreeBatch(batchBuffer [MaxBatchSize]byte, batchLen uint32, merkleRootBuffer [32]byte) bool {
	batchPtr := (*C.uchar)(unsafe.Pointer(&batchBuffer[0]))
	merkleRootPtr := (*C.uchar)(unsafe.Pointer(&merkleRootBuffer[0]))
	return (bool)(C.verify_batch_merkle_root_ffi(batchPtr, (C.uint32_t)(batchLen), merkleRootPtr))
}
