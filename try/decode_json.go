package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type Data struct {
	Number  int
	Comment string
}

var Data1 = Data{
	Number:  17,
	Comment: "Good",
}

func Handler(w http.ResponseWriter, req *http.Request) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content legnth\n", http.StatusBadRequest)
		return
	}
	reqBodyBuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var reqData map[string][]string
	if err := json.Unmarshal(reqBodyBuffer, &reqData); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
}

func NewHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body) // io.Reader <- 標準入力
	var reqData map[string][]string
	if err := decoder.Decode(&reqData); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
}

func JsonHandler(w http.ResponseWriter, req *http.Request) {
	//面倒な読み出し
	{
		//Request->JSON
		var reqBodyBuffer []byte
		//略）長さを持ってきて確保
		//err != io.EOFはダメらしい
		if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
			http.Error(w, "fail", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		var DataBody Data
		if err := json.Unmarshal(reqBodyBuffer, &DataBody); err != nil {
			return
		}

		//JSON->Response
		jsonData, _ := json.Marshal(DataBody)
		w.Write(jsonData)
	}

	//簡単な方
	{
		var DataBody Data
		if err := json.NewDecoder(req.Body).Decode(&DataBody); err != nil {
			return
		}

		if err := json.NewEncoder(w).Encode(DataBody); err != nil {
			return
		}
	}
}
