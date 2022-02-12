package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
)

//type MarsPhoto []struct {
type	MarsPhotos []struct {
		ID     int `json:"id"`
		Sol    int `json:"sol"`
		Camera struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RoverID  int    `json:"rover_id"`
			FullName string `json:"full_name"`
		} `json:"camera"`
		ImgSrc    string `json:"img_src"`
		EarthDate string `json:"earth_date"`
		Rover     struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			LandingDate string `json:"landing_date"`
			LaunchDate  string `json:"launch_date"`
			Status      string `json:"status"`
		} `json:"rover"`
	} // `json:"photos"`
//}
 
func main() {
    url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&page=2&api_key=DEMO_KEY"
   method := "GET"
   client := &http.Client{}

   req, err := http.NewRequest(method, url, nil)
   if err != nil {
      fmt.Println(err)
   }

   res, err := client.Do (req)
   if err != nil{
        fmt.Println(err)
        return
    }
    defer res.Body.Close()
    //fmt.Println(res.Body)

    var record  MarsPhotos // MarsPhoto
    if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
                log.Println(err)
        }
     //fmt.Println(record)
     
     for ID := range record {
        fmt.Println(ID)
     }


}





