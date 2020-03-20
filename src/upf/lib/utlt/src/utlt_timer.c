#include "utlt_timer.h"

#include <string.h>
#include <stdio.h>

#include "utlt_debug.h"
#include "utlt_pool.h"

static int TimerCmpFunc(ListNode *pnode1, ListNode *pnode2);

typedef struct _TimerBlk {
    ListNode        node;
    TimerList       *timerList;
    
    int             type;
    int             isRunning;
    uint32_t        expireTime;
    uint32_t        duration;
    
    ExpireFunc      expireFunc;
    uintptr_t       param[6];
} TimerBlk;

PoolDeclare(timerPool, TimerBlk, MAX_NUM_OF_TIMER);

static int TimerCmpFunc(ListNode *pnode1, ListNode *pnode2) {
    TimerBlk *tm1 = (TimerBlk *)pnode1;
    TimerBlk *tm2 = (TimerBlk *)pnode2;

    return (tm1->expireTime < tm2->expireTime ? -1 : 1);
}

Status TimerPoolInit(void) {
    PoolInit(&timerPool, MAX_NUM_OF_TIMER);

    return STATUS_OK;
}

Status TimerFinal(void) {
    if (PoolCap(&timerPool) != PoolSize(&timerPool))
        UTLT_Error("%d not freed in timerPool[%d]",
                    PoolCap(&timerPool) - PoolSize(&timerPool),
                    PoolCap(&timerPool));
    
    PoolTerminate(&timerPool);

    return STATUS_OK;
}

uint32_t TimerGetPoolSize(void) {
    // The number of available space in this pool
    return PoolSize(&timerPool);
}

void TimerListInit(TimerList *tmList) {
    memset(tmList, 0x00, sizeof(TimerList));
    ListInit(&tmList->active);
    ListInit(&tmList->idle);
    return;
}

// Check expire time and update active and idle list
Status TimerExpireCheck(TimerList *tmList, uintptr_t data) {
    uint32_t curTime = TimeMsec(TimeNow());
    TimerBlk *tm = ListFirst(&(tmList->active));

    while (tm) {
        if (tm->expireTime < curTime) {
            tm->expireFunc(data, tm->param);
            
            if (tm->isRunning) {
                ListRemove(&(tmList->active), tm);

                if (tm->type == TIMER_TYPE_PERIOD) {
                    tm->expireTime = curTime + tm->duration;
                    
                    ListInsertSorted(&(tmList->active), tm, TimerCmpFunc);
                } else {
                    ListInsertSorted(&(tmList->idle), tm, TimerCmpFunc);
                    
                    tm->isRunning = 0;
                }
            }
            tm = ListFirst(&(tmList->active));
        } else {
            break;
        }
    }
    return STATUS_OK;
}

Status TimerStart(TimerBlkID id) {
    uint32_t curTime = TimeMsec(TimeNow());
    TimerBlk *tm = (TimerBlk *)id;

    if (tm->isRunning)
        ListRemove(&(tm->timerList->active), tm);
    else
        ListRemove(&(tm->timerList->idle), tm);

    tm->expireTime = curTime + tm->duration;

    ListInsertSorted(&(tm->timerList->active), tm, TimerCmpFunc);
    
    tm->isRunning = 1;

    return STATUS_OK;
}

Status TimerStop(TimerBlkID id) {
    TimerBlk *tm = (TimerBlk *)id;

    if (tm->isRunning) {
        ListRemove(&(tm->timerList->active), tm);
        ListInsertSorted(&(tm->timerList->idle), tm, TimerCmpFunc);

        tm->isRunning = 0;
    }

    return STATUS_OK;
}

TimerBlkID TimerCreate(TimerList *tmList, int type, uint32_t duration, ExpireFunc expireFunc) {
    TimerBlk *tm = NULL;
    
    PoolAlloc(&timerPool, tm);
    UTLT_Assert(tm, return (TimerBlkID)NULL, "The pool of timer create is empty");
    
    memset((char*)tm, 0x00, sizeof(TimerBlk));

    tm->timerList = tmList;

    ListInsertSorted(&(tm->timerList->idle), tm, TimerCmpFunc);

    tm->type = type;
    tm->duration = duration;
    tm->expireFunc = expireFunc;

    return (TimerBlkID)tm;
}

void TimerDelete(TimerBlkID id) {
    TimerBlk *tm = (TimerBlk *)id;

    if (tm->isRunning)
        ListRemove(&(tm->timerList->active), tm);
    else
        ListRemove(&(tm->timerList->idle), tm);

    PoolFree(&timerPool, tm);

    return;
}

Status TimerSet(int paramID, TimerBlkID id, uintptr_t param) {
    TimerBlk *tm = (TimerBlk *)id;

    UTLT_Assert(paramID >= 0 && paramID < 6, return STATUS_ERROR, "Wrong paramID for setting timer parameter");
    tm->param[paramID] = param;
    
    return STATUS_OK;
}