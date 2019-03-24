#ifndef _NTP_H
#define _NTP_H

#include <stdint.h>

struct NtpPacket 
{

  uint8_t li_vn_mode;      // Eight bits. li, vn, and mode.
                           // li.   Two bits.   Leap indicator.
                           // vn.   Three bits. Version number of the protocol.
                           // mode. Three bits. Client will pick mode 3 for client.

  uint8_t stratum;         // Eight bits. Stratum level of the local clock.
  uint8_t poll;            // Eight bits. Maximum interval between successive messages.
  uint8_t precision;       // Eight bits. Precision of the local clock.

  uint32_t root_delay;      // 32 bits. Total round trip delay time.
  uint32_t root_dispersion; // 32 bits. Max error aloud from primary clock source.
  uint32_t ref_id;          // 32 bits. Reference clock identifier.

  uint32_t ref_tm_s;        // 32 bits. Reference time-stamp seconds.
  uint32_t ref_tm_f;        // 32 bits. Reference time-stamp fraction of a second.

  uint32_t orig_tm_s;       // 32 bits. Originate time-stamp seconds.
  uint32_t orig_tm_f;       // 32 bits. Originate time-stamp fraction of a second.

  uint32_t rx_tm_s;         // 32 bits. Received time-stamp seconds.
  uint32_t rx_tm_f;         // 32 bits. Received time-stamp fraction of a second.

  uint32_t tx_tm_s;         // 32 bits and the most important field the client cares about. Transmit time-stamp seconds.
  uint32_t tx_tm_f;         // 32 bits. Transmit time-stamp fraction of a second.

};

struct NtpPacket DefaultPacket();

#endif
