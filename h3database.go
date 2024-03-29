package h3light

type DatabaseCell string

// These tables are private in the stdlib, but we need them to do byte <-> single char
// conversion instead of the dual-char in normal hex, so replicated them here.
const (
	hextable        = "0123456789abcdef"
	reverseHexTable = "" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\xff\xff\xff\xff\xff\xff" +
		"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"
)

func DatabaseCellFromCell(cell uint64) DatabaseCell {
	c := Cell(cell)
	res := c.Resolution()

	buf := make([]byte, res+2)
	buf[0] = hextable[byte((c>>49)&0b111)]
	buf[1] = hextable[byte((c>>45)&0b1111)]
	for i := 1; i < res+1; i++ {
		buf[i+1] = hextable[byte((c>>(45-i*3))&0b111)]
	}

	return DatabaseCell(buf)
}

func (dc DatabaseCell) Resolution() int {
	return len(dc) - 2
}

func (dc DatabaseCell) Int64() int64 {
	cell := uint64(0x800000000000000)

	res := dc.Resolution()

	// set resolution
	cell |= uint64(res) << 52

	// set res 0
	cell |= uint64(reverseHexTable[dc[0]]<<4|reverseHexTable[dc[1]]) << 45

	// set res i
	for i := 1; i < (res + 1); i++ {
		cell |= uint64(reverseHexTable[dc[i+1]]) << (45 - i*3)
	}

	// set empty res to 0x7 according to H3 format
	for i := res + 1; i < 16; i++ {
		cell |= uint64(7) << (45 - i*3)
	}

	return int64(cell)
}

func (dc DatabaseCell) Cell() Cell {
	return Cell(dc.Int64())
}

func (dc *DatabaseCell) CellPtr() *Cell {
	if dc == nil {
		return nil
	}

	ret := dc.Cell()

	return &ret
}

func (dc DatabaseCell) Parent(res int) DatabaseCell {
	if dc.Resolution() < res {
		return ""
	}

	return dc[:res+2]
}
