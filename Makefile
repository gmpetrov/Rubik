
SRC= main.go \
	 Cube.go \
	 Solver.go
VALID_ARG="R2 D’ B’ D F2 R F2 R2 U L’ F2 U’ B’ L2 R D B’ R’ B2 L2 F2 L2 R2 U2 D2"
NOT_VALID_ARG="salut ca va233 "

all:
	go run $(SRC) $(VALID_ARG)
test:
	go test -v
