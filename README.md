# Easter Egg Hunt Hint Code

This repository contains a simple example program where the user gets hints about an easter egg's location by providing correct answers to questions. In the example code (Norwegian questions), there are two hidden eggs, and two possible hints per egg.

You can download the project via Git, or as a zip archive. We will assume you downloaded it to your HOME directory.

To build and run the program, you must first install [Go](https://golang.org/dl/). When you have done so, open a terminal and type:

On Windows:

    cd easter
    go build . -o easter.exe
    easter.exe

On Unix systems:

    cd easter
    go build . -o easter
    ./easter

If you are planning to run the program on another computer then where you are building it, you can rely on cross compilation via setting of environment variables. Here is an example building for Linux 32bit machines from another Unix environment:

    GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o easter-linux-386

See a full explanation of environment variables that affect Go builds [here](https://golang.org/cmd/go/#hdr-Environment_variables)

The code is to be regarded as an example.
