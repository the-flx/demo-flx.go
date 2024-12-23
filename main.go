/**
 * $File: main.go $
 * $Date: 2024-12-21 16:17:05 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package main

import (
	"fmt"

	"github.com/the-flx/flx.go"
)

func main() {
	fmt.Println(flx.Score("switch-to-buffer", "stb").Score)
}
