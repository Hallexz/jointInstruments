# Collaborative Text Editor with Melody

This is a simple collaborative text editor application built using the Melody WebSocket library in Go.

## Features

- Real-time collaboration: Multiple users can edit the same document simultaneously, and their changes are reflected in real-time.
- Basic operations: Users can perform the following operations:
  - Add characters
  - Delete characters
  - Subscribe to the document
  - Unsubscribe from the document

## Usage

1. Clone the repository:
   ```
   git clone https://github.com/your-username/collaborative-text-editor.git
   ```

2. Navigate to the project directory:
   ```
   cd collaborative-text-editor
   ```

3. Build the application:
   ```
   go build -o main .
   ```

4. Run the application:
   ```
   ./main
   ```

5. Open your web browser and navigate to `http://localhost:5000`. You should see the text editor.

6. Open multiple browser windows or tabs to simulate multiple users collaborating on the document.

7. Perform various operations (add, delete, subscribe, unsubscribe) by sending messages in the following format:
   - Add character: `add:character`
   - Delete character: `delete:position`
   - Subscribe: `subscribe:username`
   - Unsubscribe: `unsubscribe:username`

   The operations will be logged in the server console.

## Docker

You can also run the application using Docker:

1. Build the Docker image:
   ```
   docker build -t collaborative-text-editor .
   ```

2. Run the Docker container:
   ```
   docker run -p 5000:5000 collaborative-text-editor
   ```

3. Open your web browser and navigate to `http://localhost:5000` to access the text editor.

## Dependencies

- [Melody](https://github.com/olahol/melody): A WebSocket library for Go.

## Contributing

If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

