package server

func Init() {
	r := CreateRouter()

	r.Listen("127.0.0.1:5000")
}
