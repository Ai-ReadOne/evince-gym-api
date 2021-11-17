package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// creating an object for the gym customers
type GymMember struct {
	LastName          string `json:"lastname"`
	FirstName         string `json:"firstname"`
	GymID             string `json:"gymid"`
	Gender            string `json:"gender"`
	Age               string `json:"age"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Instructor        string `json:"instructor"`
	LastSeen          string `json:"lastseen"`
	MebershipValidity string `json:"membershipvalidity"`
}

// creating an onject for the gym instructors
type GymInstructor struct {
	LastName     string `json:"lastname"`
	FirstName    string `json:"firstname"`
	InstructorID string `json:"instructorid"`
	Gender       string `json:"gender"`
	Age          string `json:"age"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	EmpDate      string `json:"empdate"`
}

// creating a record object that stores all the instructors and gym customers
type Records struct {
	sync.Mutex
	members     map[string]GymMember
	instructors map[string]GymInstructor
}

// checking the url and passing a handler aed on the passed url
// i did this so as to make sure that each action is perfomed by its intended Handler
func (rec *Records) resolveUrl(wr http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.String(), "/")
	if len(path) == 2 && path[1] == "" && req.Method == "GET" {
		rec.welcome(wr, req)
	} else if len(path) == 2 && path[1] == "all-members" && req.Method == "GET" {

		rec.allMembers(wr, req)

	} else if len(path) == 2 && path[1] == "all-instructors" && req.Method == "GET" {

		rec.allInstructors(wr, req)

	} else if len(path) == 2 && path[1] == "create-instructor" && req.Method == "POST" {

		rec.createNewInstructor(wr, req)

	} else if len(path) == 2 && path[1] == "create-member" && req.Method == "POST" {

		rec.createNewMember(wr, req)

	} else if len(path) == 3 && path[1] == "get-instructor" && req.Method == "GET" {

		rec.instructorByID(wr, req, path)

	} else if len(path) == 3 && path[1] == "get-member" && req.Method == "GET" {

		rec.memberByID(wr, req, path)

	} else {
		fmt.Printf(req.URL.String())
		rec.invalidUrl(wr, req)
	}
}

func (mem *Records) allMembers(wr http.ResponseWriter, req *http.Request) {
	fmt.Print("it is running gym members list")
	members := make([]GymMember, len(mem.members))
	mem.Lock()
	m := 0
	for _, member := range mem.members {
		members[m] = member
		m++
	}
	mem.Unlock()

	jsonBytes, error := json.Marshal(members)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.Header().Set("contnent-type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonBytes)
}

func (inst *Records) allInstructors(wr http.ResponseWriter, req *http.Request) {
	instructors := make([]GymInstructor, len(inst.instructors))

	inst.Lock()
	i := 0
	for _, instructor := range inst.instructors {
		instructors[i] = instructor
		i++
	}
	inst.Unlock()

	jsonBytes, error := json.Marshal(instructors)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.Header().Add("contnent-type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonBytes)
}

func (inst *Records) instructorByID(wr http.ResponseWriter, req *http.Request, path []string) {
	instructorID := path[2]
	inst.Lock()
	instructor, ok := inst.instructors[instructorID]
	inst.Unlock()
	if !ok {
		wr.WriteHeader(http.StatusNotFound)
		fmt.Print(ok)
		return

	}

	jsonBytes, error := json.Marshal(instructor)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.Header().Add("contnent-type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonBytes)
	fmt.Print(req)
}

func (mem *Records) memberByID(wr http.ResponseWriter, req *http.Request, path []string) {
	memberID := path[2]
	mem.Lock()
	member, ok := mem.members[memberID]
	mem.Unlock()
	if !ok {
		wr.WriteHeader(http.StatusNotFound)
		fmt.Print(ok)
		return

	}

	fmt.Print(mem.members[memberID])
	jsonBytes, error := json.Marshal(member)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.Header().Add("contnent-type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonBytes)
	fmt.Printf(memberID)
}

func (rec *Records) createNewMember(wr http.ResponseWriter, req *http.Request) {
	data, error := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if error != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write([]byte(error.Error()))
	}

	var newMember GymMember
	error = json.Unmarshal(data, &newMember)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write([]byte(error.Error()))
	}

	time := time.Now().UnixNano()
	newMember.GymID = fmt.Sprintf("%s%d", newMember.FirstName, time)
	wr.Header().Add("content-type", "application/json")

	rec.Lock()
	rec.members[newMember.GymID] = newMember
	defer rec.Unlock()

	wr.Write([]byte("new member created succesfully"))
	wr.WriteHeader(http.StatusOK)

}

func (rec *Records) createNewInstructor(wr http.ResponseWriter, req *http.Request) {

	data, error := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if error != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write([]byte(error.Error()))
	}

	var newInstructor GymInstructor
	error = json.Unmarshal(data, &newInstructor)
	if error != nil {
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write([]byte(error.Error()))
	}

	time := time.Now().UnixNano()
	newInstructor.InstructorID = fmt.Sprintf("%s%d", newInstructor.FirstName, time)

	wr.Header().Add("content-type", "application/json")
	rec.Lock()
	rec.instructors[newInstructor.InstructorID] = newInstructor
	defer rec.Unlock()
	wr.Write([]byte("new instructor created succesfully"))
	wr.WriteHeader(http.StatusOK)
}

func recordInit() *Records {
	fmt.Print("it is creating new gym members")
	return &Records{
		members: map[string]GymMember{
			"id1": {
				LastName:          "aransiola",
				FirstName:         "IRAHIM",
				GymID:             "KSKD",
				Gender:            "Male",
				Age:               "22",
				Phone:             "09080049446",
				Email:             "ridwanibrahim97@hotmail.com",
				Instructor:        "farouq202134",
				LastSeen:          "2020-12-5",
				MebershipValidity: "2021-12-30",
			},
		},
		instructors: map[string]GymInstructor{
			"id1": {
				LastName:     "aransiola",
				FirstName:    "IRAHIM",
				InstructorID: "KSKD",
				Gender:       "Male",
				Age:          "22",
				Phone:        "09080049446",
				Email:        "ridwanibrahim97@hotmail.com",
				EmpDate:      "2021-12-30",
			},
		},
	}
}

func (rec *Records) welcome(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(" Welcome the evince system Gym application"))
}

func (rec *Records) invalidUrl(wr http.ResponseWriter, req *http.Request) {
	link := req.URL.String()
	wr.WriteHeader(http.StatusNotFound)
	wr.Write([]byte(link))
}

func main() {
	records := recordInit()
	fmt.Println("This api is running on 127.0.0.1:8000")
	http.HandleFunc("/", records.resolveUrl)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
