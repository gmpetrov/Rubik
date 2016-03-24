
SRC= main.go \
	 Cube.go \
	 Solver.go \
	 Moves.go
VALID_ARG="R2 D’ B’ D F2 R F2 R2 U L’ F2 U’ B’ L2 R D B’ R’ B2 L2 F2 L2 R2 U2 D2 F2 B' L' B2 R'"
NOT_VALID_ARG="salut ca va233 "
VALID_ARG_BASIC="R2"

all:
	go run $(SRC) $(VALID_ARG)
test:
	go test -v
