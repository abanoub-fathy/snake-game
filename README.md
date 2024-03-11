# Snake Game 🐍
Welcome to the Golang Snake Game repository! This is a simple implementation of the classic Snake game written in [go programming language](https://go.dev/) and the [Ebiten](https://ebitengine.org/) open-source game engine.. Have fun playing and exploring the code!

<div align="center">
  <img src="https://github.com/abanoub-fathy/snake-game/assets/95833413/14af49c9-9904-4af7-bc37-f266e79cdd6e" alt="Snake Game">
</div>


## Table of Contents
- [Features](#features)
- [How to Play](#how-to-play)
- [Controls](#controls)
- [Installation](#installation)
- [Improvements](#improvements)
- [Contributing](#contributing)
- [License](#license)

## Features

- Control a snake to eat food and grow in size
- Snake wraps around the game board when reaching the edges
- Gradually increasing the snake's speed

## Controls

- Use the arrow keys to move the snake: ↑ (up), ↓ (down), ← (left), → (right)
- Use Space key to play again when the game is over.

## How to Play

You can play the game in two ways:
- Play in the Browser: Visit this [link](https://abanoub-fathy.github.io/snake-game/) to play the game directly in your browser.
- Play on your Machine: Install Golang on your machine, clone this repository, and build the game following the instructions mentioned in the Installation section below.

## Installation

To play the Snake game, you need to have Golang installed on your system. If you haven't installed it yet, you can download it from [Golang's official website](https://go.dev/).

Clone this repository using the following command:

```bash
git clone https://github.com/abanoub-fathy/snake-game.git
```

Change directory to the cloned repository:

```bash
cd snake-game
```

Install dependencies

```bash
go get
```

Build the game:

```bash
go build .
```

Once you've built the game, you can start playing by running the executable:

```bash
./snakeGame
```


## Improvements 

This is a basic Snake game implementation, and there's room for exciting additions:

- Sound Effects: Enhance the gameplay experience by incorporating sound effects for actions like eating food, hitting snake body, and game over.
- Levels: Introduce difficulty levels by changing the size of the play area, or introducing obstacles. This can add a layer of challenge and keep players engaged.
- Pausing the Game: Implement a pause functionality using keyboard input or a dedicated button to allow players to temporarily stop the game and resume later. This can be useful for taking breaks or strategizing their next move.

These are just a few ideas to get you started. Feel free to explore other creative enhancements to make this Snake game even more enjoyable!


## Contributing

We welcome contributions to this project! Please submit pull requests with clear descriptions of your changes. Ensure proper code formatting and testing before submitting.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

