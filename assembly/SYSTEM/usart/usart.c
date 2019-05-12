#include "sys.h"
#include "usart.h"	
#include "delay.h"	
#include "ntp.h"
#include "string.h"
//////////////////////////////////////////////////////////////////////////////////
//如果使用ucos,则包括下面的头文件即可.
#if SYSTEM_SUPPORT_OS
#include "includes.h" //ucos 使用
#endif

//extern u8 control_0;
//extern u16 A,B;
u8 HMI;
unsigned char g_RecvBuffer[261];
unsigned char g_RecvBufferSize=0;
int g_StreamEndFlag = 0;
//u8 rec_buf[];

//////////////////////////////////////////////////////////////////
//加入以下代码,支持printf函数,而不需要选择use MicroLIB
#if 1
#pragma import(__use_no_semihosting)
//标准库需要的支持函数
struct __FILE
{
	int handle;
};

FILE __stdout;       
//定义_sys_exit()以避免使用半主机模式    
_sys_exit(int x) 
{ 
	x = x; 
} 
//重定义fputc函数 
int fputc(int ch, FILE *f)
{      
	while((USART1->SR&0X40)==0);//循环发送,直到发送完毕   
    USART1->DR = (u8) ch;      
	return ch;
}
#endif

void uart2_init(u32 bound)
{
	//GPIO端口设置
	GPIO_InitTypeDef GPIO_InitStructure;
	USART_InitTypeDef USART_InitStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE); //使能USART1，GPIOA时钟
	RCC_APB1PeriphClockCmd(RCC_APB1Periph_USART2, ENABLE);

	//USART1_TX   GPIOA.9
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_2; //PA.2
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AF_PP; //复用推挽输出
	GPIO_Init(GPIOA, &GPIO_InitStructure);					//初始化GPIOA.2

	//USART1_RX	  GPIOA.10初始化
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_3;							//PA3
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN_FLOATING; //浮空输入
	GPIO_Init(GPIOA, &GPIO_InitStructure);								//初始化GPIOA.3

	//Usart1 NVIC 配置
	NVIC_InitStructure.NVIC_IRQChannel = USART2_IRQn;
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 3; //抢占优先级3
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 3;				//子优先级3
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE;						//IRQ通道使能
	NVIC_Init(&NVIC_InitStructure);														//根据指定的参数初始化VIC寄存器

	//USART 初始化设置
	USART_InitStructure.USART_BaudRate = bound;																			//串口波特率
	USART_InitStructure.USART_WordLength = USART_WordLength_8b;											//字长为8位数据格式
	USART_InitStructure.USART_StopBits = USART_StopBits_1;													//一个停止位
	USART_InitStructure.USART_Parity = USART_Parity_No;															//无奇偶校验位
	USART_InitStructure.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //无硬件数据流控制
	USART_InitStructure.USART_Mode = USART_Mode_Rx | USART_Mode_Tx;									//收发模式

	USART_Init(USART2, &USART_InitStructure);			 //初始化串口2
	USART_ITConfig(USART2, USART_IT_RXNE, ENABLE); //开启串口接受中断(一个字节)
	USART_Cmd(USART2, ENABLE);										 //使能串口2
	USART_ITConfig(USART2, USART_IT_IDLE, ENABLE); //开启串口接受中断(一帧数据)
}

void usart1_send(u8 data)
{
	USART1->DR = data;
	while ((USART1->SR & 0x40) == 0);
}
#if EN_USART1_RX //如果使能了接收

