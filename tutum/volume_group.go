package tutum

import "encoding/json"

type VolumeGroupListResponse struct {
	Objects []VolumeGroup `json:"objects"`
}

type VolumeGroup struct {
	Name         string   `json:"name"`
	Resource_uri string   `json:"resource_uri"`
	Services     []string `json:"services"`
	State        string   `json:"state"`
	Uuid         string   `json:"uuid"`
	Volume       []string `json:"volume"`
}

func ListVolumeGroups() ([]VolumeGroup, error) {

	url := "volumegroup/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeGroupListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Objects, nil
}

func GetVolumeGroup(uuid string) (VolumeGroup, error) {

	url := "volumegroup/" + uuid + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response VolumeGroup

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil

}
