package main

var planets = map[int]string{
	1: "Mercury",
	2: "Venus",
	3: "Earth",
	4: "Mars",
	5: "Jupiter",
	6: "Saturn",
	7: "Uranus",
	8: "Neptune",
}

func GetPlanetName(ID int) string {
	if planet, ok := planets[ID]; ok {
		return planet
	}
	return "Pluto"
}
