package visa

/*
#include <stdlib.h>
#include "visa.h"
*/
import "C"

const (
	SPEC_VERSION = C.VI_SPEC_VERSION

	// Attributes (platform independent size)
	ATTR_RSRC_CLASS                  = C.VI_ATTR_RSRC_CLASS
	ATTR_RSRC_NAME                   = C.VI_ATTR_RSRC_NAME
	ATTR_RSRC_IMPL_VERSION           = C.VI_ATTR_RSRC_IMPL_VERSION
	ATTR_RSRC_LOCK_STATE             = C.VI_ATTR_RSRC_LOCK_STATE
	ATTR_MAX_QUEUE_LENGTH            = C.VI_ATTR_MAX_QUEUE_LENGTH
	ATTR_USER_DATA_32                = C.VI_ATTR_USER_DATA_32
	ATTR_FDC_CHNL                    = C.VI_ATTR_FDC_CHNL
	ATTR_FDC_MODE                    = C.VI_ATTR_FDC_MODE
	ATTR_FDC_GEN_SIGNAL_EN           = C.VI_ATTR_FDC_GEN_SIGNAL_EN
	ATTR_FDC_USE_PAIR                = C.VI_ATTR_FDC_USE_PAIR
	ATTR_SEND_END_EN                 = C.VI_ATTR_SEND_END_EN
	ATTR_TERMCHAR                    = C.VI_ATTR_TERMCHAR
	ATTR_TMO_VALUE                   = C.VI_ATTR_TMO_VALUE
	ATTR_GPIB_READDR_EN              = C.VI_ATTR_GPIB_READDR_EN
	ATTR_IO_PROT                     = C.VI_ATTR_IO_PROT
	ATTR_DMA_ALLOW_EN                = C.VI_ATTR_DMA_ALLOW_EN
	ATTR_ASRL_BAUD                   = C.VI_ATTR_ASRL_BAUD
	ATTR_ASRL_DATA_BITS              = C.VI_ATTR_ASRL_DATA_BITS
	ATTR_ASRL_PARITY                 = C.VI_ATTR_ASRL_PARITY
	ATTR_ASRL_STOP_BITS              = C.VI_ATTR_ASRL_STOP_BITS
	ATTR_ASRL_FLOW_CNTRL             = C.VI_ATTR_ASRL_FLOW_CNTRL
	ATTR_RD_BUF_OPER_MODE            = C.VI_ATTR_RD_BUF_OPER_MODE
	ATTR_RD_BUF_SIZE                 = C.VI_ATTR_RD_BUF_SIZE
	ATTR_WR_BUF_OPER_MODE            = C.VI_ATTR_WR_BUF_OPER_MODE
	ATTR_WR_BUF_SIZE                 = C.VI_ATTR_WR_BUF_SIZE
	ATTR_SUPPRESS_END_EN             = C.VI_ATTR_SUPPRESS_END_EN
	ATTR_TERMCHAR_EN                 = C.VI_ATTR_TERMCHAR_EN
	ATTR_DEST_ACCESS_PRIV            = C.VI_ATTR_DEST_ACCESS_PRIV
	ATTR_DEST_BYTE_ORDER             = C.VI_ATTR_DEST_BYTE_ORDER
	ATTR_SRC_ACCESS_PRIV             = C.VI_ATTR_SRC_ACCESS_PRIV
	ATTR_SRC_BYTE_ORDER              = C.VI_ATTR_SRC_BYTE_ORDER
	ATTR_SRC_INCREMENT               = C.VI_ATTR_SRC_INCREMENT
	ATTR_DEST_INCREMENT              = C.VI_ATTR_DEST_INCREMENT
	ATTR_WIN_ACCESS_PRIV             = C.VI_ATTR_WIN_ACCESS_PRIV
	ATTR_WIN_BYTE_ORDER              = C.VI_ATTR_WIN_BYTE_ORDER
	ATTR_GPIB_ATN_STATE              = C.VI_ATTR_GPIB_ATN_STATE
	ATTR_GPIB_ADDR_STATE             = C.VI_ATTR_GPIB_ADDR_STATE
	ATTR_GPIB_CIC_STATE              = C.VI_ATTR_GPIB_CIC_STATE
	ATTR_GPIB_NDAC_STATE             = C.VI_ATTR_GPIB_NDAC_STATE
	ATTR_GPIB_SRQ_STATE              = C.VI_ATTR_GPIB_SRQ_STATE
	ATTR_GPIB_SYS_CNTRL_STATE        = C.VI_ATTR_GPIB_SYS_CNTRL_STATE
	ATTR_GPIB_HS488_CBL_LEN          = C.VI_ATTR_GPIB_HS488_CBL_LEN
	ATTR_CMDR_LA                     = C.VI_ATTR_CMDR_LA
	ATTR_VXI_DEV_CLASS               = C.VI_ATTR_VXI_DEV_CLASS
	ATTR_MAINFRAME_LA                = C.VI_ATTR_MAINFRAME_LA
	ATTR_MANF_NAME                   = C.VI_ATTR_MANF_NAME
	ATTR_MODEL_NAME                  = C.VI_ATTR_MODEL_NAME
	ATTR_VXI_VME_INTR_STATUS         = C.VI_ATTR_VXI_VME_INTR_STATUS
	ATTR_VXI_TRIG_STATUS             = C.VI_ATTR_VXI_TRIG_STATUS
	ATTR_VXI_VME_SYSFAIL_STATE       = C.VI_ATTR_VXI_VME_SYSFAIL_STATE
	ATTR_WIN_BASE_ADDR_32            = C.VI_ATTR_WIN_BASE_ADDR_32
	ATTR_WIN_SIZE_32                 = C.VI_ATTR_WIN_SIZE_32
	ATTR_ASRL_AVAIL_NUM              = C.VI_ATTR_ASRL_AVAIL_NUM
	ATTR_MEM_BASE_32                 = C.VI_ATTR_MEM_BASE_32
	ATTR_ASRL_CTS_STATE              = C.VI_ATTR_ASRL_CTS_STATE
	ATTR_ASRL_DCD_STATE              = C.VI_ATTR_ASRL_DCD_STATE
	ATTR_ASRL_DSR_STATE              = C.VI_ATTR_ASRL_DSR_STATE
	ATTR_ASRL_DTR_STATE              = C.VI_ATTR_ASRL_DTR_STATE
	ATTR_ASRL_END_IN                 = C.VI_ATTR_ASRL_END_IN
	ATTR_ASRL_END_OUT                = C.VI_ATTR_ASRL_END_OUT
	ATTR_ASRL_REPLACE_CHAR           = C.VI_ATTR_ASRL_REPLACE_CHAR
	ATTR_ASRL_RI_STATE               = C.VI_ATTR_ASRL_RI_STATE
	ATTR_ASRL_RTS_STATE              = C.VI_ATTR_ASRL_RTS_STATE
	ATTR_ASRL_XON_CHAR               = C.VI_ATTR_ASRL_XON_CHAR
	ATTR_ASRL_XOFF_CHAR              = C.VI_ATTR_ASRL_XOFF_CHAR
	ATTR_WIN_ACCESS                  = C.VI_ATTR_WIN_ACCESS
	ATTR_RM_SESSION                  = C.VI_ATTR_RM_SESSION
	ATTR_VXI_LA                      = C.VI_ATTR_VXI_LA
	ATTR_MANF_ID                     = C.VI_ATTR_MANF_ID
	ATTR_MEM_SIZE_32                 = C.VI_ATTR_MEM_SIZE_32
	ATTR_MEM_SPACE                   = C.VI_ATTR_MEM_SPACE
	ATTR_MODEL_CODE                  = C.VI_ATTR_MODEL_CODE
	ATTR_SLOT                        = C.VI_ATTR_SLOT
	ATTR_INTF_INST_NAME              = C.VI_ATTR_INTF_INST_NAME
	ATTR_IMMEDIATE_SERV              = C.VI_ATTR_IMMEDIATE_SERV
	ATTR_INTF_PARENT_NUM             = C.VI_ATTR_INTF_PARENT_NUM
	ATTR_RSRC_SPEC_VERSION           = C.VI_ATTR_RSRC_SPEC_VERSION
	ATTR_INTF_TYPE                   = C.VI_ATTR_INTF_TYPE
	ATTR_GPIB_PRIMARY_ADDR           = C.VI_ATTR_GPIB_PRIMARY_ADDR
	ATTR_GPIB_SECONDARY_ADDR         = C.VI_ATTR_GPIB_SECONDARY_ADDR
	ATTR_RSRC_MANF_NAME              = C.VI_ATTR_RSRC_MANF_NAME
	ATTR_RSRC_MANF_ID                = C.VI_ATTR_RSRC_MANF_ID
	ATTR_INTF_NUM                    = C.VI_ATTR_INTF_NUM
	ATTR_TRIG_ID                     = C.VI_ATTR_TRIG_ID
	ATTR_GPIB_REN_STATE              = C.VI_ATTR_GPIB_REN_STATE
	ATTR_GPIB_UNADDR_EN              = C.VI_ATTR_GPIB_UNADDR_EN
	ATTR_DEV_STATUS_BYTE             = C.VI_ATTR_DEV_STATUS_BYTE
	ATTR_FILE_APPEND_EN              = C.VI_ATTR_FILE_APPEND_EN
	ATTR_VXI_TRIG_SUPPORT            = C.VI_ATTR_VXI_TRIG_SUPPORT
	ATTR_TCPIP_ADDR                  = C.VI_ATTR_TCPIP_ADDR
	ATTR_TCPIP_HOSTNAME              = C.VI_ATTR_TCPIP_HOSTNAME
	ATTR_TCPIP_PORT                  = C.VI_ATTR_TCPIP_PORT
	ATTR_TCPIP_DEVICE_NAME           = C.VI_ATTR_TCPIP_DEVICE_NAME
	ATTR_TCPIP_NODELAY               = C.VI_ATTR_TCPIP_NODELAY
	ATTR_TCPIP_KEEPALIVE             = C.VI_ATTR_TCPIP_KEEPALIVE
	ATTR_4882_COMPLIANT              = C.VI_ATTR_4882_COMPLIANT
	ATTR_USB_SERIAL_NUM              = C.VI_ATTR_USB_SERIAL_NUM
	ATTR_USB_INTFC_NUM               = C.VI_ATTR_USB_INTFC_NUM
	ATTR_USB_PROTOCOL                = C.VI_ATTR_USB_PROTOCOL
	ATTR_USB_MAX_INTR_SIZE           = C.VI_ATTR_USB_MAX_INTR_SIZE
	ATTR_PXI_DEV_NUM                 = C.VI_ATTR_PXI_DEV_NUM
	ATTR_PXI_FUNC_NUM                = C.VI_ATTR_PXI_FUNC_NUM
	ATTR_PXI_BUS_NUM                 = C.VI_ATTR_PXI_BUS_NUM
	ATTR_PXI_CHASSIS                 = C.VI_ATTR_PXI_CHASSIS
	ATTR_PXI_SLOTPATH                = C.VI_ATTR_PXI_SLOTPATH
	ATTR_PXI_SLOT_LBUS_LEFT          = C.VI_ATTR_PXI_SLOT_LBUS_LEFT
	ATTR_PXI_SLOT_LBUS_RIGHT         = C.VI_ATTR_PXI_SLOT_LBUS_RIGHT
	ATTR_PXI_TRIG_BUS                = C.VI_ATTR_PXI_TRIG_BUS
	ATTR_PXI_STAR_TRIG_BUS           = C.VI_ATTR_PXI_STAR_TRIG_BUS
	ATTR_PXI_STAR_TRIG_LINE          = C.VI_ATTR_PXI_STAR_TRIG_LINE
	ATTR_PXI_SRC_TRIG_BUS            = C.VI_ATTR_PXI_SRC_TRIG_BUS
	ATTR_PXI_DEST_TRIG_BUS           = C.VI_ATTR_PXI_DEST_TRIG_BUS
	ATTR_PXI_MEM_TYPE_BAR0           = C.VI_ATTR_PXI_MEM_TYPE_BAR0
	ATTR_PXI_MEM_TYPE_BAR1           = C.VI_ATTR_PXI_MEM_TYPE_BAR1
	ATTR_PXI_MEM_TYPE_BAR2           = C.VI_ATTR_PXI_MEM_TYPE_BAR2
	ATTR_PXI_MEM_TYPE_BAR3           = C.VI_ATTR_PXI_MEM_TYPE_BAR3
	ATTR_PXI_MEM_TYPE_BAR4           = C.VI_ATTR_PXI_MEM_TYPE_BAR4
	ATTR_PXI_MEM_TYPE_BAR5           = C.VI_ATTR_PXI_MEM_TYPE_BAR5
	ATTR_PXI_MEM_BASE_BAR0_32        = C.VI_ATTR_PXI_MEM_BASE_BAR0_32
	ATTR_PXI_MEM_BASE_BAR1_32        = C.VI_ATTR_PXI_MEM_BASE_BAR1_32
	ATTR_PXI_MEM_BASE_BAR2_32        = C.VI_ATTR_PXI_MEM_BASE_BAR2_32
	ATTR_PXI_MEM_BASE_BAR3_32        = C.VI_ATTR_PXI_MEM_BASE_BAR3_32
	ATTR_PXI_MEM_BASE_BAR4_32        = C.VI_ATTR_PXI_MEM_BASE_BAR4_32
	ATTR_PXI_MEM_BASE_BAR5_32        = C.VI_ATTR_PXI_MEM_BASE_BAR5_32
	ATTR_PXI_MEM_BASE_BAR0_64        = C.VI_ATTR_PXI_MEM_BASE_BAR0_64
	ATTR_PXI_MEM_BASE_BAR1_64        = C.VI_ATTR_PXI_MEM_BASE_BAR1_64
	ATTR_PXI_MEM_BASE_BAR2_64        = C.VI_ATTR_PXI_MEM_BASE_BAR2_64
	ATTR_PXI_MEM_BASE_BAR3_64        = C.VI_ATTR_PXI_MEM_BASE_BAR3_64
	ATTR_PXI_MEM_BASE_BAR4_64        = C.VI_ATTR_PXI_MEM_BASE_BAR4_64
	ATTR_PXI_MEM_BASE_BAR5_64        = C.VI_ATTR_PXI_MEM_BASE_BAR5_64
	ATTR_PXI_MEM_SIZE_BAR0_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR0_32
	ATTR_PXI_MEM_SIZE_BAR1_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR1_32
	ATTR_PXI_MEM_SIZE_BAR2_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR2_32
	ATTR_PXI_MEM_SIZE_BAR3_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR3_32
	ATTR_PXI_MEM_SIZE_BAR4_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR4_32
	ATTR_PXI_MEM_SIZE_BAR5_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR5_32
	ATTR_PXI_MEM_SIZE_BAR0_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR0_64
	ATTR_PXI_MEM_SIZE_BAR1_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR1_64
	ATTR_PXI_MEM_SIZE_BAR2_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR2_64
	ATTR_PXI_MEM_SIZE_BAR3_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR3_64
	ATTR_PXI_MEM_SIZE_BAR4_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR4_64
	ATTR_PXI_MEM_SIZE_BAR5_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR5_64
	ATTR_PXI_IS_EXPRESS              = C.VI_ATTR_PXI_IS_EXPRESS
	ATTR_PXI_SLOT_LWIDTH             = C.VI_ATTR_PXI_SLOT_LWIDTH
	ATTR_PXI_MAX_LWIDTH              = C.VI_ATTR_PXI_MAX_LWIDTH
	ATTR_PXI_ACTUAL_LWIDTH           = C.VI_ATTR_PXI_ACTUAL_LWIDTH
	ATTR_PXI_DSTAR_BUS               = C.VI_ATTR_PXI_DSTAR_BUS
	ATTR_PXI_DSTAR_SET               = C.VI_ATTR_PXI_DSTAR_SET
	ATTR_PXI_ALLOW_WRITE_COMBINE     = C.VI_ATTR_PXI_ALLOW_WRITE_COMBINE
	ATTR_TCPIP_HISLIP_OVERLAP_EN     = C.VI_ATTR_TCPIP_HISLIP_OVERLAP_EN
	ATTR_TCPIP_HISLIP_VERSION        = C.VI_ATTR_TCPIP_HISLIP_VERSION
	ATTR_TCPIP_HISLIP_MAX_MESSAGE_KB = C.VI_ATTR_TCPIP_HISLIP_MAX_MESSAGE_KB
	ATTR_TCPIP_IS_HISLIP             = C.VI_ATTR_TCPIP_IS_HISLIP

	ATTR_JOB_ID              = C.VI_ATTR_JOB_ID
	ATTR_EVENT_TYPE          = C.VI_ATTR_EVENT_TYPE
	ATTR_SIGP_STATUS_ID      = C.VI_ATTR_SIGP_STATUS_ID
	ATTR_RECV_TRIG_ID        = C.VI_ATTR_RECV_TRIG_ID
	ATTR_INTR_STATUS_ID      = C.VI_ATTR_INTR_STATUS_ID
	ATTR_STATUS              = C.VI_ATTR_STATUS
	ATTR_RET_COUNT_32        = C.VI_ATTR_RET_COUNT_32
	ATTR_BUFFER              = C.VI_ATTR_BUFFER
	ATTR_RECV_INTR_LEVEL     = C.VI_ATTR_RECV_INTR_LEVEL
	ATTR_OPER_NAME           = C.VI_ATTR_OPER_NAME
	ATTR_GPIB_RECV_CIC_STATE = C.VI_ATTR_GPIB_RECV_CIC_STATE
	ATTR_RECV_TCPIP_ADDR     = C.VI_ATTR_RECV_TCPIP_ADDR
	ATTR_USB_RECV_INTR_SIZE  = C.VI_ATTR_USB_RECV_INTR_SIZE
	ATTR_USB_RECV_INTR_DATA  = C.VI_ATTR_USB_RECV_INTR_DATA
	ATTR_PXI_RECV_INTR_SEQ   = C.VI_ATTR_PXI_RECV_INTR_SEQ
	ATTR_PXI_RECV_INTR_DATA  = C.VI_ATTR_PXI_RECV_INTR_DATA

	// Attributes (platform dependent size)
	ATTR_USER_DATA = C.VI_ATTR_USER_DATA
	ATTR_RET_COUNT = C.VI_ATTR_RET_COUNT

	ATTR_WIN_BASE_ADDR     = C.VI_ATTR_WIN_BASE_ADDR
	ATTR_WIN_SIZE          = C.VI_ATTR_WIN_SIZE
	ATTR_MEM_BASE          = C.VI_ATTR_MEM_BASE
	ATTR_MEM_SIZE          = C.VI_ATTR_MEM_SIZE
	ATTR_PXI_MEM_BASE_BAR0 = C.VI_ATTR_PXI_MEM_BASE_BAR0
	ATTR_PXI_MEM_BASE_BAR1 = C.VI_ATTR_PXI_MEM_BASE_BAR1
	ATTR_PXI_MEM_BASE_BAR2 = C.VI_ATTR_PXI_MEM_BASE_BAR2
	ATTR_PXI_MEM_BASE_BAR3 = C.VI_ATTR_PXI_MEM_BASE_BAR3
	ATTR_PXI_MEM_BASE_BAR4 = C.VI_ATTR_PXI_MEM_BASE_BAR4
	ATTR_PXI_MEM_BASE_BAR5 = C.VI_ATTR_PXI_MEM_BASE_BAR5
	ATTR_PXI_MEM_SIZE_BAR0 = C.VI_ATTR_PXI_MEM_SIZE_BAR0
	ATTR_PXI_MEM_SIZE_BAR1 = C.VI_ATTR_PXI_MEM_SIZE_BAR1
	ATTR_PXI_MEM_SIZE_BAR2 = C.VI_ATTR_PXI_MEM_SIZE_BAR2
	ATTR_PXI_MEM_SIZE_BAR3 = C.VI_ATTR_PXI_MEM_SIZE_BAR3
	ATTR_PXI_MEM_SIZE_BAR4 = C.VI_ATTR_PXI_MEM_SIZE_BAR4
	ATTR_PXI_MEM_SIZE_BAR5 = C.VI_ATTR_PXI_MEM_SIZE_BAR5

	// Event Types
	EVENT_IO_COMPLETION    = C.VI_EVENT_IO_COMPLETION
	EVENT_TRIG             = C.VI_EVENT_TRIG
	EVENT_SERVICE_REQ      = C.VI_EVENT_SERVICE_REQ
	EVENT_CLEAR            = C.VI_EVENT_CLEAR
	EVENT_EXCEPTION        = C.VI_EVENT_EXCEPTION
	EVENT_GPIB_CIC         = C.VI_EVENT_GPIB_CIC
	EVENT_GPIB_TALK        = C.VI_EVENT_GPIB_TALK
	EVENT_GPIB_LISTEN      = C.VI_EVENT_GPIB_LISTEN
	EVENT_VXI_VME_SYSFAIL  = C.VI_EVENT_VXI_VME_SYSFAIL
	EVENT_VXI_VME_SYSRESET = C.VI_EVENT_VXI_VME_SYSRESET
	EVENT_VXI_SIGP         = C.VI_EVENT_VXI_SIGP
	EVENT_VXI_VME_INTR     = C.VI_EVENT_VXI_VME_INTR
	EVENT_PXI_INTR         = C.VI_EVENT_PXI_INTR
	EVENT_TCPIP_CONNECT    = C.VI_EVENT_TCPIP_CONNECT
	EVENT_USB_INTR         = C.VI_EVENT_USB_INTR

	ALL_ENABLED_EVENTS = C.VI_ALL_ENABLED_EVENTS

	// Completion and Error Codes
	SUCCESS_EVENT_EN       = C.VI_SUCCESS_EVENT_EN
	SUCCESS_EVENT_DIS      = C.VI_SUCCESS_EVENT_DIS
	SUCCESS_QUEUE_EMPTY    = C.VI_SUCCESS_QUEUE_EMPTY
	SUCCESS_TERM_CHAR      = C.VI_SUCCESS_TERM_CHAR
	SUCCESS_MAX_CNT        = C.VI_SUCCESS_MAX_CNT
	SUCCESS_DEV_NPRESENT   = C.VI_SUCCESS_DEV_NPRESENT
	SUCCESS_TRIG_MAPPED    = C.VI_SUCCESS_TRIG_MAPPED
	SUCCESS_QUEUE_NEMPTY   = C.VI_SUCCESS_QUEUE_NEMPTY
	SUCCESS_NCHAIN         = C.VI_SUCCESS_NCHAIN
	SUCCESS_NESTED_SHARED  = C.VI_SUCCESS_NESTED_SHARED
	SUCCESS_NESTED_EXCLUSI = C.VI_SUCCESS_NESTED_EXCLUSIVE
	SUCCESS_SYNC           = C.VI_SUCCESS_SYNC

	WARN_QUEUE_OVERFLOW  = C.VI_WARN_QUEUE_OVERFLOW
	WARN_CONFIG_NLOADED  = C.VI_WARN_CONFIG_NLOADED
	WARN_NULL_OBJECT     = C.VI_WARN_NULL_OBJECT
	WARN_NSUP_ATTR_STATE = C.VI_WARN_NSUP_ATTR_STATE
	WARN_UNKNOWN_STATUS  = C.VI_WARN_UNKNOWN_STATUS
	WARN_NSUP_BUF        = C.VI_WARN_NSUP_BUF
	WARN_EXT_FUNC_NIMPL  = C.VI_WARN_EXT_FUNC_NIMPL

	ERROR_SYSTEM_ERROR     = C.VI_ERROR_SYSTEM_ERROR
	ERROR_INV_OBJECT       = C.VI_ERROR_INV_OBJECT
	ERROR_RSRC_LOCKED      = C.VI_ERROR_RSRC_LOCKED
	ERROR_INV_EXPR         = C.VI_ERROR_INV_EXPR
	ERROR_RSRC_NFOUND      = C.VI_ERROR_RSRC_NFOUND
	ERROR_INV_RSRC_NAME    = C.VI_ERROR_INV_RSRC_NAME
	ERROR_INV_ACC_MODE     = C.VI_ERROR_INV_ACC_MODE
	ERROR_TMO              = C.VI_ERROR_TMO
	ERROR_CLOSING_FAILED   = C.VI_ERROR_CLOSING_FAILED
	ERROR_INV_DEGREE       = C.VI_ERROR_INV_DEGREE
	ERROR_INV_JOB_ID       = C.VI_ERROR_INV_JOB_ID
	ERROR_NSUP_ATTR        = C.VI_ERROR_NSUP_ATTR
	ERROR_NSUP_ATTR_STATE  = C.VI_ERROR_NSUP_ATTR_STATE
	ERROR_ATTR_READONLY    = C.VI_ERROR_ATTR_READONLY
	ERROR_INV_LOCK_TYPE    = C.VI_ERROR_INV_LOCK_TYPE
	ERROR_INV_ACCESS_KEY   = C.VI_ERROR_INV_ACCESS_KEY
	ERROR_INV_EVENT        = C.VI_ERROR_INV_EVENT
	ERROR_INV_MECH         = C.VI_ERROR_INV_MECH
	ERROR_HNDLR_NINSTALLED = C.VI_ERROR_HNDLR_NINSTALLED
	ERROR_INV_HNDLR_REF    = C.VI_ERROR_INV_HNDLR_REF
	ERROR_INV_CONTEXT      = C.VI_ERROR_INV_CONTEXT
	ERROR_QUEUE_OVERFLOW   = C.VI_ERROR_QUEUE_OVERFLOW
	ERROR_NENABLED         = C.VI_ERROR_NENABLED
	ERROR_ABORT            = C.VI_ERROR_ABORT
	ERROR_RAW_WR_PROT_VIOL = C.VI_ERROR_RAW_WR_PROT_VIOL
	ERROR_RAW_RD_PROT_VIOL = C.VI_ERROR_RAW_RD_PROT_VIOL
	ERROR_OUTP_PROT_VIOL   = C.VI_ERROR_OUTP_PROT_VIOL
	ERROR_INP_PROT_VIOL    = C.VI_ERROR_INP_PROT_VIOL
	ERROR_BERR             = C.VI_ERROR_BERR
	ERROR_IN_PROGRESS      = C.VI_ERROR_IN_PROGRESS
	ERROR_INV_SETUP        = C.VI_ERROR_INV_SETUP
	ERROR_QUEUE_ERROR      = C.VI_ERROR_QUEUE_ERROR
	ERROR_ALLOC            = C.VI_ERROR_ALLOC
	ERROR_INV_MASK         = C.VI_ERROR_INV_MASK
	ERROR_IO               = C.VI_ERROR_IO
	ERROR_INV_FMT          = C.VI_ERROR_INV_FMT
	ERROR_NSUP_FMT         = C.VI_ERROR_NSUP_FMT
	ERROR_LINE_IN_USE      = C.VI_ERROR_LINE_IN_USE
	//ERROR_LINE_NRESERVED   = C.VI_ERROR_LINE_NRESERVED
	ERROR_NSUP_MODE        = C.VI_ERROR_NSUP_MODE
	ERROR_SRQ_NOCCURRED    = C.VI_ERROR_SRQ_NOCCURRED
	ERROR_INV_SPACE        = C.VI_ERROR_INV_SPACE
	ERROR_INV_OFFSET       = C.VI_ERROR_INV_OFFSET
	ERROR_INV_WIDTH        = C.VI_ERROR_INV_WIDTH
	ERROR_NSUP_OFFSET      = C.VI_ERROR_NSUP_OFFSET
	ERROR_NSUP_VAR_WIDTH   = C.VI_ERROR_NSUP_VAR_WIDTH
	ERROR_WINDOW_NMAPPED   = C.VI_ERROR_WINDOW_NMAPPED
	ERROR_RESP_PENDING     = C.VI_ERROR_RESP_PENDING
	ERROR_NLISTENERS       = C.VI_ERROR_NLISTENERS
	ERROR_NCIC             = C.VI_ERROR_NCIC
	ERROR_NSYS_CNTLR       = C.VI_ERROR_NSYS_CNTLR
	ERROR_NSUP_OPER        = C.VI_ERROR_NSUP_OPER
	ERROR_INTR_PENDING     = C.VI_ERROR_INTR_PENDING
	ERROR_ASRL_PARITY      = C.VI_ERROR_ASRL_PARITY
	ERROR_ASRL_FRAMING     = C.VI_ERROR_ASRL_FRAMING
	ERROR_ASRL_OVERRUN     = C.VI_ERROR_ASRL_OVERRUN
	ERROR_TRIG_NMAPPED     = C.VI_ERROR_TRIG_NMAPPED
	ERROR_NSUP_ALIGN_OFFSE = C.VI_ERROR_NSUP_ALIGN_OFFSET
	ERROR_USER_BUF         = C.VI_ERROR_USER_BUF
	ERROR_RSRC_BUSY        = C.VI_ERROR_RSRC_BUSY
	ERROR_NSUP_WIDTH       = C.VI_ERROR_NSUP_WIDTH
	ERROR_INV_PARAMETER    = C.VI_ERROR_INV_PARAMETER
	ERROR_INV_PROT         = C.VI_ERROR_INV_PROT
	ERROR_INV_SIZE         = C.VI_ERROR_INV_SIZE
	ERROR_WINDOW_MAPPED    = C.VI_ERROR_WINDOW_MAPPED
	ERROR_NIMPL_OPER       = C.VI_ERROR_NIMPL_OPER
	ERROR_INV_LENGTH       = C.VI_ERROR_INV_LENGTH
	ERROR_INV_MODE         = C.VI_ERROR_INV_MODE
	ERROR_SESN_NLOCKED     = C.VI_ERROR_SESN_NLOCKED
	ERROR_MEM_NSHARED      = C.VI_ERROR_MEM_NSHARED
	ERROR_LIBRARY_NFOUND   = C.VI_ERROR_LIBRARY_NFOUND
	ERROR_NSUP_INTR        = C.VI_ERROR_NSUP_INTR
	ERROR_INV_LINE         = C.VI_ERROR_INV_LINE
	ERROR_FILE_ACCESS      = C.VI_ERROR_FILE_ACCESS
	ERROR_FILE_IO          = C.VI_ERROR_FILE_IO
	ERROR_NSUP_LINE        = C.VI_ERROR_NSUP_LINE
	ERROR_NSUP_MECH        = C.VI_ERROR_NSUP_MECH
	ERROR_INTF_NUM_NCONFIG = C.VI_ERROR_INTF_NUM_NCONFIG
	ERROR_CONN_LOST        = C.VI_ERROR_CONN_LOST
	ERROR_MACHINE_NAVAIL   = C.VI_ERROR_MACHINE_NAVAIL
	ERROR_NPERMISSION      = C.VI_ERROR_NPERMISSION

	// Other VISA Definitions
	FIND_BUFLEN = C.VI_FIND_BUFLEN

	INTF_GPIB     = C.VI_INTF_GPIB
	INTF_VXI      = C.VI_INTF_VXI
	INTF_GPIB_VXI = C.VI_INTF_GPIB_VXI
	INTF_ASRL     = C.VI_INTF_ASRL
	INTF_PXI      = C.VI_INTF_PXI
	INTF_TCPIP    = C.VI_INTF_TCPIP
	INTF_USB      = C.VI_INTF_USB

	PROT_NORMAL        = C.VI_PROT_NORMAL
	PROT_FDC           = C.VI_PROT_FDC
	PROT_HS488         = C.VI_PROT_HS488
	PROT_4882_STRS     = C.VI_PROT_4882_STRS
	PROT_USBTMC_VENDOR = C.VI_PROT_USBTMC_VENDOR

	FDC_NORMAL = C.VI_FDC_NORMAL
	FDC_STREAM = C.VI_FDC_STREAM

	LOCAL_SPACE     = C.VI_LOCAL_SPACE
	A16_SPACE       = C.VI_A16_SPACE
	A24_SPACE       = C.VI_A24_SPACE
	A32_SPACE       = C.VI_A32_SPACE
	A64_SPACE       = C.VI_A64_SPACE
	PXI_ALLOC_SPACE = C.VI_PXI_ALLOC_SPACE
	PXI_CFG_SPACE   = C.VI_PXI_CFG_SPACE
	PXI_BAR0_SPACE  = C.VI_PXI_BAR0_SPACE
	PXI_BAR1_SPACE  = C.VI_PXI_BAR1_SPACE
	PXI_BAR2_SPACE  = C.VI_PXI_BAR2_SPACE
	PXI_BAR3_SPACE  = C.VI_PXI_BAR3_SPACE
	PXI_BAR4_SPACE  = C.VI_PXI_BAR4_SPACE
	PXI_BAR5_SPACE  = C.VI_PXI_BAR5_SPACE
	OPAQUE_SPACE    = C.VI_OPAQUE_SPACE

	UNKNOWN_LA      = C.VI_UNKNOWN_LA
	UNKNOWN_SLOT    = C.VI_UNKNOWN_SLOT
	UNKNOWN_LEVEL   = C.VI_UNKNOWN_LEVEL
	UNKNOWN_CHASSIS = C.VI_UNKNOWN_CHASSIS

	QUEUE         = C.VI_QUEUE
	HNDLR         = C.VI_HNDLR
	SUSPEND_HNDLR = C.VI_SUSPEND_HNDLR
	ALL_MECH      = C.VI_ALL_MECH

	ANY_HNDLR = C.VI_ANY_HNDLR

	TRIG_ALL         = C.VI_TRIG_ALL
	TRIG_SW          = C.VI_TRIG_SW
	TRIG_TTL0        = C.VI_TRIG_TTL0
	TRIG_TTL1        = C.VI_TRIG_TTL1
	TRIG_TTL2        = C.VI_TRIG_TTL2
	TRIG_TTL3        = C.VI_TRIG_TTL3
	TRIG_TTL4        = C.VI_TRIG_TTL4
	TRIG_TTL5        = C.VI_TRIG_TTL5
	TRIG_TTL6        = C.VI_TRIG_TTL6
	TRIG_TTL7        = C.VI_TRIG_TTL7
	TRIG_ECL0        = C.VI_TRIG_ECL0
	TRIG_ECL1        = C.VI_TRIG_ECL1
	TRIG_ECL2        = C.VI_TRIG_ECL2
	TRIG_ECL3        = C.VI_TRIG_ECL3
	TRIG_ECL4        = C.VI_TRIG_ECL4
	TRIG_ECL5        = C.VI_TRIG_ECL5
	TRIG_STAR_SLOT1  = C.VI_TRIG_STAR_SLOT1
	TRIG_STAR_SLOT2  = C.VI_TRIG_STAR_SLOT2
	TRIG_STAR_SLOT3  = C.VI_TRIG_STAR_SLOT3
	TRIG_STAR_SLOT4  = C.VI_TRIG_STAR_SLOT4
	TRIG_STAR_SLOT5  = C.VI_TRIG_STAR_SLOT5
	TRIG_STAR_SLOT6  = C.VI_TRIG_STAR_SLOT6
	TRIG_STAR_SLOT7  = C.VI_TRIG_STAR_SLOT7
	TRIG_STAR_SLOT8  = C.VI_TRIG_STAR_SLOT8
	TRIG_STAR_SLOT9  = C.VI_TRIG_STAR_SLOT9
	TRIG_STAR_SLOT10 = C.VI_TRIG_STAR_SLOT10
	TRIG_STAR_SLOT11 = C.VI_TRIG_STAR_SLOT11
	TRIG_STAR_SLOT12 = C.VI_TRIG_STAR_SLOT12
	TRIG_STAR_INSTR  = C.VI_TRIG_STAR_INSTR
	TRIG_PANEL_IN    = C.VI_TRIG_PANEL_IN
	TRIG_PANEL_OUT   = C.VI_TRIG_PANEL_OUT
	TRIG_STAR_VXI0   = C.VI_TRIG_STAR_VXI0
	TRIG_STAR_VXI1   = C.VI_TRIG_STAR_VXI1
	TRIG_STAR_VXI2   = C.VI_TRIG_STAR_VXI2

	TRIG_PROT_DEFAULT   = C.VI_TRIG_PROT_DEFAULT
	TRIG_PROT_ON        = C.VI_TRIG_PROT_ON
	TRIG_PROT_OFF       = C.VI_TRIG_PROT_OFF
	TRIG_PROT_SYNC      = C.VI_TRIG_PROT_SYNC
	TRIG_PROT_RESERVE   = C.VI_TRIG_PROT_RESERVE
	TRIG_PROT_UNRESERVE = C.VI_TRIG_PROT_UNRESERVE

	READ_BUF           = C.VI_READ_BUF
	WRITE_BUF          = C.VI_WRITE_BUF
	READ_BUF_DISCARD   = C.VI_READ_BUF_DISCARD
	WRITE_BUF_DISCARD  = C.VI_WRITE_BUF_DISCARD
	IO_IN_BUF          = C.VI_IO_IN_BUF
	IO_OUT_BUF         = C.VI_IO_OUT_BUF
	IO_IN_BUF_DISCARD  = C.VI_IO_IN_BUF_DISCARD
	IO_OUT_BUF_DISCARD = C.VI_IO_OUT_BUF_DISCARD

	FLUSH_ON_ACCESS = C.VI_FLUSH_ON_ACCESS
	FLUSH_WHEN_FULL = C.VI_FLUSH_WHEN_FULL
	FLUSH_DISABLE   = C.VI_FLUSH_DISABLE

	NMAPPED              = C.VI_NMAPPED
	USE_OPERS            = C.VI_USE_OPERS
	DEREF_ADDR           = C.VI_DEREF_ADDR
	DEREF_ADDR_BYTE_SWAP = C.VI_DEREF_ADDR_BYTE_SWAP

	TMO_IMMEDIATE = C.VI_TMO_IMMEDIATE
	TMO_INFINITE  = C.VI_TMO_INFINITE

	NO_LOCK        = C.VI_NO_LOCK
	EXCLUSIVE_LOCK = C.VI_EXCLUSIVE_LOCK
	SHARED_LOCK    = C.VI_SHARED_LOCK
	LOAD_CONFIG    = C.VI_LOAD_CONFIG

	NO_SEC_ADDR = C.VI_NO_SEC_ADDR

	ASRL_PAR_NONE  = C.VI_ASRL_PAR_NONE
	ASRL_PAR_ODD   = C.VI_ASRL_PAR_ODD
	ASRL_PAR_EVEN  = C.VI_ASRL_PAR_EVEN
	ASRL_PAR_MARK  = C.VI_ASRL_PAR_MARK
	ASRL_PAR_SPACE = C.VI_ASRL_PAR_SPACE

	ASRL_STOP_ONE  = C.VI_ASRL_STOP_ONE
	ASRL_STOP_ONE5 = C.VI_ASRL_STOP_ONE5
	ASRL_STOP_TWO  = C.VI_ASRL_STOP_TWO

	ASRL_FLOW_NONE     = C.VI_ASRL_FLOW_NONE
	ASRL_FLOW_XON_XOFF = C.VI_ASRL_FLOW_XON_XOFF
	ASRL_FLOW_RTS_CTS  = C.VI_ASRL_FLOW_RTS_CTS
	ASRL_FLOW_DTR_DSR  = C.VI_ASRL_FLOW_DTR_DSR

	ASRL_END_NONE     = C.VI_ASRL_END_NONE
	ASRL_END_LAST_BIT = C.VI_ASRL_END_LAST_BIT
	ASRL_END_TERMCHAR = C.VI_ASRL_END_TERMCHAR
	ASRL_END_BREAK    = C.VI_ASRL_END_BREAK

	STATE_ASSERTED   = C.VI_STATE_ASSERTED
	STATE_UNASSERTED = C.VI_STATE_UNASSERTED
	STATE_UNKNOWN    = C.VI_STATE_UNKNOWN

	BIG_ENDIAN    = C.VI_BIG_ENDIAN
	LITTLE_ENDIAN = C.VI_LITTLE_ENDIAN

	DATA_PRIV  = C.VI_DATA_PRIV
	DATA_NPRIV = C.VI_DATA_NPRIV
	PROG_PRIV  = C.VI_PROG_PRIV
	PROG_NPRIV = C.VI_PROG_NPRIV
	BLCK_PRIV  = C.VI_BLCK_PRIV
	BLCK_NPRIV = C.VI_BLCK_NPRIV
	D64_PRIV   = C.VI_D64_PRIV
	D64_NPRIV  = C.VI_D64_NPRIV
	D64_2EVME  = C.VI_D64_2EVME
	D64_SST160 = C.VI_D64_SST160
	D64_SST267 = C.VI_D64_SST267
	D64_SST320 = C.VI_D64_SST320

	WIDTH_8  = C.VI_WIDTH_8
	WIDTH_16 = C.VI_WIDTH_16
	WIDTH_32 = C.VI_WIDTH_32
	WIDTH_64 = C.VI_WIDTH_64

	GPIB_REN_DEASSERT        = C.VI_GPIB_REN_DEASSERT
	GPIB_REN_ASSERT          = C.VI_GPIB_REN_ASSERT
	GPIB_REN_DEASSERT_GTL    = C.VI_GPIB_REN_DEASSERT_GTL
	GPIB_REN_ASSERT_ADDRESS  = C.VI_GPIB_REN_ASSERT_ADDRESS
	GPIB_REN_ASSERT_LLO      = C.VI_GPIB_REN_ASSERT_LLO
	GPIB_REN_ASSERT_ADDRESS_ = C.VI_GPIB_REN_ASSERT_ADDRESS_LLO
	GPIB_REN_ADDRESS_GTL     = C.VI_GPIB_REN_ADDRESS_GTL

	GPIB_ATN_DEASSERT        = C.VI_GPIB_ATN_DEASSERT
	GPIB_ATN_ASSERT          = C.VI_GPIB_ATN_ASSERT
	GPIB_ATN_DEASSERT_HANDSH = C.VI_GPIB_ATN_DEASSERT_HANDSHAKE
	GPIB_ATN_ASSERT_IMMEDIAT = C.VI_GPIB_ATN_ASSERT_IMMEDIATE

	GPIB_HS488_DISABLED = C.VI_GPIB_HS488_DISABLED
	GPIB_HS488_NIMPL    = C.VI_GPIB_HS488_NIMPL

	GPIB_UNADDRESSED = C.VI_GPIB_UNADDRESSED
	GPIB_TALKER      = C.VI_GPIB_TALKER
	GPIB_LISTENER    = C.VI_GPIB_LISTENER

	VXI_CMD16        = C.VI_VXI_CMD16
	VXI_CMD16_RESP16 = C.VI_VXI_CMD16_RESP16
	VXI_RESP16       = C.VI_VXI_RESP16
	VXI_CMD32        = C.VI_VXI_CMD32
	VXI_CMD32_RESP16 = C.VI_VXI_CMD32_RESP16
	VXI_CMD32_RESP32 = C.VI_VXI_CMD32_RESP32
	VXI_RESP32       = C.VI_VXI_RESP32

	ASSERT_SIGNAL       = C.VI_ASSERT_SIGNAL
	ASSERT_USE_ASSIGNED = C.VI_ASSERT_USE_ASSIGNED
	ASSERT_IRQ1         = C.VI_ASSERT_IRQ1
	ASSERT_IRQ2         = C.VI_ASSERT_IRQ2
	ASSERT_IRQ3         = C.VI_ASSERT_IRQ3
	ASSERT_IRQ4         = C.VI_ASSERT_IRQ4
	ASSERT_IRQ5         = C.VI_ASSERT_IRQ5
	ASSERT_IRQ6         = C.VI_ASSERT_IRQ6
	ASSERT_IRQ7         = C.VI_ASSERT_IRQ7

	UTIL_ASSERT_SYSRESET  = C.VI_UTIL_ASSERT_SYSRESET
	UTIL_ASSERT_SYSFAIL   = C.VI_UTIL_ASSERT_SYSFAIL
	UTIL_DEASSERT_SYSFAIL = C.VI_UTIL_DEASSERT_SYSFAIL

	VXI_CLASS_MEMORY   = C.VI_VXI_CLASS_MEMORY
	VXI_CLASS_EXTENDED = C.VI_VXI_CLASS_EXTENDED
	VXI_CLASS_MESSAGE  = C.VI_VXI_CLASS_MESSAGE
	VXI_CLASS_REGISTER = C.VI_VXI_CLASS_REGISTER
	VXI_CLASS_OTHER    = C.VI_VXI_CLASS_OTHER

	PXI_ADDR_NONE = C.VI_PXI_ADDR_NONE
	PXI_ADDR_MEM  = C.VI_PXI_ADDR_MEM
	PXI_ADDR_IO   = C.VI_PXI_ADDR_IO
	PXI_ADDR_CFG  = C.VI_PXI_ADDR_CFG

	TRIG_UNKNOWN = C.VI_TRIG_UNKNOWN

	PXI_LBUS_UNKNOWN         = C.VI_PXI_LBUS_UNKNOWN
	PXI_LBUS_NONE            = C.VI_PXI_LBUS_NONE
	PXI_LBUS_STAR_TRIG_BUS_0 = C.VI_PXI_LBUS_STAR_TRIG_BUS_0
	PXI_LBUS_STAR_TRIG_BUS_1 = C.VI_PXI_LBUS_STAR_TRIG_BUS_1
	PXI_LBUS_STAR_TRIG_BUS_2 = C.VI_PXI_LBUS_STAR_TRIG_BUS_2
	PXI_LBUS_STAR_TRIG_BUS_3 = C.VI_PXI_LBUS_STAR_TRIG_BUS_3
	PXI_LBUS_STAR_TRIG_BUS_4 = C.VI_PXI_LBUS_STAR_TRIG_BUS_4
	PXI_LBUS_STAR_TRIG_BUS_5 = C.VI_PXI_LBUS_STAR_TRIG_BUS_5
	PXI_LBUS_STAR_TRIG_BUS_6 = C.VI_PXI_LBUS_STAR_TRIG_BUS_6
	PXI_LBUS_STAR_TRIG_BUS_7 = C.VI_PXI_LBUS_STAR_TRIG_BUS_7
	PXI_LBUS_STAR_TRIG_BUS_8 = C.VI_PXI_LBUS_STAR_TRIG_BUS_8
	PXI_LBUS_STAR_TRIG_BUS_9 = C.VI_PXI_LBUS_STAR_TRIG_BUS_9
	PXI_STAR_TRIG_CONTROLLER = C.VI_PXI_STAR_TRIG_CONTROLLER

	// Backward Compatibility Macros
	ERROR_INV_SESSION    = C.VI_ERROR_INV_SESSION
	INFINITE             = C.VI_INFINITE
	NORMAL               = C.VI_NORMAL
	FDC                  = C.VI_FDC
	HS488                = C.VI_HS488
	ASRL488              = C.VI_ASRL488
	ASRL_IN_BUF          = C.VI_ASRL_IN_BUF
	ASRL_OUT_BUF         = C.VI_ASRL_OUT_BUF
	ASRL_IN_BUF_DISCARD  = C.VI_ASRL_IN_BUF_DISCARD
	ASRL_OUT_BUF_DISCARD = C.VI_ASRL_OUT_BUF_DISCARD

	// National Instruments
	INTF_RIO      = C.VI_INTF_RIO
	INTF_FIREWIRE = C.VI_INTF_FIREWIRE

	ATTR_SYNC_MXI_ALLOW_EN = C.VI_ATTR_SYNC_MXI_ALLOW_EN

	// This is for VXI SERVANT resources
	EVENT_VXI_DEV_CMD      = C.VI_EVENT_VXI_DEV_CMD
	ATTR_VXI_DEV_CMD_TYPE  = C.VI_ATTR_VXI_DEV_CMD_TYPE
	ATTR_VXI_DEV_CMD_VALUE = C.VI_ATTR_VXI_DEV_CMD_VALUE

	VXI_DEV_CMD_TYPE_16 = C.VI_VXI_DEV_CMD_TYPE_16
	VXI_DEV_CMD_TYPE_32 = C.VI_VXI_DEV_CMD_TYPE_32

	// mode values include VI_VXI_RESP16, VI_VXI_RESP32, and the next 2 values
	VXI_RESP_NONE       = C.VI_VXI_RESP_NONE
	VXI_RESP_PROT_ERROR = C.VI_VXI_RESP_PROT_ERROR

	// This is for VXI TTL Trigger routing
	ATTR_VXI_TRIG_LINES_EN = C.VI_ATTR_VXI_TRIG_LINES_EN
	ATTR_VXI_TRIG_DIR      = C.VI_ATTR_VXI_TRIG_DIR

	// This allows extended Serial support on Win32 and on NI ENET Serial products
	ATTR_ASRL_DISCARD_NULL  = C.VI_ATTR_ASRL_DISCARD_NULL
	ATTR_ASRL_CONNECTED     = C.VI_ATTR_ASRL_CONNECTED
	ATTR_ASRL_BREAK_STATE   = C.VI_ATTR_ASRL_BREAK_STATE
	ATTR_ASRL_BREAK_LEN     = C.VI_ATTR_ASRL_BREAK_LEN
	ATTR_ASRL_ALLOW_TRANSMI = C.VI_ATTR_ASRL_ALLOW_TRANSMIT
	ATTR_ASRL_WIRE_MODE     = C.VI_ATTR_ASRL_WIRE_MODE

	ASRL_WIRE_485_4         = C.VI_ASRL_WIRE_485_4
	ASRL_WIRE_485_2_DTR_ECH = C.VI_ASRL_WIRE_485_2_DTR_ECHO
	ASRL_WIRE_485_2_DTR_CTR = C.VI_ASRL_WIRE_485_2_DTR_CTRL
	ASRL_WIRE_485_2_AUTO    = C.VI_ASRL_WIRE_485_2_AUTO
	ASRL_WIRE_232_DTE       = C.VI_ASRL_WIRE_232_DTE
	ASRL_WIRE_232_DCE       = C.VI_ASRL_WIRE_232_DCE
	ASRL_WIRE_232_AUTO      = C.VI_ASRL_WIRE_232_AUTO

	EVENT_ASRL_BREAK    = C.VI_EVENT_ASRL_BREAK
	EVENT_ASRL_CTS      = C.VI_EVENT_ASRL_CTS
	EVENT_ASRL_DSR      = C.VI_EVENT_ASRL_DSR
	EVENT_ASRL_DCD      = C.VI_EVENT_ASRL_DCD
	EVENT_ASRL_RI       = C.VI_EVENT_ASRL_RI
	EVENT_ASRL_CHAR     = C.VI_EVENT_ASRL_CHAR
	EVENT_ASRL_TERMCHAR = C.VI_EVENT_ASRL_TERMCHAR

	SUCCESS = C.VI_SUCCESS

	// Other VISA Definitions

	NULL = C.VI_NULL

	TRUE  = C.VI_TRUE
	FALSE = C.VI_FALSE

// #if defined(NIVISA_PXI) || defined(PXISAVISA_PXI)
// 	VI_ATTR_PXI_USE_PREALLOC_POOL = C.VI_ATTR_PXI_USE_PREALLOC_POOL
// #endif

// #if defined(NIVISA_USB)
// 	VI_ATTR_USB_BULK_OUT_PIPE = C.VI_ATTR_USB_BULK_OUT_PIPE
// 	VI_ATTR_USB_BULK_IN_PIPE = C.VI_ATTR_USB_BULK_IN_PIPE
// 	VI_ATTR_USB_INTR_IN_PIPE = C.VI_ATTR_USB_INTR_IN_PIPE
// 	VI_ATTR_USB_CLASS = C.VI_ATTR_USB_CLASS
// 	VI_ATTR_USB_SUBCLASS = C.VI_ATTR_USB_SUBCLASS
// 	VI_ATTR_USB_ALT_SETTING = C.VI_ATTR_USB_ALT_SETTING
// 	VI_ATTR_USB_END_IN = C.VI_ATTR_USB_END_IN
// 	VI_ATTR_USB_NUM_INTFCS = C.VI_ATTR_USB_NUM_INTFCS
// 	VI_ATTR_USB_NUM_PIPES = C.VI_ATTR_USB_NUM_PIPES
// 	VI_ATTR_USB_BULK_OUT_STATUS = C.VI_ATTR_USB_BULK_OUT_STATUS
// 	VI_ATTR_USB_BULK_IN_STATUS = C.VI_ATTR_USB_BULK_IN_STATUS
// 	VI_ATTR_USB_INTR_IN_STATUS = C.VI_ATTR_USB_INTR_IN_STATUS
// 	VI_ATTR_USB_CTRL_PIPE = C.VI_ATTR_USB_CTRL_PIPE

// 	VI_USB_PIPE_STATE_UNKNOWN = C.VI_USB_PIPE_STATE_UNKNOWN
// 	VI_USB_PIPE_READY = C.VI_USB_PIPE_READY
// 	VI_USB_PIPE_STALLED = C.VI_USB_PIPE_STALLED

// 	VI_USB_END_NONE = C.VI_USB_END_NONE
// 	VI_USB_END_SHORT = C.VI_USB_END_SHORT
// 	VI_USB_END_SHORT_OR_COUNT = C.VI_USB_END_SHORT_OR_COUNT
// #endif
)
