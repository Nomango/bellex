#include "delay.h"
#include "sys.h"
#include "ntp.h"
#include "string.h"




struct NtpPacket DefaultPacket()
{
	struct NtpPacket packet;
	memset(&packet, 0, sizeof(struct NtpPacket));
	*((char *)&packet + 0) = 0x1b;  // 00 011 011 (or 0x1B)
	return packet;
}


//struct NtpPacket packet = DefaultPacket();






