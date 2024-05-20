package texture

import (
	"fmt"
	tafilesread "totala_reader/ta_files_read"
)

var describeActions bool

type GafEntry struct {
	Name   string
	Frames []*GafFrame
}

type GafFrame struct {
	Pixels [][]uint8 // each value is an index from palette
}

func ReadTextureFromReader(r *tafilesread.Reader, verbose bool) []*GafEntry {
	describeActions = verbose
	// Reading header
	version := r.ReadIntFromBytesArray(0, 0)
	entries := r.ReadIntFromBytesArray(0, 4)
	always0 := r.ReadIntFromBytesArray(0, 8)
	if describeActions {
		fmt.Printf("GAF file %s\nVersion %d, entries %d, should be zero %d\n", r.FileName, version, entries, always0)
	}
	var entryPointers []int
	for i := 0; i < entries; i++ {
		entryPointers = append(entryPointers, r.ReadIntFromBytesArray(12, i*4))
	}
	if describeActions {
		fmt.Printf("Pointers acquired.\n")
	}

	entriesArray := make([]*GafEntry, entries)

	for index, off := range entryPointers {
		frames := r.ReadUint16FromBytesArray(off, 0)
		always1 := r.ReadUint16FromBytesArray(off, 2)
		always0 = r.ReadIntFromBytesArray(off, 4)
		name := r.ReadFixedLengthStringFromBytesArray(off, 8, 32)
		if describeActions {
			fmt.Printf("GAF entry #%d at offset %d:\n", index, off)
			fmt.Printf("  Name \"%s\", %d frames, %d should be one, %d should be zero\n",
				name, frames, always1, always0)
		}
		entry := &GafEntry{
			Name: name,
		}
		// Read each GAF frame entry for GAF entry
		frameEntryOffset := 0
		for gfe := 0; gfe < frames; gfe++ {
			ptrFrameEntry := r.ReadIntFromBytesArray(off, 40+frameEntryOffset)
			unknown := r.ReadIntFromBytesArray(off, 44+frameEntryOffset)
			if describeActions {
				fmt.Printf("    GAF frame entry %d:\n", gfe)
				fmt.Printf("      Pointer to the data: %d, unknown value: %d\n", ptrFrameEntry, unknown)
			}
			entry.Frames = append(entry.Frames, readGafFrameData(r, ptrFrameEntry))
			frameEntryOffset += 8
		}
		entriesArray[index] = entry
	}
	return entriesArray
}

func readGafFrameData(r *tafilesread.Reader, offset int) *GafFrame {
	width := r.ReadUint16FromBytesArray(offset, 0)
	height := r.ReadUint16FromBytesArray(offset, 2)
	xPos := r.ReadUint16FromBytesArray(offset, 4)
	yPos := r.ReadUint16FromBytesArray(offset, 6)
	unknownByte := r.ReadByteFromBytesArray(offset, 8)
	compressed := r.ReadByteFromBytesArray(offset, 9) != 0
	framePointers := r.ReadUint16FromBytesArray(offset, 10)
	unknown2 := r.ReadIntFromBytesArray(offset, 12)
	ptrFrameData := r.ReadIntFromBytesArray(offset, 16)
	unknown3 := r.ReadIntFromBytesArray(offset, 20)
	if describeActions {
		fmt.Printf("      GAF Frame Data: \n")
		fmt.Printf("        %dx%dpx, xPos %d, yPos %d\n", width, height, xPos, yPos)
		fmt.Printf("        Unknown1: %d, compressed %v\n", unknownByte, compressed)
		fmt.Printf("        Frame pointers: %d, Unknown2 %d\n", framePointers, unknown2)
		fmt.Printf("        PtrFrameData: %d, Unknown3 %d\n", ptrFrameData, unknown3)
	}

	frame := &GafFrame{}

	// 	If the FramePointers
	// member were not 0, then instead of pixels, PtrFrameData would point to a
	// list of pointers that had that many entries.
	if framePointers > 1 {
		panic("Multi frame-pointer reading is not implemented yet!")
		frame.Pixels = [][]byte{{0}}
		return frame
	}
	// read the raw data itself
	if compressed {
		frame.Pixels = readCompressedPixels(r, ptrFrameData, width, height)
	} else {
		frame.Pixels = readUncompressedPixels(r, ptrFrameData, width, height)
	}
	return frame
}

func readUncompressedPixels(r *tafilesread.Reader, offset, width, height int) [][]uint8 {
	// Row and column indices (width and height) must be swapped, as the data is written row-by-row and read column-by-column
	var pixels = make([][]uint8, width)
	for i := 0; i < width; i++ {
		pixels[i] = make([]uint8, height)
		for j := 0; j < height; j++ {
			index := i + j*width
			pixels[i][j] = r.ReadByteFromBytesArray(offset, index)
		}
	}
	return pixels
}

func readCompressedPixels(r *tafilesread.Reader, offset, width, height int) [][]uint8 {
	var pixels = make([][]uint8, width)
	for i := 0; i < width; i++ {
		pixels[i] = make([]uint8, height)
	}

	currOffset := 0
	fmt.Printf("Reading compressed entry... Size %dx%d;\n", width, height)

	for currY := 0; currY < height; currY++ {
		currX := 0
		thisLineOffset := 0
		thisLineBytes := r.ReadUint16FromBytesArray(offset, currOffset)
		thisLineOffset += 2

		for thisLineOffset < thisLineBytes+2 {
			mask := r.ReadByteFromBytesArray(offset, currOffset+thisLineOffset)
			thisLineOffset++

			if mask&0x01 == 0x01 { // transparency: skip (mask >> 1) pixels
				currX += int(mask >> 1)
			} else if mask&0x02 == 0x02 { // copy next byte ((mask >> 2) + 1) times
				nextByte := r.ReadByteFromBytesArray(offset, currOffset+thisLineOffset)
				thisLineOffset++
				for i := byte(0); i < (mask>>2)+1; i++ {
					pixels[currX][currY] = nextByte
					currX++
				}
			} else { // copy next ((mask >> 0x02) + 1) bytes to output
				for i := byte(0); i < (mask>>0x02)+1; i++ {
					nextByte := r.ReadByteFromBytesArray(offset, currOffset+thisLineOffset)
					thisLineOffset++
					pixels[currX][currY] = nextByte
					currX++
				}
			}
		}
		currOffset += thisLineOffset
	}

	return pixels
}
