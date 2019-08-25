// Copyright 2019 Intel Corporation. All Rights Reserved.
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
// Kernel drivers FPGA constants and structures
// This file is an input for "go tool cgo -godefs -- -Iuapi types.go | gofmt > ztypes.go"
//
// Include files should be copied from OPAE-SDK release, e.g. from
// https://github.com/OPAE/opae-sdk/tree/master/libopae/plugins/xfpga/
//
// +build ignore
// Code generated by carefully reading documentation and cherry-picking right kernel headers. DO NOT EDIT.
//

package linux

/*
#include "fpga-dfl.h"
#include "intel-fpga.h"
*/
import "C"

// Upstream FPGA DFL kernel driver
const (
	// DFL driver IOCTLs
	DFL_FPGA_GET_API_VERSION      = C.DFL_FPGA_GET_API_VERSION
	DFL_FPGA_CHECK_EXTENSION      = C.DFL_FPGA_CHECK_EXTENSION
	DFL_FPGA_PORT_RESET           = C.DFL_FPGA_PORT_RESET
	DFL_FPGA_PORT_GET_INFO        = C.DFL_FPGA_PORT_GET_INFO
	DFL_FPGA_PORT_GET_REGION_INFO = C.DFL_FPGA_PORT_GET_REGION_INFO
	DFL_FPGA_PORT_DMA_MAP         = C.DFL_FPGA_PORT_DMA_MAP
	DFL_FPGA_PORT_DMA_UNMAP       = C.DFL_FPGA_PORT_DMA_UNMAP
	DFL_FPGA_FME_PORT_PR          = C.DFL_FPGA_FME_PORT_PR
	DFL_FPGA_FME_PORT_RELEASE     = C.DFL_FPGA_FME_PORT_RELEASE
	DFL_FPGA_FME_PORT_ASSIGN      = C.DFL_FPGA_FME_PORT_ASSIGN

	// Flags in DflFpgaPortRegionInfo
	DFL_PORT_REGION_READ  = C.DFL_PORT_REGION_READ
	DFL_PORT_REGION_WRITE = C.DFL_PORT_REGION_WRITE
	DFL_PORT_REGION_MMAP  = C.DFL_PORT_REGION_MMAP

	// Index in DflFpgaPortRegionInfo
	DFL_PORT_REGION_INDEX_AFU = C.DFL_PORT_REGION_INDEX_AFU
	DFL_PORT_REGION_INDEX_STP = C.DFL_PORT_REGION_INDEX_STP
)

type DflFpgaPortInfo C.struct_dfl_fpga_port_info
type DflFpgaPortRegionInfo C.struct_dfl_fpga_port_region_info
type DflFpgaPortDMAMap C.struct_dfl_fpga_port_dma_map
type DflFpgaPortDMAUnmap C.struct_dfl_fpga_port_dma_unmap
type DflFpgaFmePortPR C.struct_dfl_fpga_fme_port_pr

// Out-of-tree Intel FPGA kernel driver
const (
	// Intel FPGA driver IOCTLs
	FPGA_GET_API_VERSION         = C.FPGA_GET_API_VERSION
	FPGA_CHECK_EXTENSION         = C.FPGA_CHECK_EXTENSION
	FPGA_PORT_RESET              = C.FPGA_PORT_RESET
	FPGA_PORT_GET_INFO           = C.FPGA_PORT_GET_INFO
	FPGA_PORT_GET_REGION_INFO    = C.FPGA_PORT_GET_REGION_INFO
	FPGA_PORT_DMA_MAP            = C.FPGA_PORT_DMA_MAP
	FPGA_PORT_DMA_UNMAP          = C.FPGA_PORT_DMA_UNMAP
	FPGA_PORT_UMSG_ENABLE        = C.FPGA_PORT_UMSG_ENABLE
	FPGA_PORT_UMSG_DISABLE       = C.FPGA_PORT_UMSG_DISABLE
	FPGA_PORT_UMSG_SET_MODE      = C.FPGA_PORT_UMSG_SET_MODE
	FPGA_PORT_UMSG_SET_BASE_ADDR = C.FPGA_PORT_UMSG_SET_BASE_ADDR
	FPGA_PORT_ERR_SET_IRQ        = C.FPGA_PORT_ERR_SET_IRQ
	FPGA_PORT_UAFU_SET_IRQ       = C.FPGA_PORT_UAFU_SET_IRQ
	FPGA_FME_PORT_PR             = C.FPGA_FME_PORT_PR
	FPGA_FME_PORT_RELEASE        = C.FPGA_FME_PORT_RELEASE
	FPGA_FME_PORT_ASSIGN         = C.FPGA_FME_PORT_ASSIGN
	FPGA_FME_GET_INFO            = C.FPGA_FME_GET_INFO
	FPGA_FME_ERR_SET_IRQ         = C.FPGA_FME_ERR_SET_IRQ

	// Capabilities in IntelFpgaPortInfo
	FPGA_PORT_CAP_ERR_IRQ  = C.FPGA_PORT_CAP_ERR_IRQ
	FPGA_PORT_CAP_UAFU_IRQ = C.FPGA_PORT_CAP_UAFU_IRQ

	// Flags in IntelFpgaPortRegionInfo
	FPGA_REGION_READ  = C.FPGA_REGION_READ
	FPGA_REGION_WRITE = C.FPGA_REGION_WRITE
	FPGA_REGION_MMAP  = C.FPGA_REGION_MMAP

	// Index in IntelFpgaPortRegionInfo
	FPGA_PORT_INDEX_UAFU = C.FPGA_PORT_INDEX_UAFU
	FPGA_PORT_INDEX_STP  = C.FPGA_PORT_INDEX_STP

	// Flags in IntelFpgaPortDmaMap
	FPGA_DMA_TO_DEV   = C.FPGA_DMA_TO_DEV
	FPGA_DMA_FROM_DEV = C.FPGA_DMA_FROM_DEV

	// Capabilities in IntelFpgaFmeInfo
	FPGA_FME_CAP_ERR_IRQ = C.FPGA_FME_CAP_ERR_IRQ
)

type IntelFpgaPortInfo C.struct_fpga_port_info
type IntelFpgaPortRegionInfo C.struct_fpga_port_region_info
type IntelFpgaPortDmaMap C.struct_fpga_port_dma_map
type IntelFpgaPortDmaUnmap C.struct_fpga_port_dma_unmap
type IntelFpgaPortUmsgCfg C.struct_fpga_port_umsg_cfg
type IntelFpgaPortUmsgBaseAddr C.struct_fpga_port_umsg_base_addr
type IntelFpgaPortErrIrqSet C.struct_fpga_port_err_irq_set
type IntelFpgaPortUafuIrqSet C.struct_fpga_port_uafu_irq_set
type IntelFpgaFmePortPR C.struct_fpga_fme_port_pr
type IntelFpgaFmePortRelease C.struct_fpga_fme_port_release
type IntelFpgaFmePortAssign C.struct_fpga_fme_port_assign
type IntelFpgaFmeInfo C.struct_fpga_fme_info
type IntelFpgaFmeErrIrqSet C.struct_fpga_fme_err_irq_set
