#include "bell.h"
#include "string.h"
#include "ds1302.h"
#include "key.h"
#include "key_send.h"
#include "../BELL/request.h"
#include "usart.h"	
unsigned char SCH[100];
unsigned char sch[100];
char g_BellID[BELL_ID_LENGTH + 1] = "87654321";
char g_BellCode[BELL_CODE_LENGTH + 1] = "00000000";
extern u16 bell_flag;
int len;
const char* GetBellID(void)
{
    return g_BellID;
}

const char* GetBellCode(void)
{
    return g_BellCode;
}

void SetBellCode(const char* code, int size)
{
    if (size != BELL_CODE_LENGTH)
        return;

    strncpy(g_BellCode, code, size);
}

char* GetBellStatus(void)
{
//		如果继电器通，则返回working，否则为idle
	if (bell_flag == 0)
	{
		return "idle";
	}
	
	else if (bell_flag == 1)
	{
		return "working";
	}
}

void HandleConncetResponse(unsigned char* recv, unsigned char size)
{
    int i;
    char code[9];

    if (strncmp("unique_code:", recv, 12) != 0)
    {
        return;
    }

    // 处理这个请求必须满足21个字符
    if (size != (12 + 8 + 2))
        return;

    for (i = 0; i < 8; i++)
    {
        code[i] = recv[12 + i];
    }
    code[8] = '\0';

    SetBellCode(code, 8);
}

void HandleHeartResponse(unsigned char* recv, unsigned char size)
{
    int i;
		unsigned char buf[100];
    unsigned char Beat = 0;

    if (strncmp("status:", recv, 7) != 0)
    {
        return;
    }

    
    if (size != (7 + 1 + 1))
        return;
		
		Beat = recv[7];
		
		if(Beat == '1'){
				return;
//				LED = ~LED;
		}
		else{
			LED = 0;
			

		}
		
}

void HandleTimeResponse(unsigned char* recv, unsigned char size)
{
    int i;
    unsigned int server_now_time[7];

    if (strncmp("current_time:", recv, 13) != 0)
    {
        return;
    }

    
    if (size != (13 + 14 + 2))
        return;

    for (i = 0; i < 7; i++)
    {
        server_now_time[i] = (int)((recv[13 + i * 2] - '0') * 10 + (recv[13 + i * 2 + 1] - '0'));
    }

    ds1302_init(write, server_now_time);
}

void Handle_Schudeule_Rec(unsigned char* recv, unsigned char size)
{
    int i, j=0, t=0;

    if (strncmp("schedule:", recv, 9) != 0)
    {
        return;
    }
		
		i = strlen(recv);
		for (j = 0; j < i; j++)
    {
				SCH[j] = '\0';		
    }
		
		i = i - 10;
		j = 0;
		
    for (j = 0, t = 0; j < i; j+=2, t++)
    {
				if ( recv[9+j] != ';')
				{
					SCH[t] = (recv[9+j] - 48)*10 + (recv[10+j] - 48);		
				}
    }
		
		len = strlen(SCH);
		t = 0;
}



   /* 定时打铃 */
void Handle_Regular_Belling(unsigned char* recv, unsigned char size)
{
		if ((strncmp("bell:current;", recv, 13) != 0) && (strncmp("bell", recv, 4) == 0))
		{
			sch[0] = (recv[5] - 48)*10 + recv[6] - 48;
			sch[1] = (recv[7] - 48)*10 + recv[8] - 48;
			
			SCH[len] = sch[0];
			SCH[len+1] = sch[1];
			
		}
		
		else
		{
			return;
		}


}

int bell_on(unsigned char* recv, unsigned char size)
{
    int i, j=0;
	
		i = strlen(SCH);
	
    for (j = 0; j < i; j+=2)
    {
        if( (now[2] == SCH[j]) && (now[1] == SCH[j+1]))
				{
//					printf("SCH\r\n");
					return 1;
				}
    }
		
	

		if((now[2] == sch[0]) && (now[1] == sch[1]))
			return 1;

		if (strncmp("bell:current;", recv, 13) == 0)
    {
//			 printf("current\r\n");
       return 1;
    }
		

		return 0;
}

void restart(void)
{
	int i, j;
	if ((now[2] == 0) && (now[1] == 0))
	{
		i = strlen(SCH);
		for (j = 0; j < i; j++)
    {
				SCH[j] = '\0';		
    }
		
		SendScheduleTimeRequest();
		delay_ms(500);
		SendScheduleTimeRequest();
		delay_ms(500);
	}
}
