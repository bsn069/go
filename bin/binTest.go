package bin

import (
	"fmt"
	"time"
)

// for test
type SBinTest struct {
	m_typeId uint32
	m_id     uint32
	m_bQuit  bool
}

func NewBinTest(typeId, id uint32) *SBinTest {
	pBinTest := &SBinTest{
		m_typeId: typeId,
		m_id:     id,
		m_bQuit:  false,
	}
	return pBinTest
}

func (this *SBinTest) AsyncCmd(strCmd, strParam string) {

}

func (this *SBinTest) Update(u32ElapseMs uint32, now *time.Time) bool {
	// fmt.Println("Update", this.TypeId(), this.Id(), u32ElapseMs, now)
	return !this.m_bQuit
}

func (this *SBinTest) TypeId() uint32 {
	return this.m_typeId
}

func (this *SBinTest) Id() uint32 {
	return this.m_id
}

func (this *SBinTest) Quit() {
	fmt.Println(this.TypeId(), this.Id())
	this.m_bQuit = true
}

func (this *SBinTest) FrameMs() uint32 {
	return 1000
}
