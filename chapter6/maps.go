package main

import "fmt"

func main() {
	elements := make(map[string]map[string]string)

	elements["H"] = map[string]string{
		"name":  "Hydrogen",
		"state": "gas",
	}

	elements["He"] = map[string]string{
		"name":  "Helium",
		"state": "gas",
	}

	elements["Li"] = map[string]string{
		"name":  "Lithium",
		"state": "solid",
	}

	elements["Be"] = map[string]string{
		"name":  "Beryllium",
		"state": "solid",
	}

	elements["B"] = map[string]string{
		"name":  "Boron",
		"state": "solid",
	}

	elements["C"] = map[string]string{
		"name":  "Carbon",
		"state": "solid",
	}

	elements["N"] = map[string]string{
		"name":  "Nitrogen",
		"state": "gas",
	}

	elements["O"] = map[string]string{
		"name":  "Oxygen",
		"state": "gas",
	}

	elements["F"] = map[string]string{
		"name":  "Fluoride",
		"state": "gas",
	}

	elements["Ne"] = map[string]string{
		"name":  "Neon",
		"state": "gas",
	}

	if element, ok := elements["Li"]; ok {
		fmt.Println(element["name"], element["state"])
	}

	if name, ok := elements["Un"]; ok { // Will not run as Un is not a key
		fmt.Println(name, ok)
	}

}
