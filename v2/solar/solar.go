package solar

import "fmt"

func Solar() map[int]string {
	continent := map[int]string{1001: "Mercury", 1002: "Venus", 1003: "Earth", 1004: "Mars", 1005: "Jupiter", 1006: "Saturn", 1007: "Uranus", 1008: "Neptune", 1009: "Pluto"}
	return continent
}
func Solarnames() []string {
	continent := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
	return continent
}

func SolarDistance(id int) map[int]string {
	continent := map[int]string{1011: "Mercury", 1012: "Venus", 1013: "Earth", 1014: "Mars", 1015: "Jupiter", 1016: "Saturn", 1017: "Uranus", 1018: "Neptune", 1019: "Pluto"}

	for i, cont := range continent {

	}

	return continent
}

type SolarStruct struct {
	ID   int
	Name string
}

var name map[string]SolarStruct

func SolarDistancev2(id int) SolarStruct {
	continent := map[int]string{1011: "Mercury", 1012: "Venus", 1013: "Earth", 1014: "Mars", 1015: "Jupiter", 1016: "Saturn", 1017: "Uranus", 1018: "Neptune", 1019: "Pluto"}

	var cname = SolarStruct{ID: 0000, Name: "Not Found"}

	for i, cont := range continent {
		fmt.Println("cont, i: ", i, cont)
		//fmt.Println("cont, i: ", reflect.Typeof(cont), i)
		if id == i {
			cname.ID = i
			cname.Name = cont
			break
		}
	}
	return cname
}
