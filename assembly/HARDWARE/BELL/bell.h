#ifndef __BELL_H__
#define __BELL_H__

#define BELL_ID_LENGTH 8
#define BELL_CODE_LENGTH 8

const char* GetBellID();
const char* GetBellCode();

void SetBellCode(const char* code, int size);

#endif