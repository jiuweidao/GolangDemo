package main

type ActCountry struct {
	DeviceData []DeviceData `json:"data"`
}

type DeviceData struct {
	Item  string `json:"item"`
	Count int    `json:"count"`
}

func main() {
	resp := ActCountry{}

	//return nil, err
	resp.DeviceData = append(resp.DeviceData, DeviceData{
		"no data",
		1,
	})


	//if resp{}

}
