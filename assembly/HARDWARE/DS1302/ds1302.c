#include "ds1302.h"
#include "delay.h"
#include "stdio.h"
#include "usart.h"
#include "string.h"
#include "stdlib.h"

unsigned char read[] = {0x81, 0x83, 0x85, 0x87, 0x89, 0x8b, 0x8d};  //���롢�֡�ʱ���ա��¡��ܡ���ļĴ�����ַ
unsigned char write[] = {0x80, 0x82, 0x84, 0x86, 0x88, 0x8a, 0x8c}; //д�롢�֡�ʱ���ա��¡��ܡ���ļĴ�����ַ

unsigned char now[7];
unsigned char rec_schedule[80];
unsigned char schedule_hour[20];
unsigned char schedule_sec[20];
unsigned char schedule_A[85] = {255, 255, 80,

								'0', '8', '3', '0',
								'0', '9', '1', '5',
								'0', '9', '2', '0',
								'1', '0', '0', '5',

								'1', '0', '2', '5',
								'1', '1', '1', '0',
								'1', '1', '1', '5',
								'1', '2', '0', '0',

								'1', '4', '0', '0',
								'1', '4', '4', '5',
								'1', '4', '5', '0',
								'1', '5', '3', '5',

								'1', '5', '5', '5',
								'1', '6', '4', '0',
								'1', '6', '4', '5',
								'1', '7', '3', '0',

								'1', '9', '0', '0',
								'1', '9', '4', '5',
								'1', '9', '5', '0',
								'2', '0', '3', '0',

								0xff, 0xfe};

void bell_init(void)
{
	GPIO_InitTypeDef GPIO_InitStructure;
	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);  //ʹ��PB,PE�˿�ʱ��
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_7 | GPIO_Pin_8; //LED0-->PB.5 �˿�����
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;	   //�������
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;	  //IO���ٶ�Ϊ50MHz
	GPIO_Init(GPIOB, &GPIO_InitStructure);				   //�����趨������ʼ��GPIOB.5
	GPIO_SetBits(GPIOB, GPIO_Pin_7);					   //PB.5 �����
	GPIO_SetBits(GPIOB, GPIO_Pin_8);
}

void ds1302_GPIO_Configuration(void)
{
	GPIO_InitTypeDef GPIO_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_12 | GPIO_Pin_13 | GPIO_Pin_14;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_Init(GPIOB, &GPIO_InitStructure);
}

void Dio_In(void)
{
	GPIO_InitTypeDef GPIO_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_13;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IPU;
	GPIO_Init(GPIOB, &GPIO_InitStructure);
}

void DIO_Out(void)
{
	GPIO_InitTypeDef GPIO_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_13;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_Init(GPIOB, &GPIO_InitStructure);
}

void write_1302byte(uint8_t dat) //дһ���ֽڵ�����sck������д����
{
	uint8_t i = 0;
	GPIO_ResetBits(GPIOB, ds1302clk);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(2);
	for (i = 0; i < 8; i++)
	{
		if (dat & 0x01)
			GPIO_SetBits(GPIOB, ds1302dat);
		else
			GPIO_ResetBits(GPIOB, ds1302dat);
		GPIO_SetBits(GPIOB, ds1302clk);
		delay_us(2);
		GPIO_ResetBits(GPIOB, ds1302clk);
		delay_us(2);
		dat >>= 1;
		delay_us(9);
	}
	//	GPIO_ResetBits(GPIOB,ds1302rst);
}

uint8_t read_1302(uint8_t add) //������
{
	uint8_t i = 0, dat1 = 0x00;
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(3);
	GPIO_SetBits(GPIOB, ds1302rst);
	delay_us(3);
	write_1302byte(add); //��д�Ĵ����ĵ�ַ
	Dio_In();
	delay_us(3);
	for (i = 0; i < 8; i++)
	{
		dat1 = dat1 >> 1;
		if (GPIO_ReadInputDataBit(GPIOB, ds1302dat) == 1)
			dat1 = dat1 | 0x80;
		GPIO_SetBits(GPIOB, ds1302clk);
		delay_us(2);
		GPIO_ResetBits(GPIOB, ds1302clk);
		delay_us(2);
		delay_us(9);
	}
	delay_us(2);
	GPIO_ResetBits(GPIOB, ds1302rst);
	delay_us(4);
	DIO_Out();
	delay_us(2);
	return dat1;
}

void write_1302(uint8_t add, uint8_t dat) //��ָ���Ĵ���д��һ���ֽڵ�����
{
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(3); //��΢��ʱ
	GPIO_SetBits(GPIOB, ds1302rst);
	delay_us(2); //ʱ���Լ2us
	write_1302byte(add);
	delay_us(2);
	write_1302byte(dat);
	delay_us(2);
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(2);
}

void ds1302_init(uint8_t *write, uint32_t *time) //��ʼ��1302
{
	uint8_t i = 0, j = 0;

	write_1302(0x8e, 0x00); //ȥ��д����
	for (i = 0; i < 7; i++)
	{
		//תBCD��
		j = time[i] % 10;
		time[i] = (time[i] / 10) * 16 + j;
	}
	for (i = 0; i < 7; i++) //���ж�ʱ
		write_1302(write[i], time[i]);
	write_1302(0x8e, 0x80); //��д����
}

void ds1302_data(uint8_t *read) //�������ݲ�ͨ�����ڴ�ӡ
{
	uint8_t i = 0;
	uint8_t time[7];
	unsigned char ge[7], tim[7];

	for (i = 0; i < 7; i++)
		time[i] = read_1302(read[i]); //�������Ѿ����

	for (i = 0; i < 7; i++)
	{
		ge[i] = time[i] % 16;  //��ʱ��ת����10��������g[i]�����ŵ������ʱ��������ĸ���λ����
		tim[i] = time[i] / 16; //����ʱ��time��i�������ŵ��������ʱ���������ʮλ����
	}

	for (i = 0; i < 7; i++)
		now[i] = tim[i] * 10 + ge[i];

	//	if(s!=(time[0]+g[0]))
	//		printf("20%d%d��%d%d��%d%d��%d%d:%d%d:%d%d ����%d\r\n",
	//		tim[6],ge[6],tim[4],ge[4],tim[3],ge[3],tim[2],ge[2],tim[1],ge[1],tim[0],ge[0],ge[5]);
	//	delay_ms(500);
}
