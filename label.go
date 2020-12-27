package main

import (
	"strconv"
)

// RowLabel returns the label for given row index
func RowLabel(row int) string {
	return strconv.Itoa(row + 1)
}

// ColLabel returns the label for given col index
// https://stackoverflow.com/a/182924
func ColLabel(col int) string {
	runes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	runeLen := len(runes)
	label := ""
	n := col + 1
	var mod int
	for {
		if n <= 0 {
			break
		}
		mod = (n - 1) % runeLen
		label = string(runes[mod]) + label
		n = (n - mod) / runeLen
	}
	return label
}
