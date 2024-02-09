# Term-Typing ⌨️🧠
## A Terminal-Based Typing App

![Screenshot of the interface](https://github.com/dustin-ward/Term-Typing/blob/master/images/v0.1.0UI.png)

### Usage

After executing the program, the text will appear greyed out. This means the keyboard is not active. Simply press 'enter' to begin typing.

Green characters have been correctly typed, while a red character means a wrong input has been received. You cannot advance until the correct key has been pressed.

Pressing 'enter' while typing will re-generate a new set of words and reset your position. Pressing 'esc' will return the application to the inactive state.

The program can be exited with 'ctrl+c' at any point in time, and with 'q' & 'esc' while in the inactive state.

### Installation
The only installation requirement is the Golang compiler.

Build the executable with:
```
go build -o ttype
```
