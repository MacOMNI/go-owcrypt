/*
 * Copyright 2020 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */
package blake2s

func blocks(d *digest, p []uint8) {
	h0, h1, h2, h3, h4, h5, h6, h7 := d.h[0], d.h[1], d.h[2], d.h[3], d.h[4], d.h[5], d.h[6], d.h[7]

	for len(p) >= BlockSize {
		d.t[0] += BlockSize
		if d.t[0] < BlockSize {
			d.t[1]++
		}
		v0, v1, v2, v3, v4, v5, v6, v7 := h0, h1, h2, h3, h4, h5, h6, h7
		v8 := iv[0]
		v9 := iv[1]
		v10 := iv[2]
		v11 := iv[3]
		v12 := iv[4] ^ d.t[0]
		v13 := iv[5] ^ d.t[1]
		v14 := iv[6] ^ d.f[0]
		v15 := iv[7] ^ d.f[1]

		var m [16]uint32

		j := 0
		for i := 0; i < 16; i++ {
			m[i] = uint32(p[j]) | uint32(p[j+1])<<8 | uint32(p[j+2])<<16 | uint32(p[j+3])<<24
			j += 4
		}

		v0 += m[0]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[2]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[4]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[6]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[5]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[7]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[3]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[1]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[8]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[10]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[12]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[14]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[13]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[15]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[11]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[9]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		v0 += m[14]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[4]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[9]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[13]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[15]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[6]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[8]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[10]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[1]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[0]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[11]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[5]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[7]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[3]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[2]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[12]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 3.
		v0 += m[11]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[12]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[5]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[15]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[2]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[13]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[0]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[8]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[10]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[3]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[7]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[9]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[1]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[4]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[6]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[14]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 4.
		v0 += m[7]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[3]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[13]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[11]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[12]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[14]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[1]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[9]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[2]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[5]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[4]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[15]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[0]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[8]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[10]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[6]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 5.
		v0 += m[9]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[5]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[2]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[10]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[4]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[15]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[7]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[0]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[14]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[11]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[6]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[3]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[8]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[13]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[12]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[1]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 6.
		v0 += m[2]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[6]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[0]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[8]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[11]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[3]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[10]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[12]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[4]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[7]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[15]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[1]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[14]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[9]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[5]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[13]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 7.
		v0 += m[12]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[1]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[14]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[4]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[13]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[10]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[15]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[5]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[0]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[6]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[9]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[8]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[2]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[11]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[3]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[7]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 8.
		v0 += m[13]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[7]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[12]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[3]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[1]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[9]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[14]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[11]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[5]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[15]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[8]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[2]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[6]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[10]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[4]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[0]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 9.
		v0 += m[6]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[14]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[11]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[0]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[3]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[8]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[9]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[15]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[12]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[13]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[1]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[10]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[4]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[5]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[7]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[2]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		// Round 10.
		v0 += m[10]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-12) | v4>>12
		v1 += m[8]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-12) | v5>>12
		v2 += m[7]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-12) | v6>>12
		v3 += m[1]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-12) | v7>>12
		v2 += m[6]
		v2 += v6
		v14 ^= v2
		v14 = v14<<(32-8) | v14>>8
		v10 += v14
		v6 ^= v10
		v6 = v6<<(32-7) | v6>>7
		v3 += m[5]
		v3 += v7
		v15 ^= v3
		v15 = v15<<(32-8) | v15>>8
		v11 += v15
		v7 ^= v11
		v7 = v7<<(32-7) | v7>>7
		v1 += m[4]
		v1 += v5
		v13 ^= v1
		v13 = v13<<(32-8) | v13>>8
		v9 += v13
		v5 ^= v9
		v5 = v5<<(32-7) | v5>>7
		v0 += m[2]
		v0 += v4
		v12 ^= v0
		v12 = v12<<(32-8) | v12>>8
		v8 += v12
		v4 ^= v8
		v4 = v4<<(32-7) | v4>>7
		v0 += m[15]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-12) | v5>>12
		v1 += m[9]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-12) | v6>>12
		v2 += m[3]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-12) | v7>>12
		v3 += m[13]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-12) | v4>>12
		v2 += m[12]
		v2 += v7
		v13 ^= v2
		v13 = v13<<(32-8) | v13>>8
		v8 += v13
		v7 ^= v8
		v7 = v7<<(32-7) | v7>>7
		v3 += m[0]
		v3 += v4
		v14 ^= v3
		v14 = v14<<(32-8) | v14>>8
		v9 += v14
		v4 ^= v9
		v4 = v4<<(32-7) | v4>>7
		v1 += m[14]
		v1 += v6
		v12 ^= v1
		v12 = v12<<(32-8) | v12>>8
		v11 += v12
		v6 ^= v11
		v6 = v6<<(32-7) | v6>>7
		v0 += m[11]
		v0 += v5
		v15 ^= v0
		v15 = v15<<(32-8) | v15>>8
		v10 += v15
		v5 ^= v10
		v5 = v5<<(32-7) | v5>>7

		h0 ^= v0 ^ v8
		h1 ^= v1 ^ v9
		h2 ^= v2 ^ v10
		h3 ^= v3 ^ v11
		h4 ^= v4 ^ v12
		h5 ^= v5 ^ v13
		h6 ^= v6 ^ v14
		h7 ^= v7 ^ v15

		p = p[BlockSize:]
	}
	d.h[0], d.h[1], d.h[2], d.h[3], d.h[4], d.h[5], d.h[6], d.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}
