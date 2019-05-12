#ifndef __BELL_H__
#define __BELL_H__

#include "delay.h"
#include "sys.h"
#include "usart.h"

//#include "oled.h"
//#include "bmp.h"
//#include "key.h"
//#include "key_send.h"
//#include "ds1302.h"
//#include "ntp.h"
//#include "handle.h"
//#include "request.h"
////#include "bell.h"
//#include "timer.h"
//#include "time.h" 

#include "string.h"

#define BELL_ID_LENGTH 8
#define BELL_CODE_LENGTH 8
extern unsigned char sch[100];
extern unsigned char SCH[100];
const char* GetBellID(void);
const char* GetBellCode(void);
char* GetBellStatus(void);
void SetBellCode(const char* code, int size);
void HandleConncetResponse(unsigned char* recv, unsigned char size);
void HandleTimeResponse(unsigned char* recv, unsigned char size);
void HandleHeartResponse(unsigned char* recv, unsigned char size);
void Handle_Schudeule_Rec(unsigned char* recv, unsigned char size);
void Handle_Regular_Belling(unsigned char* recv, unsigned char size);
int bell_on(unsigned char* recv, unsigned char size);
void restart(void);
#endif
