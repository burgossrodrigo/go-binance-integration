# Binance Integration

This project fetches OHLC (Open, High, Low, Close) data from Binance, calculates resistance and support levels, and executes trades based on these levels. The trades are then stored in a MongoDB database.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go
- MongoDB

### Installing

1. Clone the repository:

```bash
git clone https://github.com/yourusername/binance-integration.git
```

2. Navigate to the project directory:

```bash
cd binance-integration
```

3. Install the dependencies:

```bash
go get
```


## Running the Application

To run the application, use the `go run` command:

```bash
go run main.go
```


The application will immediately execute trades based on the current resistance and support levels, and then continue to do so every 6 hours.

## Built With

- [Go](https://golang.org/) - The programming language used
- [Binance API](https://binance-docs.github.io/apidocs/spot/en/) - Used to fetch OHLC data
- [TA-Lib](https://mrjbq7.github.io/ta-lib/) - Used to calculate resistance and support levels
- [MongoDB](https://www.mongodb.com/) - Used to store trade data

## Contributing

We love contributions from everyone. Contributing is a great way to learn more about blockchain and trading, and to engage with our active community.

### How to Contribute

1. Fork the repository on GitHub.
2. Clone the forked repository to your machine.
3. Make your changes and commit them to your local repository.
4. Push the changes to your forked repository on GitHub.
5. Submit a pull request to the main repository.

### Contribution Guidelines

- Please ensure your pull request adheres to the following guidelines:
  - Search previous suggestions before making a new one, as yours may be a duplicate.
  - Make an individual pull request for each suggestion.
  - Use the following format for commit messages: `Add X functionality` or `Fix issue in Y`.
  - New categories, or improvements to the existing categorisation are welcome.
  - Keep descriptions short and simple, but descriptive.
  - Start the description with a capital and end with a full stop/period.
  - Check your spelling and grammar.
  - Make sure your text editor is set to remove trailing whitespace.

### Code of Conduct

We are committed to providing a friendly, safe and welcoming environment for all. Please take a moment to read our [Code of Conduct](CODE_OF_CONDUCT.md).

### Questions

If you have any questions, please feel free to contact me.

## License

No license