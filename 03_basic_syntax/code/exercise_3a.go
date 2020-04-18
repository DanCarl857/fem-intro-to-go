package main

func main() {
	sen := "This is my sentence"

	for index, letter := range sen {
		if index%2 != 0 {
			fmt.println(string(letter))
		}
	}
}
