package table

import (
	"strings"
)

const (
	pipe       = `|`
	hyphen     = `-`
	whitespace = ` `
)

// Table prints the input data in tabulated format.
// The characters `|` & `-` are used as horizontal & verticle delimiters
// of the table respectively.
func Table(data [][]string, headerSeperator bool) []string {
	if len(data) == 0 {
		// When there are no records in table
		return []string{}
	}

	var (
		columnLengthMap                map[uint]uint
		sumOfColumnLengths, tableWidth uint
		recordLines, table             []string
		horizontalLine                 string

		// Table border characters and seperators
		leadingDelimiter  = pipe + whitespace
		verticleSeperator = whitespace + pipe + whitespace
		trailingDelimiter = whitespace + pipe
	)

	// Create map of column index -> max column width required
	columnLengthMap = calculateMaxLengthOfColumns(data)

	// Sum of max column widths of all columns
	sumOfColumnLengths = sumOfValues(columnLengthMap)

	// Total width of table = sum of max column widths +
	// 						  length of seperator with spaces ` | ` multiplied by number of seperators +
	// 						  length of leading border character with space `| `
	// 						  length of trailing border character with space ` |`
	tableWidth = uint(int(sumOfColumnLengths) +
		((len(columnLengthMap) - 1) * len(verticleSeperator)) +
		len(leadingDelimiter) + len(trailingDelimiter))

	// Create horizontal seperator line `-----` with total length. This will be used in
	// 		1. First line of table
	// 		2. Optionally, header seperator
	// 		3. Last line of table
	horizontalLine = buildHorizontalLine(tableWidth, hyphen)

	// Build data record's lines with leading, trailing and verticle seperators
	recordLines = buildRecordLines(data, verticleSeperator, leadingDelimiter, trailingDelimiter, columnLengthMap)

	// Add table start horizontal line and header record to result
	table = append(table, horizontalLine, recordLines[0])

	// Add header seperator horizontal line to result
	if headerSeperator && len(data) > 1 {
		// When header seperator is enabled
		table = append(table, horizontalLine)
	}

	// Add data records table closing horizontal line to result
	table = append(table, recordLines[1:]...)
	table = append(table, horizontalLine)

	return table
}

// calculateMaxLengthOfColumns calculates max length of each column
// and maps it against column index
func calculateMaxLengthOfColumns(data [][]string) map[uint]uint {
	maxLens := make(map[uint]uint)

	for _, record := range data {
		for i, cell := range record {
			cellLen := uint(len(cell))

			if v, ok := maxLens[uint(i)]; ok {
				// When index already exists in map
				if cellLen < v {
					// When current value is less than what we have in map.
					// Ignore it as we are looking for highest values.
					continue
				}

				// When current value is higher than what we have in map.
				// We update the value in map outside the conditional scope.
			}

			// When index doesn't exist in map.
			// We add the index and first value in map.
			maxLens[uint(i)] = cellLen
		}
	}

	return maxLens
}

// sumOfValues gives the total sum of values of the map
func sumOfValues(m map[uint]uint) uint {
	var sum uint

	for _, v := range m {
		sum += v
	}

	return sum
}

// buildHorizontalLine builds horizontal line with specified
// input character and size
func buildHorizontalLine(size uint, char string) string {
	return strings.Repeat(char, int(size))
}

// buildRecordLines converts 2D slice data to 1D record lines
// with added leading, trailing and seperator lines
func buildRecordLines(data [][]string, sep, leading, trailing string, sizeMap map[uint]uint) []string {
	var (
		line, cell    string
		lines, record []string
		index         int
	)

	for _, record = range data {
		// Add leading character.
		line = leading

		for index, cell = range record {
			if index > 0 {
				// When it is not first index, we add the seperator in the start.
				line += sep
			}

			// Add extra space to inline the current cell with highest cell in the column
			line += cell + strings.Repeat(whitespace, int(sizeMap[uint(index)])-len(cell))
		}

		// Add trailing character
		line += trailing
		lines = append(lines, line)
	}

	return lines
}
