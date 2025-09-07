package models

type BPDU struct {
	RootId     int
	BridgeId   int
	MacAddr    string
	DInterface *Interface
}

func CreateBPDU(inter *Interface, dInter *Interface) *BPDU {
	return &BPDU{
		RootId:     inter.Priority,
		BridgeId:   inter.Priority,
		MacAddr:    inter.MacAddr,
		DInterface: dInter,
	}
}
