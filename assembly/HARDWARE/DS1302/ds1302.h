#ifndef _DS1302_H
#define _DS1302_H

#include "stm32f10x.h"

#define ds1302clk GPIO_Pin_14
#define ds1302dat GPIO_Pin_13
#define ds1302rst GPIO_Pin_12    

void Dio_In(void);
void DIO_Out(void);
void ds1302_GPIO_Configuration(void);
void write_1302byte(uint8_t dat);//дһ���ֽڵ�����sck������д����
uint8_t read_1302(uint8_t add);//������
void write_1302(uint8_t add,uint8_t dat);//��ָ���Ĵ���д��һ���ֽڵ�����
void ds1302_init(uint8_t *write,uint32_t *time);//��ʼ��1302
void ds1302_data(uint8_t *read);//�������ݲ�ͨ�����ڴ�ӡ
void bell_init(void);

extern unsigned char read[7],write[7];
extern unsigned char now[7];
extern unsigned char schedule_A[85];
extern unsigned char rec_schedule[80];
extern unsigned char schedule_hour[20];
extern unsigned char schedule_sec[20];

#define bell_1 PBout(7)// PB5
#define bell_2 PBout(8)// PE5	

#endif

