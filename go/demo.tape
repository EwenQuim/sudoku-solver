# Where should we write the GIF?
Output demo.gif

# Set up a 1200x600 terminal with 46px font.
Set FontSize 26
Set Width 1200
Set Height 600

# Type a command in the terminal.
Type "echo 'Lets solve this sudoku !!!'"
Enter

Type "cat data/sudoku.txt"
Enter

Sleep 1s

Type "go run . data/sudoku.txt"

# Pause for dramatic effect...
Sleep 500ms

# Run the command by pressing enter.
Enter

# Admire the output for a bit.
Sleep 5s
