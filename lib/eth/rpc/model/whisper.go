package model

/*
Parameters

    Object - The whisper post object:

    from: DATA, 60 Bytes - (optional) The identity of the sender.
    to: DATA, 60 Bytes - (optional) The identity of the receiver. When present whisper will encrypt the message so that only the receiver can decrypt it.
    topics: Array of DATA - Array of DATA topics, for the receiver to identify messages.
    payload: DATA - The payload of the message.
    priority: QUANTITY - The integer of the priority in a range from ... (?).
    ttl: QUANTITY - integer of the time to live in seconds.
*/
type WhisperParams struct {
}

type FilterChanges struct {
	Hash       string   `json:"hash"`
	From       string   `json:"from"`
	To         string   `json:"to"`
	Expiry     string   `json:"expiry"`
	Sent       string   `json:"sent"`
	TTL        string   `json:"ttl"`
	Topics     []string `json:"topics"`
	Payload    string   `json:"payload"`
	WorkProved string   `json:"workProved"`
}
