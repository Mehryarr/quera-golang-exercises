package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Player struct {
	name         string
	currentMap   *Map
	incomingChan chan string
	sync.Mutex
}

type Map struct {
	players      map[string]*Player
	incomingChan chan MapMessage
	sync.Mutex
}

type Game struct {
	maps    map[int]*Map
	players map[string]*Player
	sync.Mutex
}

type MapMessage struct {
	sender string
	msg    string
}

func NewGame(mapIds []int) (*Game, error) {
	maps := make(map[int]*Map, len(mapIds))
	for _, id := range mapIds {
		if id <= 0 {
			return nil, errors.New("mapId cannot be negative")
		}
		if maps[id] != nil {
			return nil, errors.New("mapIds cannot be duplicate")
		}
		maps[id] = &Map{make(map[string]*Player), make(chan MapMessage, 100), sync.Mutex{}}
		go maps[id].FanOutMessages()
	}

	return &Game{maps: maps, players: make(map[string]*Player)}, nil
}

func (g *Game) ConnectPlayer(name string) error {
	name = strings.ToLower(name)
	if name == "" {
		return errors.New("name cannot be empty")
	}

	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if _, ok := g.players[name]; ok {
		return errors.New("player already connected")
	}

	g.players[name] = &Player{
		name:         name,
		incomingChan: make(chan string, 100),
	}
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	name = strings.ToLower(name)
	if p, ok := g.players[name]; !ok {
		return errors.New("player not found")
	} else {
		if m, ok := g.maps[mapId]; !ok {
			return errors.New("map not found")
		} else {
			if p.currentMap != nil {
				delete(p.currentMap.players, p.name)
			}
			m.players[p.name] = p
			p.currentMap = m
			return nil
		}
	}
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	name = strings.ToLower(name)
	if p, ok := g.players[name]; ok {
		return p, nil
	}

	return nil, errors.New("player not found")
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	if mapId <= 0 {
		return nil, errors.New("mapId cannot be negative or zero ")
	}

	if m, ok := g.maps[mapId]; !ok {
		return nil, errors.New("no such map")
	} else {
		return m, nil
	}
}

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func (m *Map) FanOutMessages() {
	for {
		msg := <-m.incomingChan
		for _, p := range m.players {
			if p.name != msg.sender {
				p.incomingChan <- fmt.Sprintf("%s says: %s", capitalize(msg.sender), msg.msg)
			}
		}
	}
}

func (p *Player) GetChannel() <-chan string {
	return p.incomingChan
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SendMessage(msg string) error {
	if msg == "" {
		return errors.New("message cannot be empty")
	}

	if p.currentMap == nil {
		return errors.New("player is not in a map")
	}

	p.Mutex.Lock()
	p.currentMap.incomingChan <- MapMessage{sender: p.name, msg: msg}
	p.Mutex.Unlock()

	return nil
}
