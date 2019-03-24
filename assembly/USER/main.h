#ifndef _MAIN_H
#define _MAIN_H

#include "delay.h"
#include "sys.h"
#include "oled.h"
#include "bmp.h"
#include "key.h"
#include "key_send.h"
#include "ds1302.h"
#include "usart.h"
#include "string.h"
#include "timer.h"
#include "time.h"

void oled_time(unsigned char net);

struct Server_schedule{
	unsigned char min;
	unsigned char hour;
};



#endif 
