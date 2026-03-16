package network

type ResponseStatus int

const (
	ResponseStatusSuccess ResponseStatus = 200
	ResponseStatusFailed  ResponseStatus = 400
)