void uart_init(u32 bound)
{
	//GPIO端口设置
	GPIO_InitTypeDef GPIO_InitStructure;
	USART_InitTypeDef USART_InitStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_USART1 | RCC_APB2Periph_GPIOA, ENABLE); //使能USART1，GPIOA时钟

	//USART1_TX   GPIOA.9
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_9; //PA.9
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AF_PP; //复用推挽输出
	GPIO_Init(GPIOA, &GPIO_InitStructure);					//初始化GPIOA.9

	//USART1_RX	  GPIOA.10初始化
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_10;						//PA10
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN_FLOATING; //浮空输入
	GPIO_Init(GPIOA, &GPIO_InitStructure);								//初始化GPIOA.10

	//Usart1 NVIC 配置
	NVIC_InitStructure.NVIC_IRQChannel = USART1_IRQn;
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 2; //抢占优先级3
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 2;				//子优先级3
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE;						//IRQ通道使能
	NVIC_Init(&NVIC_InitStructure);														//根据指定的参数初始化VIC寄存器

	//USART 初始化设置
	USART_InitStructure.USART_BaudRate = bound;																			//串口波特率
	USART_InitStructure.USART_WordLength = USART_WordLength_8b;											//字长为8位数据格式
	USART_InitStructure.USART_StopBits = USART_StopBits_1;													//一个停止位
	USART_InitStructure.USART_Parity = USART_Parity_No;															//无奇偶校验位
	USART_InitStructure.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //无硬件数据流控制
	USART_InitStructure.USART_Mode = USART_Mode_Rx | USART_Mode_Tx;									//收发模式

	USART_Init(USART1, &USART_InitStructure);			 //初始化串口1
	USART_ITConfig(USART1, USART_IT_RXNE, ENABLE); //开启串口接受中断(一个字节)
	USART_Cmd(USART1, ENABLE);										 //使能串口1
	USART_ITConfig(USART1, USART_IT_IDLE, ENABLE); //开启串口接受中断(一帧数据)
}

void USART2_IRQHandler(void) //串口1中断服务程序
{
	if (USART_GetITStatus(USART2, USART_IT_RXNE) != RESET) //接受到一个字节
	{
		USART_ClearITPendingBit(USART2, USART_IT_RXNE); //清空标志位
		g_RecvBuffer[g_RecvBufferSize] = USART_ReceiveData(USART2);
		g_RecvBufferSize++;
	}

	if (USART_GetITStatus(USART2, USART_IT_IDLE) != RESET)
	{																									//接受到一帧数据
		USART_ClearITPendingBit(USART2, USART_IT_RXNE); //清空标志位
		USART2->DR;
		USART2->SR;

		g_StreamEndFlag = 1;  // 设置流结束标志位为 1
	}
}

void USART1_IRQHandler(void) //串口1中断服务程序
{
	if (USART_GetITStatus(USART1, USART_IT_RXNE) != RESET) //接受到一个字节
	{
		USART_ClearITPendingBit(USART1, USART_IT_RXNE); //清空标志位
		HMI = USART_ReceiveData(USART1);
		g_StreamEndFlag = 1;
	}


}

void HMISends_0(char *buf1)
{
	u8 i = 0;
	while (1)
	{
		if (buf1[i] != '\0')
		{
			USART_SendData(USART2, buf1[i]); //发送一个字节
			while (USART_GetFlagStatus(USART2, USART_FLAG_TXE) == RESET)
			{
			}; //等待发送结束
			i++;
		}
		else
		{
			//			USART_SendData(USART2,'\0');
			return;
		}
	}
}

void HMISends_HMI(unsigned char *buf1)		  //字符串发送函数
{
		u8 i=0;
	while(1)
	{
	 if(buf1[i]!=0)
	 	{
			USART_SendData(USART1,buf1[i]);  //发送一个字节
			while(USART_GetFlagStatus(USART1,USART_FLAG_TXE)==RESET){};//等待发送结束
		 	i++;
		}
	 else 
	 return ;

		}
}

void HMISendb(u8 k)		         //字节发送函数
{		 
	u8 i;
	 for(i=0;i<3;i++)
	 {
	 if(k!=0)
	 	{
			USART_SendData(USART1,k);  //发送一个字节
			while(USART_GetFlagStatus(USART1,USART_FLAG_TXE)==RESET){};//等待发送结束
		}
	 else 
	 return ;

	 } 
} 

void UART2_Send_Array(unsigned char send_array[], unsigned char num)
{
	unsigned char i = 0;
	for (i = 0; i < num; i++)
	{
		USART_SendData(USART2, send_array[i]);
		while (USART_GetFlagStatus(USART2, USART_FLAG_TC) != SET)
			;
	}
}

#endif
