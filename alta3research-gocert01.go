/*

Mars rover cameras image count
Using NASA's rest APIs collect data from different cameras 
Include camera's with status active only

*/

package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
)
// Mars rover camera photo data structure type definition
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
   // NASA API with demo key	
   url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&api_key=DEMO_KEY"
   // http get method
   method := "GET"
   // defile client an http client
   client := &http.Client{}

   // building new http request
   req, err := http.NewRequest(method, url, nil)
   if err != nil {
      fmt.Println(err)
   }

   // do send http request and get http response
   res, err := client.Do (req)
   if err != nil{
        fmt.Println(err)
        return
   }
   // close connection even if any problem occurs
   defer res.Body.Close()
   // declare record of type MarsPhoto structure
   var record  MarsPhoto 
   // use json decoder to read streams of JSON data
   if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
                log.Println(err)
        }
   // declare image count variables of different cameras	
   mastCount, fhazCount, rhazCount, chemcamCount, mahliCount, mardiCount, navcamCount, 
               pancamCount, minitesCount := 0, 0, 0, 0, 0, 0, 0, 0, 0
   // for each record increment the image count of the specific camera if status flag is active
   for _, data := range record.Photos {
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

    // print report camera name Vs. image count
    fmt.Println("Camera --> image count")
    fmt.Println("MAST   --> ", mastCount)
    fmt.Println("RHAZ   --> ", rhazCount)
    fmt.Println("FHAZ   --> ", fhazCount)
    fmt.Println("CHEMCAM--> ", chemcamCount)
    fmt.Println("MAHLI  --> ", mahliCount)
    fmt.Println("MARDI  --> ", mardiCount)
    fmt.Println("NAVCAM --> ", navcamCount)
    fmt.Println("PANCAM --> ", pancamCount)
    fmt.Println("MINITES--> ", minitesCount)
}





