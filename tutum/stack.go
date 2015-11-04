package tutum

import (
	"encoding/json"
	"log"
)

/*
func ListStacks
Returns : Array of Stack objects
*/
func ListStacks() (StackListResponse, error) {
	url := "stack/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response StackListResponse
	var finalResponse StackListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse StackListResponse
			data, err := TutumCall(response.Meta.Next[8:], request, body)
			if err != nil {
				return nextResponse, err
			}
			err = json.Unmarshal(data, &nextResponse)
			if err != nil {
				return nextResponse, err
			}
			finalResponse.Objects = append(finalResponse.Objects, nextResponse.Objects...)
			response = nextResponse

		} else {
			break Loop
		}
	}

	return finalResponse, nil
}

/*
func GetStack
Argument : uuid
Returns : Stack JSON object
*/
func GetStack(uuid string) (Stack, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "stack/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func Export
Returns : String that contains the Stack details
*/
func (self *Stack) ExportStack() (string, error) {

	url := "stack/" + self.Uuid + "/export/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)

	data, err := TutumCall(url, request, body)
	if err != nil {
		return "", err
	}

	s := string(data)

	return s, nil
}

/*
func CreateStack
Argument : Stack JSON object (see documentation)
*/
func CreateStack(createRequest StackCreateRequest) (Stack, error) {
	url := "stack/"
	request := "POST"
	var response Stack

	newStack, err := json.Marshal(createRequest)
	if err != nil {
		return response, err
	}

	log.Println(string(newStack))

	data, err := TutumCall(url, request, newStack)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func Update
Argument : a Stack JSON object (see documentation)
*/
func (self *Stack) Update(createRequest StackCreateRequest) error {

	url := "stack/" + self.Uuid + "/"
	request := "PATCH"

	updatedStack, err := json.Marshal(createRequest)
	if err != nil {
		return err
	}

	log.Println(string(updatedStack))

	_, errr := TutumCall(url, request, updatedStack)
	if errr != nil {
		return errr
	}

	return nil
}

func (self *Stack) Start() error {

	url := "stack/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Stop() error {

	url := "stack/" + self.Uuid + "/stop/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Redeploy(reuse_volume ReuseVolumesOption) error {

	url := ""
	if reuse_volume.Reuse != true {
		url = "stack/" + self.Uuid + "/redeploy/?reuse_volumes=false"
	} else {
		url = "stack/" + self.Uuid + "/redeploy/"
	}

	log.Println(url)
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Terminate() error {

	url := "stack/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
