#include "sys.h"
#include "usart.h"	
#include "delay.h"	
#include "ntp.h"
#include "string.h"
//////////////////////////////////////////////////////////////////////////////////
//���ʹ��ucos,����������ͷ�ļ�����.
#if SYSTEM_SUPPORT_OS
#include "includes.h" //ucos ʹ��
#endif

//extern u8 control_0;
//extern u16 A,B;
u8 HMI;
unsigned char g_RecvBuffer[261];
unsigned char g_RecvBufferSize=0;
int g_StreamEndFlag = 0;
//u8 rec_buf[];

//////////////////////////////////////////////////////////////////
//�������´���,֧��printf����,������Ҫѡ��use MicroLIB
#if 1
#pragma import(__use_no_semihosting)
//��׼����Ҫ��֧�ֺ���
struct __FILE
{
	int handle;
};

FILE __stdout;       
//����_sys_exit()�Ա���ʹ�ð�����ģʽ    
_sys_exit(int x) 
{ 
	x = x; 
} 
//�ض���fputc���� 
int fputc(int ch, FILE *f)
{      
	while((USART1->SR&0X40)==0);//ѭ������,ֱ���������   
    USART1->DR = (u8) ch;      
	return ch;
}
#endif

void uart2_init(u32 bound)
{
	//GPIO�˿�����
	GPIO_InitTypeDef GPIO_InitStructure;
	USART_InitTypeDef USART_InitStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE); //ʹ��USART1��GPIOAʱ��
	RCC_APB1PeriphClockCmd(RCC_APB1Periph_USART2, ENABLE);

	//USART1_TX   GPIOA.9
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_2; //PA.2
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AF_PP; //�����������
	GPIO_Init(GPIOA, &GPIO_InitStructure);					//��ʼ��GPIOA.2

	//USART1_RX	  GPIOA.10��ʼ��
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_3;							//PA3
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN_FLOATING; //��������
	GPIO_Init(GPIOA, &GPIO_InitStructure);								//��ʼ��GPIOA.3

	//Usart1 NVIC ����
	NVIC_InitStructure.NVIC_IRQChannel = USART2_IRQn;
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 3; //��ռ���ȼ�3
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 3;				//�����ȼ�3
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE;						//IRQͨ��ʹ��
	NVIC_Init(&NVIC_InitStructure);														//����ָ���Ĳ�����ʼ��VIC�Ĵ���

	//USART ��ʼ������
	USART_InitStructure.USART_BaudRate = bound;																			//���ڲ�����
	USART_InitStructure.USART_WordLength = USART_WordLength_8b;											//�ֳ�Ϊ8λ���ݸ�ʽ
	USART_InitStructure.USART_StopBits = USART_StopBits_1;													//һ��ֹͣλ
	USART_InitStructure.USART_Parity = USART_Parity_No;															//����żУ��λ
	USART_InitStructure.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //��Ӳ������������
	USART_InitStructure.USART_Mode = USART_Mode_Rx | USART_Mode_Tx;									//�շ�ģʽ

	USART_Init(USART2, &USART_InitStructure);			 //��ʼ������2
	USART_ITConfig(USART2, USART_IT_RXNE, ENABLE); //�������ڽ����ж�(һ���ֽ�)
	USART_Cmd(USART2, ENABLE);										 //ʹ�ܴ���2
	USART_ITConfig(USART2, USART_IT_IDLE, ENABLE); //�������ڽ����ж�(һ֡����)
}

void usart1_send(u8 data)
{
	USART1->DR = data;
	while ((USART1->SR & 0x40) == 0);
}
#if EN_USART1_RX //���ʹ���˽���

