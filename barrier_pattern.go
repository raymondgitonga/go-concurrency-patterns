package go_concurrency_patterns

import "fmt"

type barrierResp struct {
	Err error
	Resp string
}

func barrier(endpoints ...string)  {

	// get number of endpoints to be used in buffered channel
	requestNumber := len(endpoints)

	// create channel that will be of type barrierResp and buffered size od endpoint len
	in:= make(chan barrierResp, requestNumber)

	// defer closing of channel until after fun is done
	defer close(in)

	// create responses array to hold responses from network call
	responses := make([]barrierResp, requestNumber)

	// loop through the enpoints while making network calls
	for _, endpoint := range endpoints{
		// make network calls and pass responses into the channel in
		go makeRequest(in, endpoint)
	}

	var hasError bool

	// loop trough the endpoint len
	for i := 0; i< requestNumber; i++ {
		// add responses from channel that holds the barrier resp into a variable
		resp:= <-in

		if resp.Err != nil {
			fmt.Println("Error: ", resp.Err)
			hasError = true
		}
		// add responses to response array
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses{
			fmt.Println(resp.Resp)
		}
	}

}