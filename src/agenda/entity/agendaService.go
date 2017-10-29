package entity

import "container/list"

//AgendaService .
type AgendaService struct {
	storage *Storage
}

func (agendaS *AgendaService) createMeeting(meeting Meeting) bool {
	titleFilter := func(mt Meeting) bool {
		return mt.GetTitle() == meeting.GetTitle()
	}

	if agendaS.storage.QueryMeeting(titleFilter).Len() > 0 {
		return false
	}

	overlapFilter := func(mt Meeting) bool {
		pas := meeting.GetAllPAS()
		for _, p := range pas {
			if mt.ContainParticipator(p) {
				return !(mt.GetStartDate().After(meeting.GetEndDate()) || mt.GetStartDate().Equal(meeting.GetEndDate()) || mt.GetEndDate().Before(meeting.GetStartDate()) || mt.GetEndDate().Equal(meeting.GetStartDate()))
			}
		}
		return false
	}

	if agendaS.storage.QueryMeeting(overlapFilter).Len() > 0 {
		return false
	}

	/*
	   if !agendaS.storage.ContainUser(meeting.GetSponsor()) {
	     return false
	   }

	   for _, p := range meeting.GetParticipators() {
	     if !agendaS.storage.ContainUser(p) {
	       return false
	     }
	   }*/

	agendaS.storage.CreateMeeting(meeting)
	return true
}

//AddMeetingParticipator .
func (agendaS *AgendaService) AddMeetingParticipator(user, title, ptor string) {
	mt := agendaS.QueryMeetingByTitle(user, title)
	mt.AddParticipators(ptor)
	agendaS.storage.UpdateMeetingByTitle(title, mt)
}

//DeelteMeetingParticipator .
func (agendaS *AgendaService) DeelteMeetingParticipator(user, title, ptor string) {
	mt := agendaS.QueryMeetingByTitle(user, title)
	mt.DeleteParticipator(ptor)
	agendaS.storage.UpdateMeetingByTitle(title, mt)

}

//QueryMeetingByUserAndTime .
func (agendaS *AgendaService) QueryMeetingByUserAndTime(meeting Meeting) *list.List {
	timeAndUserFilter := func(mt Meeting) bool {
		return (mt.ContainParticipator(meeting.Sponsor) && mt.StartDate.After(meeting.StartDate) && mt.EndDate.Before(meeting.EndDate))
	}

	return agendaS.storage.QueryMeeting(timeAndUserFilter)
}

//QueryMeetingByTitle .
func (agendaS *AgendaService) QueryMeetingByTitle(user, title string) Meeting {
	titleFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor() && title == mt.GetTitle()
	}
	//可能nil
	return agendaS.storage.QueryMeeting(titleFilter).Front().Value.(Meeting)
}

//DeleteMeetingByTitle .
func (agendaS *AgendaService) DeleteMeetingByTitle(user, title string) bool {
	titleFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor() && title == mt.GetTitle()
	}

	return agendaS.storage.DeleteMeeting(titleFilter) > 0
}

//DeleteMeetingAll .
func (agendaS *AgendaService) DeleteMeetingAll(user string) bool {
	allFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor()
	}

	return agendaS.storage.DeleteMeeting(allFilter) > 0
}
