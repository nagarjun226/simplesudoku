# simplesudoku
A simple Golang package that computes the solution to a given Sudoku puzzle using backtracking

## Packages
- `/sudoku` - This is the main package that handles solving a sudoku
- `/api` - defines the handler funcs and the API that the app will expose

## Run
- `go run main.go` runs the sudoku solver as a server on specified port. Submit a Get request with te incomplete Sudoku as the body

## API Usage

**Definition**

`GET /solve`
solves the given sudoku input

**Request**
Request is a JSON containing an integer array.
- The size of the array should be 81
- Array can only contain numbers from 0-9 with 0 being the representation of an empty space
```json
{
    "sudoku": [] `integer array`
}
```

**Response**
- `500 Internal Server Error` on Processing issues
- `400 Bad Request` on incorrect inputs
- `200 OK` on Success
```json
{
    "sudoku": [] `integer array`
}
```

## ToDO
- Testing
- handle edge cases in solver
- API handler accepts list of lists input