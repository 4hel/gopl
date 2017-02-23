TZ=US/Eastern go run clock.go -port 8010 &
TZ=Asia/Tokyo go run clock.go -port 8020 &
TZ=Europe/Berlin go run clock.go -port 8030 &
go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 Berlin=localhost:8030
