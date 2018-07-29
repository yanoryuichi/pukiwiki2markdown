package main

import (
	"regexp"
	"strconv"
	"strings"
)

var MARK_FOR_CODE_LINE = "			"

func convAll(text string) string {
	text = convCode(text)
	text = convHeaders(text)
	text = convLists(text)
	text = convOrderedLists(text)
	text = convDefinitionLists(text)
	text = removeInterlink(text)
	text = removeMarkForCodeLine(text)

	return text
}

func convHeaders(text string) string {
	reh3 := regexp.MustCompile("(?m)^\\*\\*\\*(.+)$")
	text = reh3.ReplaceAllString(text, "###$1")

	reh2 := regexp.MustCompile("(?m)^\\*\\*(.+)$")
	text = reh2.ReplaceAllString(text, "##$1")

	reh1 := regexp.MustCompile("(?m)^\\*(.+)$")
	text = reh1.ReplaceAllString(text, "#$1")

	return text
}

func convLists(text string) string {
	buf := []string{}

	isPreviousLineBlank := true
	isPreviousLineInsideList := false

	markMap := map[int]string{1: "- ", 2: "  - ", 3: "     - "}

	re := regexp.MustCompile("\r\n|\n\r|\n|\r")
	lines := re.Split(text, -1)

	for _, currentLine := range lines {
		if strings.Index(currentLine, "-") == 0 {
			if !isPreviousLineInsideList && !isPreviousLineBlank {
				buf = append(buf, "")
			}

			for i := 3; i >= 1; i-- {
				mark1 := strings.Repeat("-", i)
				mark2 := markMap[i]
				re := regexp.MustCompile("^" + mark1 + "\\s*(.+)$")
				if re.MatchString(currentLine) {
					ret := re.ReplaceAllString(currentLine, mark2+"$1")
					buf = append(buf, ret)
					break
				}
			}

			isPreviousLineInsideList = true
			isPreviousLineBlank = false
		} else {
			if isPreviousLineInsideList && len(strings.TrimSpace(currentLine)) != 0 {
				buf = append(buf, "")
			}

			buf = append(buf, currentLine)

			isPreviousLineInsideList = false
			isPreviousLineBlank = false
			if len(strings.TrimSpace(currentLine)) == 0 {
				isPreviousLineBlank = true
			}
		}
	}

	converted := strings.Join(buf, "\n")
	return converted
}

func convOrderedLists(text string) string {
	buf := []string{}

	isPreviousLineBlank := true
	isPreviousLineInsideList := false
	orderNum := 1

	re := regexp.MustCompile("\r\n|\n\r|\n|\r")
	lines := re.Split(text, -1)

	for _, currentLine := range lines {

		re := regexp.MustCompile("^\\+ ?(.+)$")
		ret := re.FindStringSubmatch(currentLine)

		if ret != nil {
			if !isPreviousLineInsideList && !isPreviousLineBlank {
				buf = append(buf, "")
			}
			buf = append(buf, strconv.Itoa(orderNum)+". "+ret[1])
			orderNum++
			isPreviousLineInsideList = true
		} else {
			if isPreviousLineInsideList {
				buf = append(buf, "")
			}
			buf = append(buf, currentLine)
			orderNum = 1
			isPreviousLineInsideList = false
		}

		isPreviousLineBlank = false
		if len(strings.TrimSpace(currentLine)) == 0 {
			isPreviousLineBlank = true
		}
	}

	converted := strings.Join(buf, "\n")
	return converted
}

func convDefinitionLists(text string) string {

	buf := []string{}

	isPreviousLineBlank := true
	isPreviousLineInsideList := false

	re := regexp.MustCompile("\r\n|\n\r|\n|\r")
	lines := re.Split(text, -1)

	for _, currentLine := range lines {
		re := regexp.MustCompile("^:\\s*(.+)\\|\\s*(.+)$")
		ret := re.FindStringSubmatch(currentLine)

		if ret != nil {
			if !isPreviousLineBlank || isPreviousLineInsideList {
				buf = append(buf, "")
			}

			dt, dd := ret[1], ret[2]
			buf = append(buf, dt)
			buf = append(buf, ": "+dd)

			isPreviousLineBlank = true
			isPreviousLineInsideList = true
		} else {
			if isPreviousLineInsideList && len(strings.TrimSpace(currentLine)) != 0 {
				buf = append(buf, "")
			}

			buf = append(buf, currentLine)

			isPreviousLineInsideList = false

			if len(strings.TrimSpace(currentLine)) == 0 {
				isPreviousLineBlank = true
			} else {
				isPreviousLineBlank = false
			}
		}
	}

	converted := strings.Join(buf, "\n")
	return converted
}

func convCode(text string) string {
	buf := []string{}

	isPreviousLineBlank := true
	isPreviousLineInsideCode := false

	re := regexp.MustCompile("\r\n|\n\r|\n|\r")
	lines := re.Split(text, -1)

	for _, currentLine := range lines {
		re := regexp.MustCompile("^ (.*)$")
		ret := re.FindStringSubmatch(currentLine)

		if ret != nil {
			// Current line is CODE

			if !isPreviousLineInsideCode && !isPreviousLineBlank {
				buf = append(buf, "")
			}

			if !isPreviousLineInsideCode {
				buf = append(buf, "```clike")
			}

			buf = append(buf, MARK_FOR_CODE_LINE+ret[1])

			isPreviousLineInsideCode = true
			isPreviousLineBlank = false
		} else {
			// Current line is NOT CODE

			if isPreviousLineInsideCode {
				buf = append(buf, "```")
			}

			if isPreviousLineInsideCode && len(strings.TrimSpace(currentLine)) != 0 {
				buf = append(buf, "")
			}

			buf = append(buf, currentLine)

			isPreviousLineInsideCode = false

			if len(strings.TrimSpace(currentLine)) == 0 {
				isPreviousLineBlank = true
			} else {
				isPreviousLineBlank = false
			}
		}
	}

	converted := strings.Join(buf, "\n")
	return converted
}

func removeInterlink(text string) string {
	re := regexp.MustCompile(" ?\\[#[a-z0-9]+?\\]")
	text = re.ReplaceAllString(text, "")
	return text
}

func removeMarkForCodeLine(text string) string {
	re := regexp.MustCompile("(?m)^" + MARK_FOR_CODE_LINE)
	text = re.ReplaceAllString(text, "")
	return text
}
