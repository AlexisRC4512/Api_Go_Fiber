package utils

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func SendData(url string, data interface{}) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }
    _, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    return nil
}