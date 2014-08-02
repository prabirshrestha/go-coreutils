PROGRAMS := \
	base64 \
	env \
	false \
	logname \
	ls \
	mkdir \
	mktemp \
	pwd \
	realpath \
	seq \
	sleep \
	true \
	uptime \
	whoami

ifeq ($(OS),Windows_NT)
	RMRF=rd /s /q
	EXEEXT=.exe
else
	RMRF=rm -rf
	EXEXT=
endif

all: $(PROGRAMS)

$(foreach program, $(PROGRAMS), $(program)):
	go build -o ./bin/$@$(EXEEXT) ./src/$@

clean:
	$(RMRF) bin

.PHONY: all clean $(PROGRAMS)
