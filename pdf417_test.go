package pdf417

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode2(t *testing.T) {
	assert := assert.New(t)

	data := "Ruud"

	barcode := Encode(data, 6, 2, false)

	assert.Equal(barcode.Data, data)
	assert.Equal(barcode.Columns, DEFAULT_COLUMNS)
	assert.Equal(barcode.Rows, 2)
	assert.Equal(barcode.SecurityLevel, DEFAULT_SECURITY_LEVEL)
	assert.Equal(barcode.CodeWords, []int{4, 537, 620, 119, 266, 457, 253, 518, 74, 589, 901, 37})
	assert.Equal(
		[][]int{
			[]int{
				130728,
				120256,
				125560,
				78210,
				116376,
				106894,
				74544,
				117894,
				128318,
				260649,
			},
			[]int{
				130728,
				129678,
				104332,
				123494,
				125192,
				115830,
				114524,
				120588,
				128352,
				260649,
			},
		},
		barcode.Codes,
	)
	assert.Equal(
		[][]bool{[]bool{true, true, true, true, true, true, true, true, false, true, false, true, false, true, false, false, false, true, true, true, false, true, false, true, false, true, true, true, false, false, false, false, false, false, true, true, true, true, false, true, false, true, false, false, true, true, true, true, false, false, false, true, false, false, true, true, false, false, false, true, true, false, false, false, false, false, true, false, true, true, true, false, false, false, true, true, false, true, false, false, true, true, false, false, false, true, true, false, true, false, false, false, false, true, true, false, false, false, true, true, true, false, true, false, false, true, false, false, false, true, true, false, false, true, true, false, false, false, false, true, true, true, false, false, true, true, false, false, true, false, false, false, false, true, true, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, true, true, true, true, true, true, false, true, false, false, false, true, false, true, false, false, true}, []bool{true, true, true, true, true, true, true, true, false, true, false, true, false, true, false, false, false, true, true, true, true, true, true, false, true, false, true, false, false, false, true, true, true, false, true, true, false, false, true, false, true, true, true, true, false, false, false, true, true, false, false, true, true, true, true, false, false, false, true, false, false, true, true, false, false, true, true, false, true, true, true, true, false, true, false, false, true, false, false, false, false, true, false, false, false, true, true, true, false, false, false, true, false, false, false, true, true, true, false, true, true, false, true, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, false, true, true, true, false, true, false, true, true, true, false, false, false, false, true, true, false, false, true, true, true, true, true, false, true, false, true, false, true, true, false, false, false, false, false, true, true, true, true, true, true, true, false, true, false, false, false, true, false, true, false, false, true}},
		barcode.PixelGrid(),
	)
}
