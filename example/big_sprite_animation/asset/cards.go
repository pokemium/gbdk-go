package asset

import (
	"gbdk/api/gb"
	"gbdk/api/macro"
)

var spriteBank = macro.Define(0)

var Cards = []gb.UINT8{
	0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0x81, 0x81,
	0x83, 0x83, 0x87, 0x87, 0x81, 0x81, 0x81, 0x81,
	0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81, 0x81,
	0x83, 0x83, 0x87, 0x87, 0x80, 0x80, 0xFF, 0xFF,
	0xFF, 0xFF, 0x01, 0x01, 0xC1, 0xC1, 0xC1, 0xC1,
	0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1,
	0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1, 0xC1,
	0xE1, 0xE1, 0xF1, 0xF1, 0x01, 0x01, 0xFF, 0xFF,
	0xFF, 0xFF, 0x80, 0x80, 0x83, 0x83, 0x87, 0x87,
	0x8E, 0x8E, 0x8C, 0x8C, 0x80, 0x80, 0x80, 0x80,
	0x80, 0x80, 0x80, 0x80, 0x81, 0x81, 0x83, 0x83,
	0x87, 0x87, 0x8F, 0x8F, 0x80, 0x80, 0xFF, 0xFF,
	0xFF, 0xFF, 0x01, 0x01, 0xC1, 0xC1, 0xE1, 0xE1,
	0x71, 0x71, 0x31, 0x31, 0x31, 0x31, 0x31, 0x31,
	0x71, 0x71, 0xE1, 0xE1, 0xC1, 0xC1, 0x81, 0x81,
	0xF1, 0xF1, 0xF1, 0xF1, 0x01, 0x01, 0xFF, 0xFF}