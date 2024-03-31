package auth

import (
	"fmt"
	"time"
)

type SessionManager struct {
	Lifetime   int
	SessionMap map[string]string
}

func NewSessionManager() *SessionManager {
	ret := new(SessionManager)
	ret.Lifetime = 10
	ret.SessionMap = make(map[string]string)
	return ret
}

func (sm *SessionManager) Add(key, value string) {
	sm.SessionMap[key] = value
	fmt.Println("Added")
	<-time.After(time.Duration(sm.Lifetime) * time.Second)
	delete(sm.SessionMap, key)
	fmt.Println("Deleted")
}
