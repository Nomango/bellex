#include "ds1302.h"
#include "delay.h"
#include "stdio.h"
#include "usart.h"
#include "string.h"
#include "stdlib.h"

unsigned char read[] = {0x81, 0x83, 0x85, 0x87, 0x89, 0x8b, 0x8d};  //读秒、分、时、日、月、周、年的寄存器地址
unsigned char write[] = {0x80, 0x82, 0x84, 0x86, 0x88, 0x8a, 0x8c}; //写秒、分、时、日、月、周、年的寄存器地址

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
	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);  //使能PB,PE端口时钟
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_7 | GPIO_Pin_8; //LED0-->PB.5 端口配置
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;	   //推挽输出
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;	  //IO口速度为50MHz
	GPIO_Init(GPIOB, &GPIO_InitStructure);				   //根据设定参数初始化GPIOB.5
	GPIO_SetBits(GPIOB, GPIO_Pin_7);					   //PB.5 输出高
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

void write_1302byte(uint8_t dat) //写一个字节的数据sck上升沿写数据
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

uint8_t read_1302(uint8_t add) //读数据
{
	uint8_t i = 0, dat1 = 0x00;
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(3);
	GPIO_SetBits(GPIOB, ds1302rst);
	delay_us(3);
	write_1302byte(add); //先写寄存器的地址
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

void write_1302(uint8_t add, uint8_t dat) //向指定寄存器写入一个字节的数据
{
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(3); //略微延时
	GPIO_SetBits(GPIOB, ds1302rst);
	delay_us(2); //时间大约2us
	write_1302byte(add);
	delay_us(2);
	write_1302byte(dat);
	delay_us(2);
	GPIO_ResetBits(GPIOB, ds1302rst);
	GPIO_ResetBits(GPIOB, ds1302clk);
	delay_us(2);
}

void ds1302_init(uint8_t *write, uint32_t *time) //初始化1302
{
	uint8_t i = 0, j = 0;

	write_1302(0x8e, 0x00); //去除写保护
	for (i = 0; i < 7; i++)
	{
		//转BCD码
		j = time[i] % 10;
		time[i] = (time[i] / 10) * 16 + j;
	}
	for (i = 0; i < 7; i++) //进行对时
		write_1302(write[i], time[i]);
	write_1302(0x8e, 0x80); //加写保护
}

void ds1302_data(uint8_t *read) //处理数据并通过串口打印
{
	uint8_t i = 0;
	uint8_t time[7];
	unsigned char ge[7], tim[7];

	for (i = 0; i < 7; i++)
		time[i] = read_1302(read[i]); //读数据已经完成

	for (i = 0; i < 7; i++)
	{
		ge[i] = time[i] % 16;  //此时已转换成10进制数，g[i]里面存放的是秒分时日月周年的各个位数据
		tim[i] = time[i] / 16; //而此时的time【i】里面存放的则是秒分时日月周年的十位数据
	}

	for (i = 0; i < 7; i++)
		now[i] = tim[i] * 10 + ge[i];

	//	if(s!=(time[0]+g[0]))
	//		printf("20%d%d年%d%d月%d%d日%d%d:%d%d:%d%d 星期%d\r\n",
	//		tim[6],ge[6],tim[4],ge[4],tim[3],ge[3],tim[2],ge[2],tim[1],ge[1],tim[0],ge[0],ge[5]);
	//	delay_ms(500);
}
