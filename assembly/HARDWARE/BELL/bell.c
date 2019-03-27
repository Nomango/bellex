#include "bell.h"
#include "string.h"
#include "ds1302.h"


char g_BellID[BELL_ID_LENGTH + 1] = "01234546";
char g_BellCode[BELL_CODE_LENGTH + 1] = "00000000";

const char* GetBellID()
{
    return g_BellID;
}

const char* GetBellCode()
{
    return g_BellCode;
}

void SetBellCode(const char* code, int size)
{
    if (size != BELL_CODE_LENGTH)
        return;

    strncpy(g_BellCode, code, size);
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
    if (size != (12 + 8 + 1))
        return;

    for (i = 0; i < 8; i++)
    {
        code[i] = recv[12 + i];
    }
    code[8] = '\0';

    SetBellCode(code, 8);
}

void HandleTimeResponse(unsigned char* recv, unsigned char size)
{
    int i;
    unsigned int now_time[7];

    if (strncmp("current_time:", recv, 13) != 0)
    {
        return;
    }

    
    if (size != (13 + 14 + 1))
        return;

    for (i = 0; i < 7; i++)
    {
        now_time[i] = (int)((recv[13 + i * 2] - '0') * 10 + (recv[13 + i * 2 + 1] - '0'));
    }

    ds1302_init(write, now_time);
}


