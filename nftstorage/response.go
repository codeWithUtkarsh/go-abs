package nftstorage

type Response200 struct {
	Ok    bool    `json:"ok"`
	Value RespPut `json:"value"`
}

type Response404 struct {
	Ok    bool    `json:"ok"`
	Error RespErr `json:"error"`
}

type FileInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DealInfo struct {
	BatchRootCid   string `json:"batchRootCid"`
	LastChange     string `json:"lastChange"`
	Miner          string `json:"miner"`
	Network        string `json:"network"`
	PieceCid       string `json:"pieceCid"`
	Status         string `json:"status"`
	StatusText     string `json:"statusText"`
	ChainDealID    int64  `json:"chainDealID"`
	DealActivation string `json:"dealActivation"`
	DealExpiration string `json:"dealExpiration"`
}

type PinStatus struct {
	CID     string                 `json:"cid"`
	Size    int64                  `json:"size"`
	Name    string                 `json:"name"`
	Meta    map[string]interface{} `json:"meta"`
	Status  string                 `json:"status"` // [ queued, pinning, pinned, failed ]
	Created string                 `json:"created"`
}

type RespPut struct {
	CID     string     `json:"cid"`
	Size    int64      `json:"size"`
	Created string     `json:"created"`
	Type    string     `json:"type"`
	Scope   string     `json:"scope"`
	Pin     PinStatus  `json:"pin"`
	Files   []FileInfo `json:"files"`
	Deals   []DealInfo `json:"deals"`
}

type RespErr struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type DeleteResponse200 struct {
	OK bool `json:"ok"`
}
