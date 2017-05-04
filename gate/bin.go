package gate

import (
	"fmt"
	"time"
)

type SBin struct {
	m_id    uint32
	m_bQuit bool
}

func NewBin(id uint32) *SBin {
	pBin := &SBin{
		m_id:    id,
		m_bQuit: false,
	}
	return pBin
}

func (this *SBin) AsyncCmd(strCmd, strParam string) {

}

func (this *SBin) Update(u32ElapseMs uint32, now *time.Time) bool {
	// fmt.Println("Update", this.TypeId(), this.Id())
	fmt.Println("Update", this.TypeId(), this.Id(), u32ElapseMs, now)
	return !this.m_bQuit
}

func (this *SBin) TypeId() uint32 {
	return 99
}

func (this *SBin) Id() uint32 {
	return this.m_id
}

func (this *SBin) Quit() {
	fmt.Println(this.TypeId(), this.Id())
	this.m_bQuit = true
}

func (this *SBin) FrameMs() uint32 {
	return 1
}
