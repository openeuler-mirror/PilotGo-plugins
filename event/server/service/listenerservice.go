package service

import (
	"net/http"
	"strconv"
	"sync"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"openeuler.org/PilotGo/PilotGo-plugin-event/db"
)

type Listener struct {
	Name string
	URL  string
}

var (
	eventTypeMap   map[int][]Listener
	globalEventBus *EventBus
)

type EventBus struct {
	sync.Mutex
	listeners []*Listener
	stop      chan struct{}
	wait      sync.WaitGroup
	event     chan *common.EventMessage
}

func EventBusInit() {
	eventTypeMap = make(map[int][]Listener)
	globalEventBus = &EventBus{
		event: make(chan *common.EventMessage, 20),
		stop:  make(chan struct{}),
	}
	globalEventBus.Run()
}

func Stop() {
	globalEventBus.Stop()
}

func AddListener(l *Listener) {
	globalEventBus.AddListener(l)
}

func RemoveListener(l *Listener) {
	globalEventBus.RemoveListener(l)
}

func PublishEvent(m *common.EventMessage) {
	globalEventBus.publish(m)
}

func GetEventMapTypes(l *Listener) []int {
	return globalEventBus.GetEventMapTypes(l)
}
func AddEventMap(eventtype int, l *Listener) {
	globalEventBus.AddEventMap(eventtype, l)
}

func RemoveEventMap(eventtype int, l *Listener) {
	globalEventBus.RemoveEventMap(eventtype, l)
}

func RemoveEventMaps(l *Listener) {
	globalEventBus.RemoveEventMaps(l)
}

func IsExitEventMap(l *Listener) bool {
	return globalEventBus.IsExitEventMap(l)
}

// 添加监听事件
func (e *EventBus) AddListener(l *Listener) {
	e.Lock()
	defer e.Unlock()
	e.listeners = append(e.listeners, l)
}

// 删除监听事件
func (e *EventBus) RemoveListener(l *Listener) {
	e.Lock()
	defer e.Unlock()

	for index, v := range e.listeners {
		if v.Name == l.Name && v.URL == l.URL {
			if index == len(e.listeners)-1 {
				e.listeners = e.listeners[:index]
			} else {
				e.listeners = append(e.listeners[:index], e.listeners[index+1:]...)
			}
			break
		}
	}
}

func (e *EventBus) GetEventMapTypes(l *Listener) []int {
	e.Lock()
	defer e.Unlock()
	var eventTypes []int
	for eventType, values := range eventTypeMap {
		for _, v := range values {
			if v.Name == l.Name && v.URL == l.URL {
				eventTypes = append(eventTypes, eventType)
			}
		}
	}
	return eventTypes
}

// 添加event事件
func (e *EventBus) AddEventMap(eventtpye int, l *Listener) {
	e.Lock()
	defer e.Unlock()
	eventTypeMap[eventtpye] = append(eventTypeMap[eventtpye], *l)
}

// 删除event事件
func (e *EventBus) RemoveEventMap(eventtpye int, l *Listener) {
	e.Lock()
	defer e.Unlock()
	for i, v := range eventTypeMap[eventtpye] {
		if (v.Name == l.Name) && (v.URL == l.URL) {
			if i == len(eventTypeMap[eventtpye])-1 {
				eventTypeMap[eventtpye] = eventTypeMap[eventtpye][:i]
			} else {
				eventTypeMap[eventtpye] = append(eventTypeMap[eventtpye][:i], eventTypeMap[eventtpye][i+1:]...)
			}
			break
		}
	}
}

// 删除整个插件的event事件
func (e *EventBus) RemoveEventMaps(l *Listener) {
	e.Lock()
	defer e.Unlock()
	for i, value := range eventTypeMap {
		for j, v := range value {
			if (v.Name == l.Name) && (v.URL == l.URL) {
				if j == len(value)-1 {
					eventTypeMap[i] = eventTypeMap[i][:j]
				} else {
					eventTypeMap[i] = append(eventTypeMap[i][:j], eventTypeMap[i][j+1:]...)
				}
				break
			}
		}
	}
}

// 判断监听是否存在
func (e *EventBus) IsExitEventMap(l *Listener) bool {
	e.Lock()
	defer e.Unlock()
	for _, value := range eventTypeMap {
		for _, v := range value {
			if (v.Name == l.Name) && (v.URL == l.URL) {
				return true
			}
		}
	}
	return false
}

func (e *EventBus) Run() {
	go func(e *EventBus) {
		for {
			select {
			case <-e.stop:
				e.wait.Done()
			case m := <-e.event:
				listeners, ok := eventTypeMap[m.MessageType]
				if ok { //如果有插件订阅了此事件，将消息广播给订阅的各个插件
					e.broadcast(listeners, m)
					db.WriteToDB(m.MessageData)
				} else {
					db.WriteToDB(m.MessageData)
				}
			}
		}
	}(e)

}

func (e *EventBus) Stop() {
	e.wait.Add(1)
	e.stop <- struct{}{}
	e.wait.Done()
}

func (e *EventBus) publish(m *common.EventMessage) {
	e.event <- m
}

func (e *EventBus) broadcast(listeners []Listener, msg *common.EventMessage) {
	for _, listener := range listeners {
		r, err := httputils.Post(listener.URL+"/plugin_manage/api/v1/event", &httputils.Params{
			Body: msg,
		})
		if err != nil {
			logger.Error(listener.Name + "plugin process error:" + err.Error())
		}
		if r.StatusCode != http.StatusOK {
			logger.Error(listener.Name + "plugin process error:" + strconv.Itoa(r.StatusCode))
		}
	}
}
