package neo4jclient

import "github.com/hippoai/graphgo"

// RequestAndConvert requests the database
// and converts it to Graphgo format
func (neo *Neo) RequestAndConvert(payload *Payload) (*graphgo.Output, error) {

	response, err := neo.Request(payload)
	if err != nil {
		return nil, errNeo4JRequest()
	}

	return Convert(response)

}

// RequestConvertAndGetN
func (neo *Neo) RequestConvertAndGetN(payload *Payload) (*graphgo.Output, int, error) {
	response, err := neo.Request(payload)
	if err != nil {
		return nil, 0, errNeo4JRequest()
	}

	out, err := Convert(response)
	if err != nil {
		return nil, 0, err
	}

	n, err := response.GetN()
	if err != nil {
		return nil, 0, err
	}

	return out, n, nil

}

// RequestAndGetN
func (neo *Neo) RequestAndGetN(payload *Payload) (*Response, int, error) {
	response, err := neo.Request(payload)
	if err != nil {
		return nil, 0, errNeo4JRequest()
	}

	n, err := response.GetN()
	if err != nil {
		return nil, 0, err
	}

	return response, n, nil

}
