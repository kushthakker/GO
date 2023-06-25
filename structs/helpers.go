package main

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}