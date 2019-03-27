#include "main.h"


struct Server_schedule schedule[30];

int main(void)
{

	u16 i, j;

	unsigned int time[7] = {50, 01, 02, 02, 03, 4, 12};
	//unsigned char buf[48];
	unsigned char num, factor;
	unsigned int now_time[7];

	delay_init();		  //延时函数初始化
	NVIC_Configuration(); //设置NVIC中断分组2:2位抢占优先级，2位响应优先级
	uart2_init(115200);
	KEY_Init();
	ds1302_GPIO_Configuration();	//1302模拟时序引脚初始化
	ds1302_init(write, time);		//设置1302的初始时间
	OLED_Init();					//初始化OLED
	OLED_Clear();
	bell_init();
	TIM3_Int_Init(5000 - 1, 7200 - 1);

	//	sprintf(buf,"AT+CIPMODE=1\r\n");		//设置WiFi模块为STA模式
	//	HMISends_0(buf);						//利用串口1配置模块
	//	delay_ms(500);

	//	sprintf(buf,"AT+CWJAP=\"iPhone\",\"0987654321\"\r\n");		//连接的路由器名称和密码
	//	sprintf(buf,"AT+CWJAP=\"NAO\",\"nao12345\"\r\n");
	//	HMISends_0(buf);
	//	delay_ms(5000);

	while (1)
	{
		wifi_connect();
		ds1302_data(read);
		oled_time(0);

		if (g_StreamEndFlag)
		{
			switch (g_RequestType)
			{
			case TYPE_CONNECT:
				HandleConncetResponse(g_RecvBuffer, g_RecvBufferSize);
				break;
			case TYPE_HEART_BEAT:
				break;
			case TYPE_REQUEST_TIME:
				HandleTimeResponse(g_RecvBuffer, g_RecvBufferSize);
				break;
			case TYPE_NTP_REQUEST:
				HandleNtpResponse(g_RecvBuffer, g_RecvBufferSize);
				break;
			default:
				break;
			}

			delay_ms(1000);

			// 重置所有标志
			g_StreamEndFlag = 0;
			g_RecvBufferSize = 0;
		}

		// if (I == 9)
		// {
		// 	for (i = I, j = 0; i < sum; i++, j++)
		// 		rec_schedule[j] = g_RecvBuffer[i];

		// 	num = j / 4;
		// 	I = 0;

		// 	for (i = 0; i < num; i++)
		// 	{
		// 		schedule[i].hour = rec_schedule[i * 4] * 10 + rec_schedule[i * 4 + 1] - 16;
		// 		schedule[i].min = rec_schedule[i * 4 + 2] * 10 + rec_schedule[i * 4 + 3] - 16;

		// 		USART_SendData(USART2, schedule[i].hour);
		// 		delay_ms(250);
		// 		USART_SendData(USART2, schedule[i].min);
		// 		delay_ms(250);
		// 	}
		// }

		// else if (I == 13)
		// {
		// 	for (i = I, j = 0; i < sum; i++, j++)
		// 		REC[j] = g_RecvBuffer[i];
		// 	if (j == 14)
		// 	{
		// 		for (i = 0, j = 0; i < 7; i++, j = j + 2)
		// 			now_time[i] = (((int)((REC[j]) - 48)) * 10) + ((int)((REC[j + 1]) - 48));
		// 		ds1302_init(write, now_time);
		// 	}
		// 	I = 0;
		// }

		//		for( i=0 ; i<num ; i++ ){
		//			if( schedule[i].hour == now[2] && schedule[i].min == now[1] && now[0]>0 && now[0]<5){
		////				USART_SendData(USART2,1);
		////				delay_ms(250);
		//			}
		//			else
		//				bell_2=1;
		//		}
	}
}

void oled_time(unsigned char net)
{
	if (net == 0)
	{
		//		OLED_Clear();
		OLED_ShowCHinese(28, 0, 0);  //年
		OLED_ShowCHinese(70, 0, 1);  //月
		OLED_ShowCHinese(112, 0, 2); //日
		OLED_ShowCHinese(28, 4, 3);  //时
		OLED_ShowCHinese(70, 4, 4);  //分
		OLED_ShowCHinese(112, 4, 5); //秒

		OLED_ShowNum(0, 0, now[6], 3, 16);
		OLED_ShowNum(42, 0, now[4], 3, 16);
		OLED_ShowNum(86, 0, now[3], 3, 16);
		OLED_ShowNum(0, 4, now[2], 3, 16);
		OLED_ShowNum(42, 4, now[1], 3, 16);
		OLED_ShowNum(86, 4, now[0], 3, 16);
	}
	if (net == 1)
	{
		OLED_Clear();
		OLED_ShowNum(0, 0, schedule[0].hour, 3, 12);
		OLED_ShowChar(22, 0, ':', 12);
		OLED_ShowNum(28, 0, schedule[0].min, 3, 12);

		OLED_ShowNum(0, 2, schedule[1].hour, 3, 12);
		OLED_ShowChar(22, 2, ':', 12);
		OLED_ShowNum(28, 2, schedule[1].min, 3, 12);

		OLED_ShowNum(60, 0, schedule[2].hour, 3, 12);
		OLED_ShowChar(82, 0, ':', 12);
		OLED_ShowNum(88, 0, schedule[2].min, 3, 12);

		OLED_ShowNum(60, 2, schedule[3].hour, 3, 12);
		OLED_ShowChar(82, 2, ':', 12);
		OLED_ShowNum(88, 2, schedule[3].min, 3, 12);
	}
}
