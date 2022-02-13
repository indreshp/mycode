package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
)

type MarsPhoto struct {
	Photos []struct {
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
	} `json:"photos"`
}
 
func main() {
    //url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&page=2&api_key=DEMO_KEY"
    url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&api_key=DEMO_KEY"
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

    var record  MarsPhoto // MarsPhoto
    if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
                log.Println(err)
        }
     //fmt.Println(record)
    mastCount, fhazCount, rhazCount, chemcamCount, mahliCount, mardiCount, navcamCount, 
               pancamCount, minitesCount := 0, 0, 0, 0, 0, 0, 0, 0, 0

     for _, data := range record.Photos {
        // fmt.Println(rec_id)
	if data.Rover.Status == "active" {
	   switch {
 	     case data.Camera.Name == "MAST":
	          mastCount++
   	     case data.Camera.Name == "FHAZ":
	          fhazCount++
	     case data.Camera.Name == "RHAZ":
	          rhazCount++
 	     case data.Camera.Name == "CHEMCAM":
	          chemcamCount++
   	     case data.Camera.Name == "MAHLI":
	          mahliCount++
	     case data.Camera.Name == "MARDI":
	          mardiCount++
 	     case data.Camera.Name == "NAVCAM":
	          navcamCount++
   	     case data.Camera.Name == "PANCAM":
	          pancamCount++
	     case data.Camera.Name == "MINITES":
	          minitesCount++
	  }
        }
     }
     fmt.Println("MAST Camera image count --> ", mastCount)
     fmt.Println("RHAZ Camera image count --> ", rhazCount)
     fmt.Println("FHAZ Camera image count --> ", fhazCount)
     fmt.Println("CHEMCAM Camera image count --> ", chemcamCount)
     fmt.Println("MAHLI Camera image count --> ", mahliCount)
     fmt.Println("MARDI Camera image count --> ", mardiCount)
     fmt.Println("NAVCAM Camera image count --> ", navcamCount)
     fmt.Println("PANCAM Camera image count --> ", pancamCount)
     fmt.Println("MINITES Camera image count --> ", minitesCount)
     

}





