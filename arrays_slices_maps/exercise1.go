package main

import ("fmt")

func main() {
	var songs = make([]string, 4)
	songs[0] = "BlueWale"
	songs[1] = "Levitating"
	songs[2] = "Tuisi"
	songs[3] = "dskas"

	var preferredSongs = make([]string, 2)
	copy(preferredSongs, songs[2:])
	fmt.Println(preferredSongs)




}
