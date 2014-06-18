package gopjnath

/*
#cgo pkg-config: libpjproject
#include <pjlib.h>
#include <pjlib-util.h>
#include <pjnath.h>
*/
import "C"

import (
    "syscall"
    "unsafe"
    )

type cachingPool unsafe.Pointer

var (
    cp C.struct_pj_caching_pool
    pool C.struct_pj_pool_t
    stunConfig *stunConfig
    )

func init() {
    C.pj_init()
    C.pjlib_util_init()
    C.pjnath_init()
    C.pj_caching_pool_init(cp, &C.pj_pool_factory_default_policy, 0)
    pool = C.pj_pool_create(&cp._factory,"",512,512,nil)
    
}

type TransportOp int

const (
    TransportOpStateInit        = TransportOp(C.PJ_ICE_STRANS_OP_INIT)
    TransportOpStateNegotiation = TransportOp(C.PJ_ICE_STRANS_OP_NEGOTIATION)
    TransportOpStateKeepAlive   = TransportOp(C.PJ_ICE_STRANS_OP_KEEP_ALIVE)
    )

type TransportState int

const (
    TransportStateNull      = TransportState(C.PJ_ICE_STRANS_STATE_NULL)
    TransportStateInit      = TransportState(C.PJ_ICE_STRANS_STATE_INIT)
    TransportStateReady     = TransportState(C.PJ_ICE_STRANS_STATE_READY)
    TransportStateSessReady = TransportState(C.PJ_ICE_STRANS_STATE_SESS_READY)
    TransportStateNego      = TransportState(C.PJ_ICE_STRANS_STATE_NEGO)
    TransportStateRunning   = TransportState(C.PJ_ICE_STRANS_STATE_RUNNING)
    TransportStateFailed    = TransportState(C.PJ_ICE_STRANS_STATE_FAILED)
    )

type IceSessRole int

const (
    IceSessRoleUnknown     = IceSessRole(C.PJ_ICE_SESS_ROLE_UNKNOWN)
    IceSessRoleControlled  = IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLED)
    IceSessRoleControlling = IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLING)
    )

type QosType int 

const (
    QosTypeBestEffort = QosType(C.PJ_QOS_TYPE_BEST_EFFORT)
    QosTypeBackground = QosType(C.PJ_QOS_TYPE_BACKGROUND)
    QosTypeVideo      = QosType(C.PJ_QOS_TYPE_VIDEO)
    QosTypeVoice      = QosType(C.PJ_QOS_TYPE_VOICE)
    QosTypeControl    = QosType(C.PJ_QOS_TYPE_CONTROL)
    )

func casterr(fromcgo C.pj_status_t) error {
    errno := fromcgo.(syscall.Errno)
    if !ok {
        return fromcgo
    }
    pjstatus := pjStatus(errno)
    return errno
}
