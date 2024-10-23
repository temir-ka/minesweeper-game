# minesweeper-game
Overview

This project is a command-line simulation of the classic Minesweeper game implemented in Go. The game is played on a grid consisting of bombs (*) and empty cells (.). The objective is to reveal all non-bomb cells, while avoiding the bombs. The game incorporates traditional Minesweeper rules, including bomb counting and cascading reveals for empty cells.

Features

    - Grid-based Gameplay: Players input coordinates to reveal cells on the grid.
    - Bomb Relocation on First Move: If a bomb is chosen on the first move, it is relocated to a random empty cell, ensuring a fair start.
    - Cell Reveal Rules: Revealed cells show the number of adjacent bombs or remain empty. Cascading reveals occur for empty cells with no adjacent bombs.
    - Game End Conditions: The game ends with a "Game Over" if a bomb is revealed, or "You Win" when all non-bomb cells are uncovered.
    - Input Validation: The game enforces grid size and bomb count requirements (minimum grid size of 3x3 and at least 2 bombs). It also validates player input coordinates to ensure they are within the grid bounds.
    - Statistics: After the game ends, a summary is provided, including the number of moves, grid size, and bomb count.
    - Random Map Generation: Users can choose between entering a custom grid or generating a random grid with user-specified dimensions.
    - Colorful Numbers: Numbered cells indicating adjacent bombs are color-coded for improved readability.

How to Play

    1) Clone the repository
    2) Navigate to the project folder and run the game:
    3) Follow the prompts to either enter a custom grid or generate a random map.

Development Process

This project was built using the Double Diamond approach:

    Discover: Researched Minesweeper mechanics and identified key elements to simulate.
    Define: Established the game flow, rules, and necessary features.
    Develop: Implemented the core logic, focusing on grid generation, bomb placement, and reveal mechanics.
    Deliver: Tested the game thoroughly, handled edge cases, and added bonus features.

Technologies Used

    Go: Core programming language for game logic.
    fmt: For user input/output handling.
    math/rand: For random bomb placement in the game grid.
