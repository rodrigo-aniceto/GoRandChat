var socket
var onChat


function searchChat() {
    if (socket){
        socket.close();
    }

    disableAll();
    onChat = false
    console.log ("Starting connection...");
    displayInfo ("Starting connection...");
    socket = new WebSocket("ws://localhost:5000/ws/chat");
    
    socket.onmessage = function(event) {
        const msg = JSON.parse(event.data);
            if (msg.code == 1){
                displayInfo(msg.text)
            } else if (msg.code == 2) {
                displayInfo(msg.text)
                startChat()
            } else {
                console.log ("CONNECTION ERROR")
                closeChatRemotely()
            }
    };
}

function startChat() {
    console.log ("starting chat")
    enableChat()
    onChat = true

    const chatOutput = document.getElementById("chat-output");
    chatOutput.innerHTML = "";

    socket.onmessage = function(event) {
        const msg = JSON.parse(event.data);

        if (msg.code == 3){
            displayMessage(msg.user, msg.text);
        } else if (msg.code == 2){
            displayInfo(msg.text)
        } else {
            displayInfo(msg.text)
            closeChatRemotely()
        }
    };
}


function checkMessageEnter(event){
    if ((event.key == 'Enter') && (onChat == true)){
        sendMessage()
    }
}



function sendMessage() {
    if ((socket)&&(onChat)) {
        const messageInput = document.getElementById("message-input");
        const message = messageInput.value;

        if (message.trim() !== "") {
            const msg = {
                user: UserName,
                text: message,
                code: 3
            };
            
            socket.send(JSON.stringify(msg));
            messageInput.value = "";

            displayMessage(UserName, message)            
        }
    }
}

function displayMessage (username, message) {
    
    const chatInfo = document.getElementById("chat-info");    
    chatInfo.innerHTML = ''

    const chatOutput = document.getElementById("chat-output");
    const messageDiv = document.createElement("div");
    
    
    messageDiv.innerHTML = '<strong>'+ username + ':</strong> ' + message;
    chatOutput.appendChild(messageDiv);

    const chatBox = document.getElementById('chat-box');
    chatBox.scrollTop = chatBox.scrollHeight;
}

function displayInfo (message) {
        
    const chatInfo = document.getElementById("chat-info");    
    chatInfo.innerHTML = message;
}

function closeChat() {
    if ((socket)&&(onChat)) {
        const msg = {
            user: UserName,
            text: "",
            code: 1
        };
        
        socket.send(JSON.stringify(msg));
        socket.close();
        onChat = false
        displayInfo ("Chat has been closed!");
    }
    disableChat()
}

function closeChatRemotely() {
    
    onChat = false
    socket.close()
    disableChat()
}

function enableChat (){
    document.getElementById('end-chat').disabled = false;
    document.getElementById('end-chat').style.display = ""
    document.getElementById('start-chat').disabled = true;
    document.getElementById('start-chat').style.display = "none"
    document.getElementById('chat-tools').disabled = false;
    document.getElementById('chat-tools').style.display = ""
}

function disableChat (){
    document.getElementById('end-chat').disabled = true;
    document.getElementById('end-chat').style.display = "none"
    document.getElementById('start-chat').disabled = false;
    document.getElementById('start-chat').style.display = ""
    document.getElementById('chat-tools').disabled = true;
    document.getElementById('chat-tools').style.display = "none"
}

function disableAll(){
    document.getElementById('end-chat').disabled = true;
    document.getElementById('end-chat').style.display = "none"
    document.getElementById('start-chat').disabled = true;
    document.getElementById('start-chat').style.display = "none"
    document.getElementById('chat-tools').disabled = true;
    document.getElementById('chat-tools').style.display = "none"
}

window.onload = searchChat()
