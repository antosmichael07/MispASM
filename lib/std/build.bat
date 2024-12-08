cd lib
go mod init std-lib
go build -buildmode=c-shared -o std.misplib
del std.h
