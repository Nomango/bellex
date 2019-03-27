#ifndef _MAIN_H
#define _MAIN_H

#include "delay.h"
#include "sys.h"
#include "usart.h"
#include "oled.h"
#include "bmp.h"
#include "key.h"
#include "key_send.h"
#include "ds1302.h"
#include "ntp.h"
#include "handle.h"
#include "request.h"
#include "bell.h"
#include "timer.h"
#include "time.h" 
#include "string.h"


void oled_time(unsigned char net);

struct Server_schedule{
	unsigned char min;
	unsigned char hour;
};



#endif 
