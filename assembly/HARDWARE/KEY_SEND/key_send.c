#include "key_send.h"
#include "../BELL/request.h"
#include "usart.h"	
unsigned char AT_buf[64];
unsigned char key;
extern u8 HMI;
extern unsigned char now[7];
char hou_m;
char min_m;

void HMI_DISPALY(void)
{
	u8 BUF[100];
	u8 buf[100];
	early_time();
	if ((now[1]/10) == 0)
	{
		sprintf(BUF,"t3.txt=\"%d\:0%d\"",now[2],now[1]);//采样频率1输出
		HMISends_HMI(BUF);
		HMISendb(0xff);		
		delay_ms(300);
	}
	
	else
	{
		sprintf(BUF,"t3.txt=\"%d\:%d\"",now[2],now[1]);//采样频率1输出
		HMISends_HMI(BUF);
		HMISendb(0xff);		
		delay_ms(300);
	}

	if ((min_m/10) == 0)
	{
		sprintf(buf,"t4.txt=\"%d\:0%d\"",hou_m,min_m);//采样频率1输出
		HMISends_HMI(buf);
		HMISendb(0xff);		
		delay_ms(300);
	}
	
	else
	{
		sprintf(buf,"t4.txt=\"%d\:%d\"",hou_m,min_m);//采样频率1输出
		HMISends_HMI(buf);
		HMISendb(0xff);		
		delay_ms(300);
	}

}

void early_time(void)
{
	int i, j, jj, t=0, tt=0, cnt=1;
	char hou[100];
	char min[100];
	int tim[100];
	int MIN;
	i = strlen(SCH);
//	printf("SCH%d\r\n",SCH[i]);
//	printf("SCH1%d\r\n",SCH[i+1]);
	for(j = 0; j < i; j+=2)
	{
		//printf("cnt=%d\r\n",cnt++);
		if (((SCH[j] - now[2]) > 0) || (((SCH[j+1] - now[1]) > 0) && ((SCH[j] - now[2]) == 0)))
		{
			hou[t++] = SCH[j] - now[2];
			min[tt++] = SCH[j+1] - now[1];
//			printf("hou1%d\r\n",SCH[j]);
//			printf("hou2%d\r\n",SCH[j+1]);
//			printf("hou1%d\r\n",hou[t-1]);
//			printf("hou2%d\r\n",SCH[j] - now[2]);
//			printf("min1%d\r\n",min[t-1]);
//			printf("min2%d\r\n",SCH[j+1] - now[1]);
		}
	}
	cnt = 0;
	j = 0;
//	printf("t:%d\r\n",t);
	if (t > 1)
	{
		MIN = hou[0]*60 + min[0];
		jj = 0;
		for(j = 0; j < t; j++)
		{
			tim[j] = hou[j]*60 + min[j];
			
			if (tim[j] <= MIN)
			{
				MIN = tim[j];
				jj = j;
			}
			
			//printf("tim%d\r\n",tim[j]);
		}
		
		hou_m = hou[jj] + now[2];
		min_m = min[jj] + now[1];
		//printf("hou_m%d\r\n",hou_m);
		//printf("min_m%d\r\n",min_m);
	}
	
	else if (t == 1)
	{
				hou_m = hou[0] + now[2];
				min_m = min[0] + now[1];
			//printf("hou_m%d\r\n",hou_m);
		//printf("min_m%d\r\n",min_m);
	}

}

void HMI_REC(void)
{
	if (HMI == 0xa0)
	{
		SendTimeRequest();
		delay_ms(200);
		HMI = 0;
	}
	
	else if (HMI == 0xd0)
	{
		SendScheduleTimeRequest();
		delay_ms(200);
		HMI = 0;
	}
}
void wifi_connect(void)
{
	unsigned char buf[100];
	key = KEY_Scan(0);

	if (key == KEY2_PRES)
	{
		delay_ms(100);
		SendConnect();
	}

	if (key == KEY3_PRES)
	{	
		delay_ms(100);
//		SendHeartBeat();
		SendTimeRequest();
//		SendNTPRequest();
	}

	if (key == KEY4_PRES)
	{
//		delay_ms(100);
//		sprintf(buf, "+++");	
//		HMISends_0(buf);
//		delay_ms(500);
//		sprintf(buf, "AT+CIPCLOSE\r\n");		//目标服务器得到IP地址、端口
//		HMISends_0(buf);
//		delay_ms(500);
		SendScheduleTimeRequest();
	}
}
