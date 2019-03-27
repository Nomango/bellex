#include "key_send.h"
#include "../BELL/request.h"

unsigned char AT_buf[64];
unsigned char key;

void wifi_connect(void)
{
	key = KEY_Scan(0);

	if (key == KEY2_PRES)
	{
//		sprintf(AT_buf, "AT+CIPSTART=\"UDP\",\"114.118.7.161\",123\r\n"); //������ʱ�������õ�IP��ַ���˿�
//		sprintf(AT_buf,"AT+CIPSTART=\"TCP\",\"192.168.1.101\",7777\r\n");		//����������õ�IP��ַ���˿�
		sprintf(AT_buf,"AT+CIPSTART=\"TCP\",\"132.232.126.221\",7777\r\n");		//Ŀ��������õ�IP��ַ���˿�
		HMISends_0(AT_buf);
		delay_ms(500);
		
		sprintf(AT_buf, "AT+CIPMODE=1\r\n"); //�ɷ��͵����ݵ�λ��
		HMISends_0(AT_buf);					 //���ô���1����ģ��
		delay_ms(500);

		sprintf(AT_buf, "AT+CIPSEND\r\n"); //�ɷ��͵����ݵ�λ��
		HMISends_0(AT_buf);				   //���ô���1����ģ��
		delay_ms(500);
	}

	if (key == KEY3_PRES)
	{
		delay_ms(100);
		SendConnect();
		delay_ms(100);
	}

	if (key == KEY4_PRES)
	{
		delay_ms(100);
//		SendConnect();
//		SendHeartBeat();
		SendTimeRequest();
//		SendNTPRequest();

		delay_ms(100);
	}
}
