#include "bell.h"
#include "../UTIL/util.h"

char g_BellID[BELL_ID_LENGTH + 1] = "12345678";
char g_BellCode[BELL_CODE_LENGTH + 1] = "........";

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
