package server

func Init() {
	r := CreateRouter()

	r.Listen(":5000")
}
