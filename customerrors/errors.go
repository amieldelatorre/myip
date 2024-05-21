package customerrors

type EncodeResponseError struct{}

func (err *EncodeResponseError) Error() string {
	return "Error encoding response"
}
