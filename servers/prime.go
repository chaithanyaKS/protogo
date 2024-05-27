package servers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
)

type PrimeRequest struct {
	Method string `json:"method"`
	Number int `json:"number"`

}

type PrimeResponse struct {
	Method string `json:"method"`
	Prime bool `json:"prime"`

}

func HandlePrimeConnection(c net.Conn) {
	buffer := make([]byte, 2048)
	data := make([]byte, 4048)
	defer c.Close()

	for {
		_, err := c.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("IO Error")
			}
			break
		}
		data = append(data, buffer...)
	}
	data = bytes.Trim(data, "\x00")
	jsonData, err := extractJsonData(data)
	if err != nil {
		errorString := fmt.Errorf("invaild request: %v", err)
		c.Write([]byte(errorString.Error()))
		return
	}
	isPrime := checkIsPrime(jsonData.Number)
	fmt.Println(jsonData)
	response := PrimeResponse{Method: "isPrime", Prime: isPrime}
	responseData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error while marshalling data", err)
	}
	c.Write(responseData)

}

func checkIsPrime(num int) bool {
	if num < 2 {
		fmt.Println("Number must be greater than 2.")
		return false
	}
	if num %2 == 0 {
		fmt.Println("Number must be an odd number")
		return false
	}
	sqRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqRoot; i+=2 {
		if num%i == 0 {
			fmt.Println("Non Prime Number")
			return false
		}
	}
	fmt.Println("Prime Number")
	return true
}



func extractJsonData(data []byte) (PrimeRequest, error) {
	var res PrimeRequest
	err := json.Unmarshal(data,&res)
	if err != nil {
		return PrimeRequest{}, err
	}
	if res.Method != "isPrime" {
		return PrimeRequest{}, errors.New("invalid method")
	}
	return res, nil
}