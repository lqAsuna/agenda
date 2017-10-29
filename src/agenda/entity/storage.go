package entity

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//Storage .
type Storage struct {
	userList    list.List
	meetingList list.List
	Current     *User
}

var instance *Storage
var mu sync.Mutex

//GetStorage .
func GetStorage() *Storage {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Storage{}
			instance.Current = &User{}
		}
	}
	return instance
}

//ReadCurUsr .
func (sto *Storage) ReadCurUsr(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	if err = json.Unmarshal(bytes, sto.Current); err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	fmt.Printf("%v\n", sto.Current)

	return nil
}

//WriteCurUsr .
func (sto *Storage) WriteCurUsr(filename string) error {
	ub, err := json.Marshal(sto.Current)
	if err != nil {
		fmt.Println("error:", err)
	}

	file1, err := os.Create(filename)
	defer file1.Close()
	if err != nil {
		fmt.Println(file1, err)
		return err
	}
	file1.Write(ub)
	return nil
}

//ReadFromFile .
func (sto *Storage) ReadFromFile(userfilename, meetingfilename string) error {
	var userSlice []User
	var meetingSlice []Meeting

	bytes, err := ioutil.ReadFile(userfilename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	if err = json.Unmarshal(bytes, &userSlice); err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	byte2s, err := ioutil.ReadFile(meetingfilename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	if err := json.Unmarshal(byte2s, &meetingSlice); err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	for _, v := range userSlice {
		sto.userList.PushBack(v)
	}

	for _, v := range meetingSlice {
		sto.meetingList.PushBack(v)
	}

	return nil
}

//WirteToFile .
func (sto *Storage) WirteToFile(userfilename, meetingfilename string) {
	var userSlice []User

	for e := sto.userList.Front(); e != nil; e = e.Next() {
		userSlice = append(userSlice, e.Value.(User))
	}

	ub, err := json.Marshal(userSlice)
	if err != nil {
		fmt.Println("error:", err)
	}

	file1, err := os.Create(userfilename)
	defer file1.Close()
	if err != nil {
		fmt.Println(file1, err)
		return
	}
	file1.Write(ub)

	var meetingSlice []Meeting

	for e := sto.meetingList.Front(); e != nil; e = e.Next() {
		meetingSlice = append(meetingSlice, e.Value.(Meeting))
	}

	mb, err := json.Marshal(meetingSlice)
	if err != nil {
		fmt.Println("error:", err)
	}

	file2, err := os.Create(meetingfilename)
	defer file2.Close()
	if err != nil {
		fmt.Println(file2, err)
		return
	}
	file2.Write(mb)

}

//CreateMeeting .
func (sto *Storage) CreateMeeting(meeting Meeting) {
	sto.meetingList.PushBack(meeting)
}

//QueryMeeting .
func (sto *Storage) QueryMeeting(filter func(meeting Meeting) bool) *list.List {
	returnList := list.New()
	for e := sto.meetingList.Front(); e != nil; e = e.Next() {
		if filter(e.Value.(Meeting)) {
			returnList.PushBack(e.Value.(Meeting))
		}
	}
	return returnList
}

//UpdateMeetingByTitle .
func (sto *Storage) UpdateMeetingByTitle(tl string, meeting Meeting) int {
	count := 0
	for mt := sto.meetingList.Front(); mt != nil; mt = mt.Next() {
		if mt.Value.(Meeting).Title == tl {
			mt.Value = meeting
			count++
		}
	}
	return count
}

//DeleteMeeting .
func (sto *Storage) DeleteMeeting(filter func(meeting Meeting) bool) int {
	count := 0
	var next *list.Element
	for mt := sto.meetingList.Front(); mt != nil; mt = next {
		next = mt.Next()
		if filter(mt.Value.(Meeting)) {
			sto.meetingList.Remove(mt)
			count++
		}
	}
	return count
}

/*
//ContainUser .
func (sto *Storage) ContainUser(user string) bool {
	for u := sto.userList.Front(); u != nil; u = u.Next() {
		if user == u.Value.(User).Name {
			return true
		}
	}
	return false
}*/

//PrintMU .
func (sto *Storage) PrintMU() {
	for e := sto.meetingList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(Meeting))
	}
}
