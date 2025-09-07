package models

type BPDU struct {
	RootId   int
	BridgeId int
}

func CreateBPDU(rid, bid int) *BPDU {
	return &BPDU{RootId: rid, BridgeId: bid}
}
