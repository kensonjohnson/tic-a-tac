package main

import "github.com/gdamore/tcell/v2"

type Welcome struct {
	showWelcome bool
}

func NewWelcome() Welcome {
	return Welcome{
		showWelcome: true,
	}
}

func (w *Welcome) Show() {
	w.showWelcome = true
}

func (w *Welcome) Hide() {
	w.showWelcome = false
}

func (w *Welcome) Render(screen tcell.Screen) {
	if !w.showWelcome {
		return
	}
	drawStyledString(screen, 0, 1, blueForeground, "****************************************************************************")
	drawStyledString(screen, 0, 2, yellowForeground, "             __        __   _                            _        ")
	drawStyledString(screen, 0, 3, yellowForeground, "__/\\__       \\ \\      / /__| | ___ ___  _ __ ___   ___  | |_ ___  ")
	drawStyledString(screen, 0, 4, yellowForeground, "\\    /        \\ \\ /\\ / / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\ | __/ _ \\ ")
	drawStyledString(screen, 0, 5, yellowForeground, "/_  _\\         \\ V  V /  __/ | (_| (_) | | | | | |  __/ | || (_) |")
	drawStyledString(screen, 0, 6, redForeground, "  \\/            \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___|  \\__\\___/ ")

	drawStyledString(screen, 0, 8, redForeground, " _____ ___ ____      _      _____  _    ____             ")
	drawStyledString(screen, 0, 9, redForeground, "|_   _|_ _/ ___|    / \\    |_   _|/ \\  / ___|      __/\\__")
	drawStyledString(screen, 0, 10, greenForeground, "  | |  | | |       / _ \\     | | / _ \\| |          \\    /")
	drawStyledString(screen, 0, 11, greenForeground, "  | |  | | |___   / ___ \\    | |/ ___ \\ |___       /_  _\\")
	drawStyledString(screen, 0, 12, greenForeground, "  |_| |___\\____| /_/   \\_\\   |_/_/   \\_\\____|        \\/  ")

	drawStyledString(screen, 0, 14, blueForeground, "****************************************************************************")

	drawStyledString(screen, 0, 16, greenForeground, "Get ready to play an exciting game of Tic A Tac!")

	drawStyledString(screen, 0, 18, greenForeground, "Tic A Tac Game Rules:")
	drawStyledString(screen, 0, 19, greenForeground, "1. The game is played on a 5x5 grid.")
	drawStyledString(screen, 0, 20, greenForeground, "2. You will play as 'X' and the will be 'O'.")
	drawStyledString(screen, 0, 21, greenForeground, "3. Both players make their guess and reveal at the same time.")
	drawStyledString(screen, 0, 22, greenForeground, "4a. If both players guess the same square, it becomes blocked.")
	drawStyledString(screen, 0, 23, greenForeground, "4b. Otherwise, each player gets the guess they made.")
	drawStyledString(screen, 0, 24, greenForeground, "5. The game continues until all squares are marked or blocked.")
	drawStyledString(screen, 0, 25, greenForeground, "6a. Players get 1 point for every three in a row.")
	drawStyledString(screen, 0, 26, greenForeground, "6b. Players get 2 points for every four in a row.")
	drawStyledString(screen, 0, 27, greenForeground, "6c. Players get 3 points for every five in a row.")

	drawStyledString(screen, 0, 29, blueForeground, "Press enter to play!")

}
