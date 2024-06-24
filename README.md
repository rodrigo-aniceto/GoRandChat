# GoRandChat
This repository contains a randon chat application written in Go. Users can connect to the chat server, choose a username, and exchange messages in real-time to randon people.

## How It Works
1. **WebSocket Communication:** The application uses WebSockets for real-time communication (sending and receiving messages) between the web browsers and the server.
2. **Username Selection:** Upon connecting, users are prompted to choose a username. They should also agree to the terms of use.
3. **Chat Messages:** Users can send chat messages to the server, which then broadcasts the message to all connected clients. Messages are displayed in real-time on the client side.

## Getting Started
Follow these steps to build and run the chat application:

1. Clone the Repository:
```
git clone https://github.com/rodrigo-aniceto/GoRandChat.git
cd GoRandChat
```

2. Build the Application:
```
go build
```
3. Run the Server:
```
./GoRandChat
```

4. **Access the Chat:** Open your web browser and navigate to `http://localhost:5000`. You’ll see the start screen.
   
## Dependencies
The application uses the Gorilla WebSocket library for handling WebSocket connections. Make sure you have it installed (you can use `go get github.com/gorilla/websocket`).

## Customization
Feel free to customize the application further:

- Write the welcome text and terms of use.
- Implement private messaging using encryption.
- Enhance the UI with CSS and JavaScript.
- Check if there is anything that can be optimized on the server

## Contributions
Contributions are welcome! If you find any issues or have ideas for improvements, feel free to open a pull request.

Happy chatting! 🚀
