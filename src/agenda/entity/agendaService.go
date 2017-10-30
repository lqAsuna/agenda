package entity

import (
	"container/list"
	"sync"
)

//AgendaService .
type AgendaService struct {
	storage *Storage
}

var agendaInstance *AgendaService
var agendamu sync.Mutex

//GetAgendaService .
func GetAgendaService() *AgendaService {
	if agendaInstance == nil {
		agendamu.Lock()
		defer agendamu.Unlock()
		if agendaInstance == nil {
			sto := GetStorage()
			agendaInstance = &AgendaService{}
			agendaInstance.storage = sto
		}
	}
	return agendaInstance
}

//GetAgendaServiceStorage .
func (agendaS *AgendaService) GetAgendaServiceStorage() *Storage {
	return agendaS.storage
}

//UserLogIn
func (agendaS *AgendaService) UserLogIn(user User) bool {
	filter := func(ur User) bool {
		return ur.GetName() == user.GetName()
	}
	if agendaS.storage.QueryUser(filter).size() <= 0 {
		return false
	}
	return true
}	

//UserRegister
func (agendaS *AgendaService) UserRegister(user User) bool {
	nameFilter := func(ur User) bool {
		return ur.GetName() == user.GetName()
	}
	if agendaS.storage.QueryUser(filter).size() > 0 {
		return false
	}
	agendaS.storage.CreateMeeting(user)
	return true
}

// DeleteUser(By name and password)
func (agendaS *AgendaService) DeleteUser(N string, P string) bool {
	filter := func(ur User) bool {
		return N == ur.GetName() && P == ur.GetPassword()
	}
	return agendaS.storage.DeleteUser(filter) > 0
}

//ListAllUsers
func (agendaS *AgendaService) QueryAllUsers() *list.List {
	filter := func(ur User) bool {
		return true
	}
	return agendaS.storage.QueryUser(filter)
}

func (agendaS *AgendaService) QueryMeetingByUserAndTime(meeting Meeting) *list.List {
	timeAndUserFilter := func(mt Meeting) bool {
		return (mt.ContainParticipator(meeting.Sponsor) && mt.StartDate.After(meeting.StartDate) && mt.EndDate.Before(meeting.EndDate))
	}

	return agendaS.storage.QueryMeeting(timeAndUserFilter)
}

// createMeeting
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

//AddMeetingParticipators .
func (agendaS *AgendaService) AddMeetingParticipators(user, title string, ptors []string) {
	mt := agendaS.QueryMyMeetingByTitle(user, title)
	for _, p := range ptors {
		mt.AddParticipators(p)
	}
	agendaS.storage.UpdateMeetingByTitle(title, mt)
}

//DeelteMeetingParticipators .
func (agendaS *AgendaService) DeelteMeetingParticipators(user, title string, ptors []string) {
	mt := agendaS.QueryMyMeetingByTitle(user, title)
	for _, p := range ptors {
		mt.DeleteParticipator(p)
	}
	agendaS.storage.UpdateMeetingByTitle(title, mt)

}

//Quit .
func (agendaS *AgendaService) Quit(user, title string) {
	mt := agendaS.QueryAllMeetingByTitle(user, title)
	if mt.Sponsor == user {
		mt.Sponsor = ""
	} else {
		mt.DeleteParticipator(user)
	}

	agendaS.storage.UpdateMeetingByTitle(title, mt)

}

//QueryMeetingByUserAndTime .
func (agendaS *AgendaService) QueryMeetingByUserAndTime(meeting Meeting) *list.List {
	timeAndUserFilter := func(mt Meeting) bool {
		return (mt.ContainParticipator(meeting.Sponsor) && mt.StartDate.After(meeting.StartDate) && mt.EndDate.Before(meeting.EndDate))
	}

	return agendaS.storage.QueryMeeting(timeAndUserFilter)
}

//QueryMyMeetingByTitle .
func (agendaS *AgendaService) QueryMyMeetingByTitle(user, title string) Meeting {
	titleFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor() && title == mt.GetTitle()
	}
	//可能nil
	return agendaS.storage.QueryMeeting(titleFilter).Front().Value.(Meeting)
}

//QueryAllMeetingByTitle .
func (agendaS *AgendaService) QueryAllMeetingByTitle(user, title string) Meeting {
	titleFilter := func(mt Meeting) bool {
		return mt.ContainParticipator(user) && title == mt.GetTitle()
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
