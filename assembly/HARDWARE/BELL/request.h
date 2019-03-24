#ifndef __BELL_REQUEST_H__
#define __BELL_REQUEST_H__

enum BellRequestType
{
    TYPE_CONNECT = 1,
    TYPE_HEART_BEAT,
    TYPE_REQUEST_TIME,
    TYPE_NTP_REQUEST,
};

extern BellRequestType g_RequestType;

// 请求连接服务器
void SendConnect();
// 发送心跳包
void SendHeartBeat();
// 发送时间校对请求
void SendTimeRequest();
// 发送 NTP 数据包
void SendNTPRequest();

#endif
