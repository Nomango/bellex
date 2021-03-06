#ifndef __USART_H
#define __USART_H
#include "stdio.h"	
#include "sys.h" 

#define USART_REC_LEN  			200  	//定义最大接收字节数 200
#define EN_USART1_RX 			1		//使能（1）/禁止（0）串口1接收
extern unsigned char g_RecvBuffer[261],REC[100];
extern unsigned char g_RecvBufferSize;
extern int g_StreamEndFlag;
//extern u8 rec_buf[100];
extern u8  USART_RX_BUF[USART_REC_LEN]; //接收缓冲,最大USART_REC_LEN个字节.末字节为换行符 
extern u16 USART_RX_STA;         		//接收状态标记	
//如果想串口中断接收，请不要注释以下宏定义
void uart_init(u32 bound);
void usart1_send(u8 data);
void uart2_init(u32 bound);
void HMISends();
void HMISendb(u8 k);
void HMISends_0();
void HMISenb();
void OK_FLAG(void);
void HMISendstart(void);
void UART1_Send_Array(unsigned char send_array[],unsigned char num);
void UART2_Send_Array(unsigned char send_array[],unsigned char num);
void HMISends_HMI(unsigned char *buf1);
#endif


