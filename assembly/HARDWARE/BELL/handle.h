#ifndef __BELL_HANDLE_H__
#define __BELL_HANDLE_H__

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
#include "request.h"
#include "bell.h"
#include "timer.h"
#include "time.h" 
#include "ntp.h"
#include "ds1302.h"
#include "string.h"

void HandleNtpResponse(unsigned char* recv, unsigned char size);



#endif
