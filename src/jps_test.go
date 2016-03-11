package jpsplus

import (
	"fmt"
	"testing"
)

func TestJPSplus(*testing.T) {
	// mapData := make([]bool, 0, 10000)
	// thePath := make([]Point, 0, 10000)

	ok := GetMapFromImage("map10x10.png")
	if ok {
		fmt.Printf("width = %v  heigth = %v\n", DefaultBoolMap.width(), DefaultBoolMap.height())
		fmt.Println(DefaultBoolMap.toString())
		PreprocessMap("mapPreprocessedFilename")
		jps := newJPSPlus()
		s := xyLocJPS{0, 0}
		g := xyLocJPS{0, 1}
		path, ok := jps.GetPath(s, g)
		fmt.Printf("ok = %v , path = %v\n", ok, path)
		// fmt.Printf("jps.m_simpleUnsortedPriorityQueue = %#v\n", jps.m_simpleUnsortedPriorityQueue)
		// fmt.Printf("jps.m_fastStack = %#v\n", jps.m_fastStack)
		// fmt.Printf("jps.m_mapNodes = %#v\n", jps.m_mapNodes)
		fmt.Printf("jps.m_currentIteration = %#v\n", jps.m_currentIteration)
		// fmt.Printf("jps.m_goalNode = %#v\n", jps.m_goalNode)
		// fmt.Printf("jps.m_goalRow = %#v\n", jps.m_goalRow)
		// fmt.Printf("jps.m_goalCol = %#v\n", jps.m_goalCol)

	} else {
		fmt.Println("open mapfile faild")
	}

	// reference := PrepareForSearch(mapData, width, height, "mapPreprocessedFilename")
	// thePath, ok := GetPath(reference, Point{1, 1}, Point{10, 10})
}
