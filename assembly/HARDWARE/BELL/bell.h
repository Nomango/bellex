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

const char* GetBellID();
const char* GetBellCode();

void SetBellCode(const char* code, int size);
void HandleConncetResponse(unsigned char* recv, unsigned char size);
void HandleTimeResponse(unsigned char* recv, unsigned char size);

#endif
