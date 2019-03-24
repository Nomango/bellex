#include "handle.h"
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

void HandleNtpResponse(unsigned char* recv, unsigned char size)
{
    int tx_tm_s;
	time_t unix;
	struct tm *tmbuf;

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

    printf("%d %d %d %d %d %d\r\n",
        1900 + tmbuf->tm_year,	// 年份
        1 + tmbuf->tm_mon,		// 月份
        1 + tmbuf->tm_wday,		// 周
        tmbuf->tm_mday,			// 天
        8 + tmbuf->tm_hour,		// 小时
        tmbuf->tm_min,
        tmbuf->tm_sec);
}
