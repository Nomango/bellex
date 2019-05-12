#ifndef __BELL_REQUEST_H__
#define __BELL_REQUEST_H__

#include "sys.h"
#include "usart.h"
#include "ntp.h"
#include "bell.h"
#include "string.h"

enum BellRequestType
{
    TYPE_CONNECT = 1,
    TYPE_HEART_BEAT,
    TYPE_REQUEST_TIME,
    TYPE_NTP_REQUEST,
		TYPE_SCHEDULE_REC,
		TYPE_REGULAR_REC,
		TYPE_BELL_REC,
	
};


extern enum BellRequestType g_RequestType;
void SendScheduleTimeRequest(void);
// 请求连接服务器
void SendConnect(void);
// 发送心跳包
void SendHeartBeat(void);
// 发送时间校对请求
void SendTimeRequest(void);
// 发送 NTP 数据包
void SendNTPRequest(void);

#endif
