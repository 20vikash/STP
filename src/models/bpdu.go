package models

type BPDU struct {
	RootId     int
	BridgeId   int
	MacAddr    string
	SInterface *Interface
	DInterface *Interface
}

func CreateBPDU(inter *Interface, dInter *Interface, mac string) *BPDU {
	return &BPDU{
		RootId:     inter.Priority,
		BridgeId:   inter.Priority,
		SInterface: inter,
		MacAddr:    mac,
		DInterface: dInter,
	}
}
