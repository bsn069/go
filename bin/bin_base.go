package bin

import (
	"fmt"
	"time"
)

type SBinBase struct {
	M_typeId     uint32
	M_id         uint32
	M_bQuit      bool
	M_u32FrameMs uint32
	M_cmd        map[string]func(strParam string, bHelp bool) bool
}

func InitBinBase(typeId, id uint32, pBin *SBinBase) {
	pBin.M_typeId = typeId
	pBin.M_id = id
	pBin.M_bQuit = false
	pBin.M_u32FrameMs = 1000
	pBin.M_cmd = make(map[string]func(strParam string, bHelp bool) bool)

	pBin.M_cmd["SetFrameMs"] = func(strParam string, bHelp bool) bool {
		pBin.SetFrameMs(300)
		return true
	}
}

func (this *SBinBase) AsyncCmd(strCmd, strParam string) {
	if funcExec, bOk := this.M_cmd[strCmd]; bOk {
		fmt.Println("run cmd " + strCmd)
		if !funcExec(strParam, false) {
			funcExec(strParam, true)
		}
	}
}

func (this *SBinBase) Update(u32ElapseMs uint32, now *time.Time) bool {
	fmt.Println("Update", this.TypeId(), this.Id(), u32ElapseMs, now)
	return !this.M_bQuit
}

func (this *SBinBase) TypeId() uint32 {
	return this.M_typeId
}

func (this *SBinBase) Id() uint32 {
	return this.M_id
}

func (this *SBinBase) Quit() {
	fmt.Println(this.TypeId(), this.Id())
	this.M_bQuit = true
}

func (this *SBinBase) FrameMs() uint32 {
	return this.M_u32FrameMs
}

func (this *SBinBase) SetFrameMs(u32Ms uint32) {
	this.M_u32FrameMs = u32Ms
}
