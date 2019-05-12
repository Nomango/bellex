#include "handle.h"


void HandleNtpResponse(unsigned char* recv, unsigned char size)
{
	int tx_tm_s;
	time_t unix;
	struct tm *tmbuf;
	unsigned int NTP_time[7];
	

    // 处理这个请求必须满足 48 个字节
    if (size != 48)
        return;

    // 从接收的 buffer 中计算 txTmSec 变量
    tx_tm_s = ((int)recv[40]) << 24;
    tx_tm_s += ((int)recv[41]) << 16;
    tx_tm_s += ((int)recv[42]) << 8;
    tx_tm_s += ((int)recv[43]);

    // 计算时间戳
    unix = (time_t)(((uint64_t)tx_tm_s) - 2208988800);
    tmbuf = localtime(&unix);
		
		if(tmbuf->tm_wday == 0)
			tmbuf->tm_wday=7;
		
		NTP_time[0] = tmbuf->tm_sec;
		NTP_time[1] = tmbuf->tm_min;
		NTP_time[2] = tmbuf->tm_hour + 8;
		NTP_time[3] = tmbuf->tm_mday;
		NTP_time[4] = tmbuf->tm_mon + 1;
		NTP_time[5] = tmbuf->tm_wday;
		NTP_time[6] = tmbuf->tm_year - 100;
		
		ds1302_init(write, NTP_time);
		
//    printf("%d %d %d %d %d %d %d\r\n",
//        1900 + tmbuf->tm_year,	// 年份
//        1 + tmbuf->tm_mon,		// 月份
//        tmbuf->tm_wday,		// 周
//        tmbuf->tm_mday,			// 天
//        8 + tmbuf->tm_hour,		// 小时
//        tmbuf->tm_min,
//        tmbuf->tm_sec
//				);
}
