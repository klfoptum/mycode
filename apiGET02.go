package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        "time"
)

type SpaceX []struct {
        CapsuleSerial      string    `json:"capsule_serial"`
        CapsuleID          string    `json:"capsule_id"`
        Status             string    `json:"status"`
        OriginalLaunch     time.Time `json:"original_launch"`
        OriginalLaunchUnix int       `json:"original_launch_unix"`
        Missions           []struct {
                Name   string `json:"name"`
                Flight int    `json:"flight"`
        } `json:"missions"`
        Landings   int    `json:"landings"`
        Type       string `json:"type"`
        Details    string `json:"details"`
        ReuseCount int    `json:"reuse_count"`
}

func main() {

    url := "https://api.spacexdata.com/v3/capsules"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }

    client := &http.Client{}

    resp, err := client.Do(req)
        
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }

    defer resp.Body.Close()

    var record SpaceX

    // Use json.Decode for reading streams of JSON data
    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        log.Println(err)
    }

    // the _ is an int
    // launchData is the data
    for launchNo, launchData := range record {
        fmt.Println("Capsule Record     =\n", launchData)
        fmt.Println("Record Number      =", launchNo)
        fmt.Println("Capsule Serial     =", launchData.CapsuleSerial)
    }
}