void uart_init(u32 bound)
{
	//GPIO�˿�����
	GPIO_InitTypeDef GPIO_InitStructure;
	USART_InitTypeDef USART_InitStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_USART1 | RCC_APB2Periph_GPIOA, ENABLE); //ʹ��USART1��GPIOAʱ��

	//USART1_TX   GPIOA.9
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_9; //PA.9
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AF_PP; //�����������
	GPIO_Init(GPIOA, &GPIO_InitStructure);					//��ʼ��GPIOA.9

	//USART1_RX	  GPIOA.10��ʼ��
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_10;						//PA10
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN_FLOATING; //��������
	GPIO_Init(GPIOA, &GPIO_InitStructure);								//��ʼ��GPIOA.10

	//Usart1 NVIC ����
	NVIC_InitStructure.NVIC_IRQChannel = USART1_IRQn;
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 2; //��ռ���ȼ�3
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 2;				//�����ȼ�3
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE;						//IRQͨ��ʹ��
	NVIC_Init(&NVIC_InitStructure);														//����ָ���Ĳ�����ʼ��VIC�Ĵ���

	//USART ��ʼ������
	USART_InitStructure.USART_BaudRate = bound;																			//���ڲ�����
	USART_InitStructure.USART_WordLength = USART_WordLength_8b;											//�ֳ�Ϊ8λ���ݸ�ʽ
	USART_InitStructure.USART_StopBits = USART_StopBits_1;													//һ��ֹͣλ
	USART_InitStructure.USART_Parity = USART_Parity_No;															//����żУ��λ
	USART_InitStructure.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //��Ӳ������������
	USART_InitStructure.USART_Mode = USART_Mode_Rx | USART_Mode_Tx;									//�շ�ģʽ

	USART_Init(USART1, &USART_InitStructure);			 //��ʼ������1
	USART_ITConfig(USART1, USART_IT_RXNE, ENABLE); //�������ڽ����ж�(һ���ֽ�)
	USART_Cmd(USART1, ENABLE);										 //ʹ�ܴ���1
	USART_ITConfig(USART1, USART_IT_IDLE, ENABLE); //�������ڽ����ж�(һ֡����)
}

void USART2_IRQHandler(void) //����1�жϷ������
{
	if (USART_GetITStatus(USART2, USART_IT_RXNE) != RESET) //���ܵ�һ���ֽ�
	{
		USART_ClearITPendingBit(USART2, USART_IT_RXNE); //��ձ�־λ
		g_RecvBuffer[g_RecvBufferSize] = USART_ReceiveData(USART2);
		g_RecvBufferSize++;
	}

	if (USART_GetITStatus(USART2, USART_IT_IDLE) != RESET)
	{																									//���ܵ�һ֡����
		USART_ClearITPendingBit(USART2, USART_IT_RXNE); //��ձ�־λ
		USART2->DR;
		USART2->SR;

		g_StreamEndFlag = 1;  // ������������־λΪ 1
	}
}

void USART1_IRQHandler(void) //����1�жϷ������
{
	if (USART_GetITStatus(USART1, USART_IT_RXNE) != RESET) //���ܵ�һ���ֽ�
	{
		USART_ClearITPendingBit(USART1, USART_IT_RXNE); //��ձ�־λ
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
			USART_SendData(USART2, buf1[i]); //����һ���ֽ�
			while (USART_GetFlagStatus(USART2, USART_FLAG_TXE) == RESET)
			{
			}; //�ȴ����ͽ���
			i++;
		}
		else
		{
			//			USART_SendData(USART2,'\0');
			return;
		}
	}
}

void HMISends_HMI(unsigned char *buf1)		  //�ַ������ͺ���
{
		u8 i=0;
	while(1)
	{
	 if(buf1[i]!=0)
	 	{
			USART_SendData(USART1,buf1[i]);  //����һ���ֽ�
			while(USART_GetFlagStatus(USART1,USART_FLAG_TXE)==RESET){};//�ȴ����ͽ���
		 	i++;
		}
	 else 
	 return ;

		}
}

void HMISendb(u8 k)		         //�ֽڷ��ͺ���
{		 
	u8 i;
	 for(i=0;i<3;i++)
	 {
	 if(k!=0)
	 	{
			USART_SendData(USART1,k);  //����һ���ֽ�
			while(USART_GetFlagStatus(USART1,USART_FLAG_TXE)==RESET){};//�ȴ����ͽ���
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
