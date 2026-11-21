package currency

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type valCurs struct {
	Items []valuteNode `xml:"Valute"`
}

type valuteNode struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ParseXML(data []byte) ([]Rate, error) {
	var feed valCurs
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	rates := make([]Rate, 0, len(feed.Items))
	for _, entry := range feed.Items {
		numCode, err := strconv.Atoi(strings.TrimSpace(entry.NumCode))
		if err != nil {
			return nil, err
		}

		prepared := strings.ReplaceAll(strings.TrimSpace(entry.Value), ",", ".")
		val, err := strconv.ParseFloat(prepared, 64)
		if err != nil {
			return nil, err
		}

		rates = append(rates, Rate{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(entry.CharCode),
			Value:    val,
		})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}

func SaveJSON(outPath string, rates []Rate) error {
	dir := filepath.Dir(outPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	encoded, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(outPath, encoded, 0o644)
}
