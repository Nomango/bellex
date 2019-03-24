#include "key_send.h"


void wifi_connect(void){
	
	unsigned char AT_buf[64];	
	unsigned char key;
	struct NtpPacket packet = DefaultPacket();
	unsigned char request_time[39]={255,255,34,'i','d',':','1','2','3',';','c','o','d','e',':','1',
																	'2','3',';','r','e','q',':','r','e','q','u','e','s','t','_','t',
																	'i','m','e',';','#',0xff,0xfe};
	unsigned char request_schedule[34]={255,255,29,'i','d',':','1','2','3',';','c','o','d','e',':','1',
																		'2','3',';','r','e','q',':','s','c','h','e','d','u','l','e',';',0xff,0xfe};
	
	unsigned char request_code[30]={255,255,25,'i','d',':','1','2','3',';','c','o','d','e',':',' ',
																	'r','e','q',':','c','o','n','n','e','c','t',';',0xff,0xfe};

	unsigned char heart_beat[36]={255,255,31,'i','d',':','1','2','3',';','c','o','d','e',':','1','2','3',
																	'r','e','q',':','h','e','a','r','t','_','b','e','a','t',';','#',0xff,0xfe};																		
	
											
	key=KEY_Scan(0);
	
	if(key==KEY2_PRES){
//		sprintf(AT_buf,"AT+CIPSTART=\"UDP\",\"114.118.7.161\",123\r\n");		//国家授时服务器得到IP地址、端口
//		sprintf(AT_buf,"AT+CIPSTART=\"TCP\",\"192.168.1.101\",7777\r\n");		//虚拟服务器得到IP地址、端口
		sprintf(AT_buf,"AT+CIPSTART=\"TCP\",\"132.232.126.221\",7777\r\n");		//目标服务器得到IP地址、端口
		HMISends_0(AT_buf);
		delay_ms(500);
	}

	if(key==KEY3_PRES){	 
		sprintf(AT_buf,"AT+CIPMODE=1\r\n");		//可发送的数据的位数
		HMISends_0(AT_buf);										//利用串口1配置模块
		delay_ms(500);

		sprintf(AT_buf,"AT+CIPSEND\r\n");		//可发送的数据的位数
		HMISends_0(AT_buf);										//利用串口1配置模块
		delay_ms(500);

	}
	
	if(key==KEY4_PRES){
		UART2_Send_Array(heart_beat,36);
//		UART2_Send_Array(request_code,30);
		delay_ms(500);
//		UART2_Send_Array(request_time,39);
//		UART2_Send_Array(request_schedule,34);
//		UART2_Send_Array((unsigned char*)&packet,sizeof(struct NtpPacket));
		delay_ms(100);
	}
}
	
