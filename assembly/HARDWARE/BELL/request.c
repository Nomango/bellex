#include "request.h"

char g_Buffer[261] = { 0 };
int g_BufferSize = 0;

enum BellRequestType g_RequestType = 0;

void MakeRequest(const char* request_str)
{
    int pos = 0, len = 0;
		char* heart_beat = "heart_beat";
    // 添加前标志位
    g_Buffer[0] = 0xFF;
    g_Buffer[1] = 0xFF;
    g_Buffer[2] = 0;
    pos += 3;
		if( strcmp(request_str,heart_beat) == 0)
			//GetBellStatus()内容为空
			len = sprintf(&g_Buffer[pos], "id:%s;code:%s;req:%s;data:status:%s;", GetBellID(), GetBellCode(), request_str, GetBellStatus());
		else
			len = sprintf(&g_Buffer[pos], "id:%s;code:%s;req:%s;data:;", GetBellID(), GetBellCode(), request_str);
		
    pos += len;

    // 在数据包开头添加数据包大小
    g_Buffer[2] = (char)len;

    // 添加后标志位
    g_Buffer[pos] = 0xFF;
    g_Buffer[pos + 1] = 0xFE;

    pos += 2;
    g_BufferSize = pos;

}

void SendConnect(void)
{
    MakeRequest("connect");
    UART2_Send_Array((unsigned char*)g_Buffer, (unsigned char)g_BufferSize);
    g_RequestType = TYPE_CONNECT;
}

void SendHeartBeat(void)
{
    MakeRequest("heart_beat");
    UART2_Send_Array((unsigned char*)g_Buffer, (unsigned char)g_BufferSize);
    g_RequestType = TYPE_HEART_BEAT;
}

void SendTimeRequest(void)
{
    MakeRequest("proof_time");
    UART2_Send_Array((unsigned char*)g_Buffer, (unsigned char)g_BufferSize);
    g_RequestType = TYPE_REQUEST_TIME;
}

void SendScheduleTimeRequest(void)
{
    MakeRequest("get_schedule");
    UART2_Send_Array((unsigned char*)g_Buffer, (unsigned char)g_BufferSize);
    g_RequestType = TYPE_REQUEST_TIME;
}

void SendNTPRequest(void)
{
    struct NtpPacket packet = DefaultPacket();
    UART2_Send_Array((unsigned char *)&packet, sizeof(struct NtpPacket));
    g_RequestType = TYPE_NTP_REQUEST;
}
