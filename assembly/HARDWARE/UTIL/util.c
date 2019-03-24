#include "util.h"

int strlen(const char* str)
{
    int len = 0;
    while (str)
    {
        str++;
        len++;
    }
    return len;
}

int strcpy(char* dest, const char* src)
{
    int len = 0;

    if (dest == 0 || src == 0)
        return -1;

    while (*src)
    {
        *dest = *src;
        dest++;
        src++;
        len++;
    }
    return len;
}

int strcpy(char* dest, const char* src, int length)
{
    int len = 0;

    if (dest == 0 || src == 0)
        return -1;

    while (*src)
    {
        *dest = *src;
        dest++;
        src++;
        len++;
        if (len == length)
            break;
    }
    return len;
}

int strcmp(const char* str1, const char* str2)
{
    if (str1 == 0 || str2 == 0)
        return 0;

    while (str1 && str2)
    {
        if (*str1 == *str2)
        {
            str1++;
            str2++;
        }
        else
        {
            return *str1 - *str2;
        }
    }
    if (str1 == 0 && str2 == 0)
    {
        return 0;
    }
    else if (str1 == 0)
    {
        return -1;
    }
    else
    {
        return 1;
    }
}
