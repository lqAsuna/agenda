package entity

import (
	"time"
)

//Meeting .
type Meeting struct {
	Sponsor       string
	Participators []string
	StartDate     time.Time
	EndDate       time.Time
	Title         string
}

//GetMeeting .
func GetMeeting(sp string, ptor []string, sd, ed, tl string) Meeting {
	mt := Meeting{}
	mt.Sponsor = sp
	mt.Participators = ptor
	sdt, _ := time.Parse("2006-01-02 15:04", sd)
	edt, _ := time.Parse("2006-01-02 15:04", ed)
	mt.StartDate = sdt
	mt.EndDate = edt
	mt.Title = tl

	return mt
}

//GetSponsor .
func (m *Meeting) GetSponsor() string {
	return m.Sponsor
}

//SetSponsor .
func (m *Meeting) SetSponsor(sp string) {
	m.Sponsor = sp
}

//GetParticipators .
func (m *Meeting) GetParticipators() []string {
	return m.Participators
}

//AddParticipators .
func (m *Meeting) AddParticipators(ptor string) {
	m.Participators = append(m.Participators, ptor)
}

//DeleteParticipator .
func (m *Meeting) DeleteParticipator(ptor string) {

	for i := 0; i < len(m.Participators); i++ {
		if m.Participators[i] == ptor {
			m.Participators = append(m.Participators[:i], m.Participators[i+1:]...)
			break
		}
	}
}

//ContainParticipator .
func (m *Meeting) ContainParticipator(ptor string) bool {
	if m.Sponsor == ptor {
		return true
	}
	for _, p := range m.Participators {
		if p == ptor {
			return true
		}
	}
	return false
}

//GetStartDate .
func (m *Meeting) GetStartDate() time.Time {
	return m.StartDate
}

//SetStartDate .
func (m *Meeting) SetStartDate(sd time.Time) {
	m.StartDate = sd
}

//GetEndDate .
func (m *Meeting) GetEndDate() time.Time {
	return m.EndDate
}

//SetendDate .
func (m *Meeting) SetendDate(ed time.Time) {
	m.EndDate = ed
}

//GetTitle .
func (m *Meeting) GetTitle() string {
	return m.Title
}

//SetTitle .
func (m *Meeting) SetTitle(tl string) {
	m.Title = tl
}

//GetAllPAS .
func (m *Meeting) GetAllPAS() []string {
	pas := append(m.GetParticipators(), m.GetSponsor())
	return pas
}
