PROGRAMS := \
	base64 \
	basename \
	cat \
	dirname \
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
	uuidgen \
	whoami \
	yes

ifeq ($(OS),Windows_NT)
	RMRF=rd /s /q
	EXEEXT=.exe
else
	RMRF=rm -rf
	EXEXT=
endif

all: $(PROGRAMS)

$(foreach program, $(PROGRAMS), $(program)):
	gom build -o ./bin/$@$(EXEEXT) ./src/$@

clean:
	$(RMRF) bin

.PHONY: all clean $(PROGRAMS)

