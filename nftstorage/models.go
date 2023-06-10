package nftstorage

type NftSuccessResponse struct {
	Ok    bool     `json:"ok"`
	Value Metadata `json:"value"`
}

type NftErrorResponse struct {
	Ok    bool         `json:"ok"`
	Error ErrorDetails `json:"error"`
}

type PinStatus struct {
	CID     string                 `json:"cid"`
	Size    int64                  `json:"size"`
	Name    string                 `json:"name"`
	Meta    map[string]interface{} `json:"meta"`
	Status  string                 `json:"status"` // [ queued, pinning, pinned, failed ]
	Created string                 `json:"created"`
}

type Metadata struct {
	CID     string    `json:"cid"`
	Size    int64     `json:"size"`
	Created string    `json:"created"`
	Type    string    `json:"type"`
	Scope   string    `json:"scope"`
	Pin     PinStatus `json:"pin"`
}

type ErrorDetails struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
