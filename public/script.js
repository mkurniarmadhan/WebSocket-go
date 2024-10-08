const chatContainer = document.getElementById("chat-container");
const messageInput = document.getElementById("message-input");
const sendButton = document.getElementById("send-button");

const ws = new WebSocket("ws://localhost:8081/ws");

ws.onopen = () => {
  console.log("Connected to WebSocket");
};

ws.onmessage = (event) => {
  const message = event.data;
  const messageElement = document.createElement("div");
  messageElement.textContent = message;
  chatContainer.appendChild(messageElement);
};

sendButton.addEventListener("click", () => {
  const message = messageInput.value;
  ws.send(message);
  messageInput.value = "";
});
