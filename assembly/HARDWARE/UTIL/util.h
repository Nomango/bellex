#ifndef __UTIL_H__
#define __UTIL_H__

#include "delay.h"
#include "sys.h"
#include "oled.h"
#include "bmp.h"
#include "key.h"
#include "key_send.h"
#include "ds1302.h"
#include "usart.h"
//#include "string.h"
#include "timer.h"
#include "time.h"
//#include "handle.h"
#include "request.h"
#include "bell.c"

int strlen(const char* str);
int strcpy(char* dest, const char* src);
int strncpy(char* dest, const char* src, int length);
int strcmp(const char* str1, const char* str2);

#endif
