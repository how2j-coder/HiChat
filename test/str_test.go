package test

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

// CompressString 使用 zlib 压缩字符串数据
func CompressString(data string) ([]byte, error) {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	defer writer.Close()

	_, err := writer.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	// 确保所有数据都写入压缩器
	writer.Close()
	return buf.Bytes(), nil
}

func DecompressString(compressedData []byte) (string, error) {
	reader, err := zlib.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	uncompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(uncompressedData), nil
}

func TestT(t *testing.T) {
	originalString := `
    eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmZvIjoiNzI2NzEyMDAwODQzMjY2NDU3NiIsImlzcyI6IkF1dGhfU2VydmVyIiwic3ViIjoiQXV0aF9TZXJ2ZXIiLCJhdWQiOlsiQW5kcm9pZF9BUFAiLCJXZWJfQVBQIl0sImV4cCI6MTczMjc3MzMyMSwibmJmIjoxNzMyNzczMzEyLCJpYXQiOjE3MzI3NzMzMTEsImp0aSI6IjcyNjcxMjAwMDg0MzI2NjQ1NzYifQ.B7_HmaIN3_K_B9eBz60jHK_2X15tNfVgKhawUC2edmEU7fMr-fhZrd-ZkQ5xk-xlKwj107n66dm7XXNjyzL0ieh_6wsRKgdqctI0YYdkrgNP5TMQEpdIax5tZGjY0FE7eW50_ageTYXsjL3bPq-rJOQbDvcRXjsam6HFHqOKPjrYVjuuIF1CCO7cAMG7b9-u1QhMiXntQMKSxsRM_ImtBH_MCc-tp5KOaiqsSsVKo9GHwXVsS2rv77Ra3S4pYf_HWz2QkyceLMBQThsbmbGYz0B9IzjZyTU2PREWgF5j4YbQrjM7fkGInsjSO0mfCmZHQ3vPWWK89UB6RcC4kIACnQ
789c1ccfddb26b3000c5f1ebb39fa22fd03d4a715c9652a1a1f1cd8d21a844d3ad7457e5e9cff45caff9cdacffd766b3d9b4abd3d7274c7ce28431073b8f8019dc03191b4001c3982586a37db7ab33d6ac7801fa433c0e448f9b2b3c1e16ff8838a4b9e81d63e93f64378ed70fb476cda9ef6231591b96ac802c044b09f978942542656951913962910172369cbe4911f16f3341a93c60a68d85a5e9b175f86c59913a1d4a74046ec20c58b2ff1c8311e690620972b8c27021357318a03f6f8fc3d5e398436eae67c319f30c119f9a12e440f238e4303267c046a10a8102285e3d8adf901e1678bc0af0d345d1cee339e9d0b7ae9636ab8027956ea96badce1581da6e29663bf9e975c9d5edab2536c4b66166ac7670da767d3135db6240f27bd8be6fee4277827a579486a959e6d1959f05d2f6a5b2cc817b6d1ef809843c6f86e9ea5de40822736c40f5969fc589e68265aa6d2a0b65756da33c9be959aa2f8fede4f8a83ebe7090d1b9628a6dd90fdfbdd0294fe8ef2fb07686e1abf8004f6aad6d7f77a88724bb3f1174c3f71cc012b0a76e97d0c0dbe728bb7e451e733827ee8f76b2972c9943717aa96a5049e17eccbbd24eb9888615b767a8a3a89f6b569f722ee81ae0b458a358bc04667ab564bacf6b3451a876c309dc671afa02eb0c56d8487a5dd2d4fdabc5ba1260633f808371475f7ffe050000ffffe6f2d34a
	`
	fmt.Println("Original String:", len(originalString))

	// 压缩字符串
	compressedData, err := CompressString(originalString)
	if err != nil {
		log.Fatalf("Failed to compress string: %v", err)
	}
	fmt.Printf("Compressed Data (length: %d):\n%x\n", len(compressedData), compressedData)

	// 解压缩数据
	decompressedString, err := DecompressString(compressedData)
	if err != nil {
		log.Fatalf("Failed to decompress data: %v", err)
	}
	fmt.Println("Decompressed String:", decompressedString)
}
