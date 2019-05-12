#include "main.h"


struct Server_schedule schedule[30];
u16 bell_flag;
int main(void)
{

	u16 i, j;
	unsigned int time[7] = {50, 01, 02, 02, 03, 4, 12};
	unsigned char buf[100];
//	unsigned char num, factor;
//	unsigned int now_time[7];

	delay_init();		  //��ʱ������ʼ��
	NVIC_Configuration(); //����NVIC�жϷ���2:2λ��ռ���ȼ���2λ��Ӧ���ȼ�
	uart2_init(115200);
	
	KEY_Init();
	LED_Init();
	ds1302_GPIO_Configuration();	//1302ģ��ʱ�����ų�ʼ��
	ds1302_init(write, time);		//����1302�ĳ�ʼʱ��
	OLED_Init();					//��ʼ��OLED
	OLED_Clear();
	bell_init();
	TIM3_Int_Init(5000 - 1, 7200 - 1);
	TIM4_Int_Init(10000 - 1, 7200 - 1);
	
	for(i=0; i<8; i++)
		delay_ms(1000);
	
	
//		sprintf(buf,"AT+CIPMODE=1\r\n");		//����WiFiģ��ΪSTAģʽ
//		HMISends_0(buf);						//���ô���1����ģ��
//		delay_ms(500);

//		sprintf(buf,"AT+RST\r\n");		//����WiFiģ��ΪSTAģʽ
//		HMISends_0(buf);						//���ô���1����ģ��
//		delay_ms(500);

//		//sprintf(buf,"AT+CWJAP=\"2BZNB\",\"zb970513\"\r\n");		//���ӵ�·�������ƺ�����
////		sprintf(buf,"AT+CWJAP=\"Python\",\"song7871855\"\r\n");
//		sprintf(buf,"AT+CWJAP=\"szy\",\"12345678\"\r\n");
//		HMISends_0(buf);
//		for(i=0; i<5; i++)
//		delay_ms(1000);
//		
//		sprintf(buf,"AT+CIPMUX=0\r\n");
//		HMISends_0(buf);
//		delay_ms(500);
		
		
//		sprintf(buf,"AT+CIPSTART=\"TCP\",\"132.232.126.221\",7777\r\n");		//Ŀ��������õ�IP��ַ���˿�
//		sprintf(buf, "AT+CIPSTART=\"UDP\",\"114.118.7.161\",123\r\n"); //������ʱ�������õ�IP��ַ���˿�
		sprintf(buf,"AT+CIPSTART=\"TCP\",\"47.102.102.16\",7777\r\n");
		HMISends_0(buf);
		for(i=0; i<1; i++)
		delay_ms(1000);
	
		sprintf(buf, "AT+CIPMODE=1\r\n"); //�ɷ��͵����ݵ�λ��
		HMISends_0(buf);					 //���ô���1����ģ��
		delay_ms(500);

		sprintf(buf, "AT+CIPSEND\r\n"); //�ɷ��͵����ݵ�λ��
		HMISends_0(buf);				   //���ô���1����ģ��
		delay_ms(500);
		
		uart_init(9600);
		
		g_StreamEndFlag = 0;
		g_RecvBufferSize = 0;
		SendConnect();
		delay_ms(500);
		HandleConncetResponse(g_RecvBuffer, g_RecvBufferSize);
		delay_ms(500);
		
		g_StreamEndFlag = 0;
		g_RecvBufferSize = 0;
		SendTimeRequest();
		delay_ms(500);
		HandleTimeResponse(g_RecvBuffer, g_RecvBufferSize);
		delay_ms(500);
		
		g_StreamEndFlag = 0;
		g_RecvBufferSize = 0;
		SendScheduleTimeRequest();
//		for(i=0; i<3; i++)
//		delay_ms(1000);
//		
//		
		for(i=0; i<3; i++)
		delay_ms(1000);
		g_StreamEndFlag = 0;
		g_RecvBufferSize = 0;
		SendScheduleTimeRequest();
		for(i=0; i<3; i++)
		delay_ms(1000);
		SendScheduleTimeRequest();
		g_StreamEndFlag = 0;
		g_RecvBufferSize = 0;
		
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
				HandleHeartResponse(g_RecvBuffer, g_RecvBufferSize);
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
			HMI_REC();
			Handle_Schudeule_Rec(g_RecvBuffer, g_RecvBufferSize);
			Handle_Regular_Belling(g_RecvBuffer, g_RecvBufferSize);
			delay_ms(200);


			// �������б�־
			g_StreamEndFlag = 0;
			g_RecvBufferSize = 0;
		}
		restart();
		belling();
		HMI_DISPALY();
//		printf("minute_rec:%d\r\n", SCH[1]);
//		printf("hour_rec:%d\r\n", 	SCH[0]);
//		printf("minute_rec1:%d\r\n", SCH[3]);
//		printf("hour_rec1:%d\r\n", 	SCH[2]);
//		printf("minute:%d\r\n", 		now[1]);
//		printf("hour:%d\r\n", 			now[2]);
		//printf("bell:%d\r\n",      bell_flag);
		delay_ms(200);
	}
}

void oled_time(unsigned char net)
{
	if (net == 0)
	{
		//		OLED_Clear();
		OLED_ShowCHinese(28, 0, 0);  //��
		OLED_ShowCHinese(70, 0, 1);  //��
		OLED_ShowCHinese(112, 0, 2); //��
		OLED_ShowCHinese(28, 4, 3);  //ʱ
		OLED_ShowCHinese(70, 4, 4);  //��
		OLED_ShowCHinese(112, 4, 5); //��

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
