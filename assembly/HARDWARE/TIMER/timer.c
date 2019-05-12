#include "timer.h"
#include "key_send.h"
#include "../BELL/request.h"

unsigned char send_flag=0;
unsigned char oled_flag=0;
extern u16 bell_flag;
u16 flag=0;
void TIM3_Int_Init(u16 arr,u16 psc)
{
  TIM_TimeBaseInitTypeDef  TIM_TimeBaseStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB1PeriphClockCmd(RCC_APB1Periph_TIM3, ENABLE); //时钟使能
	
	//定时器TIM3初始化
	TIM_TimeBaseStructure.TIM_Period = arr; //设置在下一个更新事件装入活动的自动重装载寄存器周期的值	
	TIM_TimeBaseStructure.TIM_Prescaler =psc; //设置用来作为TIMx时钟频率除数的预分频值
	TIM_TimeBaseStructure.TIM_ClockDivision = TIM_CKD_DIV1; //设置时钟分割:TDTS = Tck_tim
	TIM_TimeBaseStructure.TIM_CounterMode = TIM_CounterMode_Up;  //TIM向上计数模式
	TIM_TimeBaseInit(TIM3, &TIM_TimeBaseStructure); //根据指定的参数初始化TIMx的时间基数单位
 
	TIM_ITConfig(TIM3,TIM_IT_Update,ENABLE ); //使能指定的TIM3中断,允许更新中断

	//中断优先级NVIC设置
	NVIC_InitStructure.NVIC_IRQChannel = TIM3_IRQn;  //TIM3中断
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 0;  //先占优先级0级
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 3;  //从优先级3级
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE; //IRQ通道被使能
	NVIC_Init(&NVIC_InitStructure);  //初始化NVIC寄存器


	TIM_Cmd(TIM3, ENABLE);  //使能TIMx					 
}

void TIM4_Int_Init(u16 arr,u16 psc)
{
  TIM_TimeBaseInitTypeDef  TIM_TimeBaseStructure;
	NVIC_InitTypeDef NVIC_InitStructure;

	RCC_APB1PeriphClockCmd(RCC_APB1Periph_TIM4, ENABLE); //时钟使能
	
	//定时器TIM3初始化
	TIM_TimeBaseStructure.TIM_Period = arr; //设置在下一个更新事件装入活动的自动重装载寄存器周期的值	
	TIM_TimeBaseStructure.TIM_Prescaler =psc; //设置用来作为TIMx时钟频率除数的预分频值
	TIM_TimeBaseStructure.TIM_ClockDivision = TIM_CKD_DIV1; //设置时钟分割:TDTS = Tck_tim
	TIM_TimeBaseStructure.TIM_CounterMode = TIM_CounterMode_Up;  //TIM向上计数模式
	TIM_TimeBaseInit(TIM4, &TIM_TimeBaseStructure); //根据指定的参数初始化TIMx的时间基数单位
 
	TIM_ITConfig(TIM4,TIM_IT_Update,ENABLE ); //使能指定的TIM3中断,允许更新中断

	//中断优先级NVIC设置
	NVIC_InitStructure.NVIC_IRQChannel = TIM4_IRQn;  //TIM3中断
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 0;  //先占优先级0级
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 3;  //从优先级3级
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE; //IRQ通道被使能
	NVIC_Init(&NVIC_InitStructure);  //初始化NVIC寄存器


	TIM_Cmd(TIM4, ENABLE);  //使能TIMx					 
}




//定时器3中断服务程序
void TIM4_IRQHandler(void)   //TIM3中断
{
	static int i = 1;
	static int j;
	if (TIM_GetITStatus(TIM4, TIM_IT_Update) != RESET)  //检查TIM3更新中断发生与否
	{
		i++;
		if ( i%15 == 0)
		{
			/* 心跳包请求 */
			SendHeartBeat();
		}
		
		if (j == 20)
		{
			//printf("ok\r\n");
			bell_flag = 0;
		}
		
		if (j == 60)
		{
			//printf("flag\r\n");
			flag = 0;
			j = 0;
		}
		
		if ( bell_flag == 1)
		{
			//printf("j:%d\r\n",j);
			flag = 1;
		}
		
		if ( flag == 1)
		{
			j++;
		}
		
		if ( i%1200 == 0)
		{
			i = 1;
			/* 获取时间请求 */
			SendTimeRequest();
		}
		
		TIM_ClearITPendingBit(TIM4, TIM_IT_Update  );  //清除TIMx更新中断标志 
	}
}


//定时器3中断服务程序
void TIM3_IRQHandler(void)   //TIM3中断
{
	static unsigned char i=0,j=0,k=0;
	if (TIM_GetITStatus(TIM3, TIM_IT_Update) != RESET)  //检查TIM3更新中断发生与否
	{
		TIM_ClearITPendingBit(TIM3, TIM_IT_Update  );  //清除TIMx更新中断标志 
		
		if(k > (10-1)){
//			OLED_Clear();
			if( oled_flag % 2 == 1){
				oled_flag =0;
				k = 0;
			}
			else{
				oled_flag =1;
				k = 0;
			}
		}
		else
			k++;
		
		
		if( i > (120-1) ){  //60s
			j++;
			i=0;
		}
		else
			i++;
		
		if( j > (5-1) ){     //5min
			send_flag = 1;
			j = 0;
		}
		else 
			send_flag = 0;
		
	}
}


void belling(void)
{
	if (flag != 1)
	{
		//printf("flag\r\n");
		bell_flag = bell_on(g_RecvBuffer, g_RecvBufferSize);
	}
	
	if (bell_flag == 1)
	{
		GPIO_SetBits(GPIOB,GPIO_Pin_8);
	}
	
	else if (bell_flag == 0)
	{
		GPIO_ResetBits(GPIOB,GPIO_Pin_8);
	}
}









