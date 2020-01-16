package cmd

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
)

const HEADER_CODE   = "00"
const LINE_CODE     = "01"
const FOOTER_CODE   = "99"
const CODE_OFFSET = 2

type LineRecord struct {
    LineMarketDate time.Time
    LineBdiCode int
    LineStockCode string
    LineMarketType int
    LineCompanyName string
    LineStockSpec string
    LineTermDays int
    LineCurrency string
    LineOpeningPrice int
    LineMaxPrice int
    LineMinPrice int
    LineMedPrice int
    LineLastPrice int
    LineBestBuyPrice int
    LineBestSellPrice int
    LineNegTotal int
    LineTitTotal int
    LineVolumeTotal int
}

func (line LineRecord) Persist() {

}

func ImportHandler() {
    filePath, _ := filepath.Abs(*FileInput)
    fileHandler, _ := os.Open(filePath)
    defer fileHandler.Close()
    scanner := bufio.NewScanner(fileHandler)
    lines := make([][]rune, 1)

    for scanner.Scan() {
        line := []rune(scanner.Text())
        code := string(line[0:2])
        switch code {
        case LINE_CODE:
            lines = append(lines, line)
        }
    }

    ImportLines(lines)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func ImportLines(lines [][]rune) {
    for i, rawLine := range lines {
        parsedLine := parseLine(rawLine)
    }
}

func parseLine(rawLine []rune) *LineRecord {

}
