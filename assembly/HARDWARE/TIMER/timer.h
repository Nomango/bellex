#ifndef __TIMER_H
#define __TIMER_H

#include "sys.h"
#include "stm32f10x_tim.h"
#include "usart.h"
#include "oled.h"

void TIM3_Int_Init(u16 arr,u16 psc);
void TIM4_Int_Init(u16 arr,u16 psc);
void TIM6_Int_Init(u16 arr,u16 psc);
void belling(void);
extern unsigned char send_flag;
extern unsigned char oled_flag;

#endif
