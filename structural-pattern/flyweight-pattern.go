package main

type Player struct {
	ID                    uint32
	Handle, Name, Country string
	Games                 []string
}

var (
	playerIDMap = make(map[uint32]Player)
	playerHandleMap = make(map[string]Player)
	playerCountryMap = make(map[string][]Player)
)

func FindPlayerByID(ID uint32) Player {
	if p, ok := playerIDMap[ID]; ok {
		return p
	}
	return Player{}
}

func FindPlayerByHandle(handle string) Player {
    if p, ok := playerHandleMap[handle]; ok {
        return p
    }
    return Player{}
}

func FindPlayerByCountry(code string) []Player {
    if ps, ok := playerCountryMap[code]; ok {
        return ps
    }
    return []Player{}
}


func main() {
	playerOne := Player {
		ID: 1,
		Handle: "Razzor",
		Name: "Dave",
		Country: "US",
		Games: []string("COD", "PES")
	}
}
